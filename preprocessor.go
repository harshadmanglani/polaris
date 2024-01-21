package polaris

import (
	"container/list"
	"fmt"
)

type ExecutionGraph struct {
	DependencyHierarchy [][]BuilderMeta
}

type ExecutionGraphGenerator struct {
	metaDataManager MetaDataManager
}

func (e *ExecutionGraphGenerator) generateExecGraph(dataFlow *DataFlow) ExecutionGraph {
	if dataFlow.TargetData == "" {
		panic("NO_TARGET_DATA")
	}
	dependencyGraph := make(map[string][]BuilderMeta)
	inDegree := make(map[string]int, 0)
	executionGraph := ExecutionGraph{
		DependencyHierarchy: make([][]BuilderMeta, 0),
	}

	for name, builderMeta := range e.metaDataManager.builderMetaMap {
		consumes := builderMeta.EffectiveConsumes()
		for _, c := range consumes.ToSlice() {
			fmt.Println(c)
			if _, ok := dependencyGraph[name]; !ok {
				dependencyGraph[name] = make([]BuilderMeta, 0)
			}
			if producedBy, ok := e.metaDataManager.producedToProducerMap[c]; ok {
				if val, ok := e.metaDataManager.builderMetaMap[producedBy.Name]; ok {
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

			currentLevel = append(currentLevel, e.metaDataManager.builderMetaMap[node])

			for _, builder := range graph[node] {
				inDegree[builder]--
				if inDegree[builder] == 0 {
					queue.PushBack(builder)
				}
			}
		}

		executionGraph.DependencyHierarchy = append(executionGraph.DependencyHierarchy, currentLevel)
	}
	return executionGraph
}
