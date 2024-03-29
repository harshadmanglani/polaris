package polaris

import (
	"container/list"
	"fmt"
	"reflect"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
)

var dataStore IDataStore

func InitRegistry(ds IDataStore) {
	dataStore = ds
}

func RegisterWorkflow(workflowKey string, workflow IWorkflow) error {
	if dataStore == nil {
		sugar.Errorf("Datastore uninitialized. Could not register: ", reflect.TypeOf(workflow))
		return fmt.Errorf("DATASTORE_UNINITIALIZED")
	}

	if workflow.GetWorkflowMeta().TargetData == nil {
		sugar.Errorf("Empty target data. Could not register: ", reflect.TypeOf(workflow))
		return fmt.Errorf("TARGET_DATA_NIL")
	}

	sugar.Infof("Registering workflow: %s", reflect.TypeOf(workflow))

	metaDataManager, err := newMetaDataManager(workflow)
	if err != nil {
		return err
	}

	dataFlow := newDataFlow(workflow, &metaDataManager)
	dataStore.Write(workflowKey, dataFlow)

	sugar.Infof("Registering %s took ", reflect.TypeOf(workflow), time.Now())
	return nil
}

func buildSet(data []IData) mapset.Set[string] {
	set := mapset.NewSet[string]()
	for _, d := range data {
		set.Add(Name(d))
	}
	return set
}

func newBuilderMeta(builder IBuilder) BuilderMeta {
	builderInfo := builder.GetBuilderInfo()
	return BuilderMeta{
		Consumes:  buildSet(builderInfo.Consumes),
		Optionals: buildSet(builderInfo.Optionals),
		Accesses:  buildSet(builderInfo.Accesses),
		Produces:  Name(builderInfo.Produces),
		Name:      Name(builder),
		Type:      reflect.TypeOf(builder),
	}
}

type DataFlow struct {
	Name                string
	TargetData          string
	DependencyHierarchy [][]BuilderMeta
	metaDataManager     MetaDataManager
}

func newDataFlow(workflow IWorkflow, metaDataManager *MetaDataManager) DataFlow {
	return DataFlow{
		Name:                Name(workflow),
		TargetData:          Name(workflow.GetWorkflowMeta().TargetData),
		metaDataManager:     *metaDataManager,
		DependencyHierarchy: generateDependencyHierarchy(metaDataManager),
	}
}

func generateDependencyHierarchy(metaDataManager *MetaDataManager) [][]BuilderMeta {

	dependencyGraph := make(map[string][]BuilderMeta)
	inDegree := make(map[string]int, 0)
	depedencyHierarchy := make([][]BuilderMeta, 0)

	for name, builderMeta := range metaDataManager.builderMetaMap {
		consumes := builderMeta.EffectiveConsumes()
		for _, c := range consumes.ToSlice() {
			if _, ok := dependencyGraph[name]; !ok {
				dependencyGraph[name] = make([]BuilderMeta, 0)
			}
			if producedBy, ok := metaDataManager.producedToProducerMap[c]; ok {
				if val, ok := metaDataManager.builderMetaMap[producedBy.Name]; ok {
					dependencyGraph[name] = append(dependencyGraph[name], val)
				}
			}
		}
	}

	graph := make(map[string][]string)
	for node, dependencies := range dependencyGraph {
		inDegree[node] = 0
		for _, dependency := range dependencies {
			graph[dependency.Name] = append(graph[dependency.Name], node)
			inDegree[node]++
		}
	}

	queue := list.New()
	for node := range dependencyGraph {
		if inDegree[node] == 0 {
			queue.PushBack(node)
		}
	}

	for queue.Len() > 0 {
		levelSize := queue.Len()
		currentLevel := []BuilderMeta{}

		for i := 0; i < levelSize; i++ {
			element := queue.Front()
			node := element.Value.(string)
			queue.Remove(element)

			currentLevel = append(currentLevel, metaDataManager.builderMetaMap[node])

			for _, builder := range graph[node] {
				inDegree[builder]--
				if inDegree[builder] == 0 {
					queue.PushBack(builder)
				}
			}
		}

		depedencyHierarchy = append(depedencyHierarchy, currentLevel)
	}
	return depedencyHierarchy
}
