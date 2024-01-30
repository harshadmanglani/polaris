package polaris

import (
	"fmt"
	"reflect"

	mapset "github.com/deckarep/golang-set/v2"
)

type Executor struct {
	Before func(builder reflect.Type, delta []IData) // TODO: add trigger delta
	After  func(builder reflect.Type, produced IData)
}

func checkForConsumes(dataSet *DataSet, builderInfo BuilderInfo) bool {
	fmt.Println(builderInfo)
	for _, consumes := range builderInfo.Consumes {
		if _, ok := dataSet.AvailableData[Name(consumes)]; !ok {
			return false
		}
	}
	return true
}

func (e *Executor) Run(workflowKey string, workflowId string, data ...IData) DataExecutionResponse {
	depedencyHierarchy := workflowStore[workflowKey] // TODO: add error handling
	var dataSet DataSet = dataStore[workflowId]
	responseData := make(map[string]IData)
	if _, ok := dataStore[workflowId]; !ok {
		dataSet = DataSet{
			AvailableData: make(map[string]IData),
		}
	}
	activeDataSet := mapset.NewSet[string]()
	for _, d := range data {
		dataSet.AvailableData[Name(d)] = d
		activeDataSet.Add(Name(d))
	}

	processedBuilders := mapset.NewSet[BuilderMeta]()
	newlyGeneratedData := mapset.NewSet[string]()

	for true {
		for _, levelBuilders := range depedencyHierarchy.DependencyHierarchy {
			for _, builderMeta := range levelBuilders {
				// implement go routines here
				if processedBuilders.Contains(builderMeta) {
					continue
				}

				if builderMeta.EffectiveConsumes().Intersect(activeDataSet).IsEmpty() {
					continue
				}
				fmt.Println(builderMeta)
				builder := reflect.New(builderMeta.Type).Interface().(IBuilder)

				if !checkForConsumes(&dataSet, builder.GetBuilderInfo()) {
					continue
				}
				e.Before(builderMeta.Type, data)

				response := builder.Process(BuilderContext{
					DataSet: dataSet,
				})
				if response != nil {
					if Name(response) != builderMeta.Produces {
						// throw error
					}
					dataSet.AvailableData[Name(response)] = response
					activeDataSet.Add(Name(response))
					newlyGeneratedData.Add(Name(response))
					responseData[Name(response)] = response
				}
				processedBuilders.Add(builderMeta)
				e.After(builderMeta.Type, response)
			}
		}
		if newlyGeneratedData.Contains(depedencyHierarchy.TargetData) {
			break
		}
		if newlyGeneratedData.IsEmpty() {
			break
		}
	}

	return DataExecutionResponse{
		Responses: responseData,
	}
}
