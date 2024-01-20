package polaris

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
	var graph map[string][]BuilderMeta
	var inDegree map[string]int

	for name, builderMeta := range e.metaDataManager.builderMetaMap {
		consumes := builderMeta.EffectiveConsumes()
		for _, c := range <-consumes.Iter() {
			graph[name] = append(graph[name], e.metaDataManager.getMetaForProducerOf(string(c)))
			inDegree[name]++
		}
	}
	return ExecutionGraph{}
}
