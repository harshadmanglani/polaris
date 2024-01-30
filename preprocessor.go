package polaris

import (
	"container/list"
	"fmt"
	"reflect"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
)

var workflowStore map[string]DataFlow
var dataStore map[string]DataSet

func init() {
	fmt.Println("in init preprocessor")
	workflowStore = make(map[string]DataFlow)
	dataStore = make(map[string]DataSet)
	// This is where redis client would be initialized
}

type MetaDataManager struct {
	builders              map[string]IBuilder
	builderMetaMap        map[string]BuilderMeta
	producedToProducerMap map[string]BuilderMeta // TODO: make this []BuilderMeta when ResolutionSpec is implemented
}

func newMetaDataManager() MetaDataManager {
	return MetaDataManager{
		builders:              make(map[string]IBuilder),
		builderMetaMap:        make(map[string]BuilderMeta),
		producedToProducerMap: make(map[string]BuilderMeta),
	}
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

func (m *MetaDataManager) register(builder IBuilder) {
	builderMeta := newBuilderMeta(builder)
	if _, ok := m.builderMetaMap[builderMeta.Name]; ok {
		panic("Builder already exists")
	}
	m.builders[builderMeta.Name] = builder
	m.builderMetaMap[builderMeta.Name] = builderMeta
	m.producedToProducerMap[builderMeta.Produces] = builderMeta
}

type DataFlow struct {
	Name                string
	TargetData          string
	DependencyHierarchy [][]BuilderMeta
	metaDataManager     MetaDataManager
	// TODO: implement Transients
	// TODO: implement ResolutionSpec
}

func logTimeSince(message string, past time.Time) {
	fmt.Printf("%s: %s", message, time.Since(past))
}

func RegisterWorkflow(workflowKey string, workflow IWorkflow) {
	now := time.Now()
	defer logTimeSince("Time to register "+Name(workflow), now)
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
