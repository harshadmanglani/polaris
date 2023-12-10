package core

import (
	"github.com/harshadmanglani/polaris/models"
)

type ExecutionGraphGenerator struct {
	metaDataManager MetaDataManager
}

func (e *ExecutionGraphGenerator) generateExecGraph(dataFlow *models.DataFlow) models.ExecutionGraph {
	if dataFlow.TargetData == "" {
		panic("NO_TARGET_DATA")
	}
	return models.ExecutionGraph{}
}
