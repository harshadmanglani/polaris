package core

import (
	"fmt"

	"github.com/harshadmanglani/polaris/models"
)

func registerWorkflow(workflow models.IWorkflow) {
	for _, b := range workflow.GetWorkflowMeta().Builders {
		fmt.Printf("Registering Builder: %s", models.DataToString(b))
	}
	generate(workflow)
}
