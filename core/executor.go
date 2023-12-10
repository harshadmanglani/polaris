package core

import (
	"github.com/harshadmanglani/polaris/models"
)

type Executor struct {
	Before  func()
	After   func()
	OnError func()
}

func (e *Executor) Run(df *models.DataFlow, data ...models.IData) models.DataExecutionResponse {
	return models.DataExecutionResponse{}
}
