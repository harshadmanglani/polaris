package polaris

import (
	"fmt"
	"reflect"
	"sync"

	mapset "github.com/deckarep/golang-set/v2"
)

type Executor struct {
	Before func(builder reflect.Type, delta []IData) // TODO: add trigger delta
	After  func(builder reflect.Type, produced IData)
}

func checkForConsumes(dataSet *DataSet, builderInfo BuilderInfo) bool {
	for _, consumes := range builderInfo.Consumes {
		if _, ok := dataSet.AvailableData[Name(consumes)]; !ok {
			return false
		}
	}
	return true
}

func (e *Executor) Sequential(workflowKey string, workflowId string, data ...IData) (DataExecutionResponse, error) {
	defer func() {
		if err := recover(); err != nil {
			sugar.Errorf("Workflow execution panicked: %v", err)
		}
	}()

	if dataStore == nil {
		sugar.Errorf("Datastore uninitialized. Could not run workflow with key: ", workflowKey)
		return DataExecutionResponse{}, fmt.Errorf("DATASTORE_UNINITIALIZED")
	}

	var dataSet DataSet
	var dataFlow DataFlow
	if dataFlowInterface, ok := dataStore.Read(workflowKey); !ok {
		return DataExecutionResponse{}, fmt.Errorf("WORKFLOW_KEY_NOT_FOUND")
	} else {
		dataFlow = dataFlowInterface.(DataFlow)
	}

	if dataSetInterface, ok := dataStore.Read(workflowId); !ok {
		dataSet = DataSet{
			AvailableData: make(map[string]IData),
		}
	} else {
		dataSet = dataSetInterface.(DataSet)
	}

	responseData := make(map[string]IData)

	activeDataSet := mapset.NewSet[string]()
	for _, d := range data {
		dataSet.AvailableData[Name(d)] = d
		activeDataSet.Add(Name(d))
	}

	processedBuilders := mapset.NewSet[BuilderMeta]()
	newlyGeneratedData := mapset.NewSet[string]()

	for {
		for _, levelBuilders := range dataFlow.DependencyHierarchy {
			for _, builderMeta := range levelBuilders {
				if processedBuilders.Contains(builderMeta) {
					continue
				}

				if builderMeta.EffectiveConsumes().Intersect(activeDataSet).IsEmpty() {
					continue
				}
				builder := reflect.New(builderMeta.Type).Interface().(IBuilder)

				if !checkForConsumes(&dataSet, builder.GetBuilderInfo()) {
					continue
				}
				if e.Before != nil {
					e.Before(builderMeta.Type, data)
				}

				response := builder.Process(BuilderContext{
					DataSet: dataSet,
				})
				if response != nil {
					if Name(response) != builderMeta.Produces {
						sugar.Errorf("Builder %s did not produce %s, instead it produced %s", builderMeta.Name, builderMeta.Produces, Name(response))
						return DataExecutionResponse{}, fmt.Errorf("INVALID_PRODUCED_DATA")
					}
					dataSet.AvailableData[Name(response)] = response
					activeDataSet.Add(Name(response))
					newlyGeneratedData.Add(Name(response))
					responseData[Name(response)] = response
				}
				processedBuilders.Add(builderMeta)

				if e.After != nil {
					e.After(builderMeta.Type, response)
				}
			}
		}
		if newlyGeneratedData.Contains(dataFlow.TargetData) {
			break
		}
		if newlyGeneratedData.IsEmpty() {
			break
		}
		activeDataSet.Clear()
		activeDataSet = activeDataSet.Union(newlyGeneratedData)
		newlyGeneratedData.Clear()
	}

	return DataExecutionResponse{
		Responses: responseData,
	}, nil
}

/*
 * This is experimental and severely untested. Use with caution.
 */
func (e *Executor) Parallel(workflowKey string, workflowId string, data ...IData) (DataExecutionResponse, error) {
	defer func() {
		if err := recover(); err != nil {
			sugar.Errorf("Workflow execution panicked: %v", err)
		}
	}()
	
	var dataSet DataSet
	var dataFlow DataFlow
	if dataFlowInterface, ok := dataStore.Read(workflowKey); !ok {
		return DataExecutionResponse{}, fmt.Errorf("WORKFLOW_KEY_NOT_FOUND")
	} else {
		dataFlow = dataFlowInterface.(DataFlow)
	}

	if dataSetInterface, ok := dataStore.Read(workflowId); !ok {
		dataSet = DataSet{
			AvailableData: make(map[string]IData),
		}
	} else {
		dataSet = dataSetInterface.(DataSet)
	}

	responseData := make(map[string]IData)

	activeDataSet := mapset.NewSet[string]()
	for _, d := range data {
		dataSet.AvailableData[Name(d)] = d
		activeDataSet.Add(Name(d))
	}

	processedBuilders := mapset.NewSet[BuilderMeta]()
	newlyGeneratedData := mapset.NewSet[string]()

	for {
		for _, levelBuilders := range dataFlow.DependencyHierarchy {
			var wg sync.WaitGroup
			for _, builderMeta := range levelBuilders {
				wg.Add(1)
				go e.executeBuilder(processedBuilders, builderMeta, activeDataSet, dataSet, data, newlyGeneratedData, responseData, &wg)
			}
			wg.Wait()
		}
		if newlyGeneratedData.Contains(dataFlow.TargetData) {
			break
		}
		if newlyGeneratedData.IsEmpty() {
			break
		}
		activeDataSet.Clear()
		activeDataSet = activeDataSet.Union(newlyGeneratedData)
		newlyGeneratedData.Clear()
	}

	return DataExecutionResponse{
		Responses: responseData,
	}, nil
}

func (e *Executor) executeBuilder(processedBuilders mapset.Set[BuilderMeta],
	builderMeta BuilderMeta,
	activeDataSet mapset.Set[string],
	dataSet DataSet,
	data []IData,
	newlyGeneratedData mapset.Set[string],
	responseData map[string]IData,
	wg *sync.WaitGroup) {

	if processedBuilders.Contains(builderMeta) {
		wg.Done()
		return
	}

	if builderMeta.EffectiveConsumes().Intersect(activeDataSet).IsEmpty() {
		wg.Done()
		return
	}
	builder := reflect.New(builderMeta.Type).Interface().(IBuilder)

	if !checkForConsumes(&dataSet, builder.GetBuilderInfo()) {
		wg.Done()
		return
	}

	if e.Before != nil {
		e.Before(builderMeta.Type, data)
	}
	response := builder.Process(BuilderContext{
		DataSet: dataSet,
	})
	if response != nil {
		if Name(response) != builderMeta.Produces {
			sugar.Errorf("Builder %s did not produce %s, instead it produced %s", builderMeta.Name, builderMeta.Produces, Name(response))
			wg.Done()
			return
			// TODO: return error here
		}
		dataSet.AvailableData[Name(response)] = response
		activeDataSet.Add(Name(response))
		newlyGeneratedData.Add(Name(response))
		responseData[Name(response)] = response
	}
	processedBuilders.Add(builderMeta)
	if e.After != nil {
		e.After(builderMeta.Type, response)
	}
	wg.Done()
}
