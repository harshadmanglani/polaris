package polaris

type Executor struct {
	Before  func()
	After   func()
	OnError func()
}

func (e *Executor) Run(df *DataFlow, data ...IData) DataExecutionResponse {
	return DataExecutionResponse{}
}
