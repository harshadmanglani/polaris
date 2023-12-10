package core

import (
	"github.com/harshadmanglani/polaris/models"
)

type DataFlowBuilder struct {
	metaDataManager MetaDataManager
	dataFlow        models.DataFlow
}

func (dfb *DataFlowBuilder) buildDataFlow(workflow models.IWorkflow) {
	dfb.metaDataManager = MetaDataManager{}
	for _, b := range workflow.GetWorkflowMeta().Builders {
		dfb.metaDataManager.register(b)
	}
	dfb.dataFlow = models.DataFlow{
		Name:       models.Name(workflow),
		TargetData: models.Name(workflow.GetWorkflowMeta().TargetData),
	}
	execGraphGenerator := ExecutionGraphGenerator{metaDataManager: dfb.metaDataManager}
	dfb.dataFlow.ExecGraph = execGraphGenerator.generateExecGraph(&dfb.dataFlow)
}

func RegisterWorkflow(workflow models.IWorkflow) models.DataFlow {
	dfb := &DataFlowBuilder{}
	dfb.buildDataFlow(workflow)
	return dfb.dataFlow
}
