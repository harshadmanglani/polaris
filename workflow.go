package polaris

type IWorkflow interface {
	GetWorkflowMeta() WorkflowMeta
}

type WorkflowMeta struct {
	Builders   []IBuilder
	TargetData IData
}
