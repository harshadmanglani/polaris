package core

import (
	"fmt"

	"github.com/harshadmanglani/polaris/models"
)

func generate(workflow models.IWorkflow) models.ExecutionGraph {
	if workflow.GetWorkflowMeta().TargetData == nil {
		panic("EMPTY_TARGET_DATA")
	}
	builders := workflow.GetWorkflowMeta().Builders
	targetDataBuilderIndex := -1
	for i, b := range builders {
		if models.DataToString(b.GetBuilderMeta().Produces) == models.DataToString(workflow.GetWorkflowMeta().TargetData) {
			targetDataBuilderIndex = i
		}
	}
	fmt.Printf("Target data is produced by: %s", models.DataToString(builders[targetDataBuilderIndex]))
	return models.ExecutionGraph{}
}
