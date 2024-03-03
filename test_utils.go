package polaris

type workflowTerminated struct {
}

type alphaConsumes struct {
}

type alphaProduces struct {
}

type betaProduces struct {
}

type alphaBuilder struct {
}

func (a alphaBuilder) Process(context BuilderContext) IData {
	return alphaProduces{}
}

func (a alphaBuilder) GetBuilderInfo() BuilderInfo {
	return BuilderInfo{
		Consumes: []IData{
			alphaConsumes{},
		},
		Produces:  alphaProduces{},
		Optionals: nil,
		Accesses:  nil,
	}
}

type betaBuilder struct {
}

func (b betaBuilder) Process(context BuilderContext) IData {
	return betaProduces{}
}

func (b betaBuilder) GetBuilderInfo() BuilderInfo {
	return BuilderInfo{
		Consumes: []IData{
			alphaProduces{},
		},
		Produces:  betaProduces{},
		Optionals: nil,
		Accesses:  nil,
	}
}

type terminator struct {
}

func (t terminator) Process(context BuilderContext) IData {
	return workflowTerminated{}
}

func (t terminator) GetBuilderInfo() BuilderInfo {
	return BuilderInfo{
		Consumes: []IData{
			alphaProduces{},
			betaProduces{},
		},
		Produces:  workflowTerminated{},
		Optionals: nil,
		Accesses:  nil,
	}
}

type testWorkflow struct {
}

func (tW testWorkflow) GetWorkflowMeta() WorkflowMeta {
	return WorkflowMeta{
		Builders: []IBuilder{
			alphaBuilder{},
			betaBuilder{},
			terminator{},
		},
		TargetData: workflowTerminated{},
	}
}

type redundantBuilderFailureWorklow struct {
}

func (rBFW redundantBuilderFailureWorklow) GetWorkflowMeta() WorkflowMeta {
	return WorkflowMeta{
		Builders: []IBuilder{
			alphaBuilder{},
			alphaBuilder{},
			terminator{},
		},
		TargetData: workflowTerminated{},
	}
}

type emptyTargetDataWorkflow struct {
}

func (eTDW emptyTargetDataWorkflow) GetWorkflowMeta() WorkflowMeta {
	return WorkflowMeta{
		Builders: []IBuilder{
			alphaBuilder{},
			betaBuilder{},
			terminator{},
		},
	}
}

type mockStorage struct {
	store map[string]interface{}
}

func (ms *mockStorage) Read(key string) (interface{}, bool) {
	val, ok := ms.store[key]
	return val, ok
}

func (ms *mockStorage) Write(key string, val interface{}) {
	ms.store[key] = val
}
