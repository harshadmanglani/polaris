package example

import "github.com/harshadmanglani/polaris/models"

type OmsWorkflow struct {
}

func (omsW OmsWorkflow) GetWorkflowMeta() models.WorkflowMeta {
	return models.WorkflowMeta{
		Builders:   []models.IBuilder{Validator{}, RiskChecker{}},
		TargetData: WorkflowTerminated{},
	}
}

func main() {
}
