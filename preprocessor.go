package polaris

import (
	"container/list"
)

var workflowStore map[string]DataFlow
var dataStore map[string]DataSet

func init() {
	workflowStore = make(map[string]DataFlow)
	dataStore = make(map[string]DataSet)
	// This is where redis client would be initialized
}

type DataFlow struct {
	Name                string
	TargetData          string
	DependencyHierarchy [][]BuilderMeta
	metaDataManager     MetaDataManager
	// TODO: implement Transients
	// TODO: implement ResolutionSpec
}

func RegisterWorkflow(workflowKey string, workflow IWorkflow) {
	metaDataManager := newMetaDataManager()
	for _, b := range workflow.GetWorkflowMeta().Builders {
		metaDataManager.register(b)
	}
	dataFlow := DataFlow{
		Name:                Name(workflow),
		TargetData:          Name(workflow.GetWorkflowMeta().TargetData),
		metaDataManager:     metaDataManager,
		DependencyHierarchy: preprocess(&metaDataManager),
	}
	workflowStore[workflowKey] = dataFlow
}

// TODO:
// 1. refactor to optimize
// 2. measure time
// 3. rename
// 4. add error handling for null target data
func preprocess(metaDataManager *MetaDataManager) [][]BuilderMeta {
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
