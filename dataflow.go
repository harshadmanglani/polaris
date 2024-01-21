package polaris

type DataFlow struct {
	Name       string
	TargetData string
	ExecGraph  ExecutionGraph
	// TODO: implement Transients
	// TODO: implement ResolutionSpec
}

type DataFlowBuilder struct {
	metaDataManager MetaDataManager
	dataFlow        DataFlow
}

func (dfb *DataFlowBuilder) buildDataFlow(workflow IWorkflow) {
	dfb.metaDataManager = newMetaDataManager()
	for _, b := range workflow.GetWorkflowMeta().Builders {
		dfb.metaDataManager.register(b)
	}
	dfb.dataFlow = DataFlow{
		Name:       Name(workflow),
		TargetData: Name(workflow.GetWorkflowMeta().TargetData),
	}
	execGraphGenerator := ExecutionGraphGenerator{metaDataManager: dfb.metaDataManager}
	dfb.dataFlow.ExecGraph = execGraphGenerator.generateExecGraph(&dfb.dataFlow)
}

func RegisterWorkflow(workflow IWorkflow) DataFlow {
	dfb := &DataFlowBuilder{}
	dfb.buildDataFlow(workflow)
	return dfb.dataFlow
}
