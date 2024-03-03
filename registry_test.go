package polaris

import (
	"testing"

	mapset "github.com/deckarep/golang-set/v2"
)

type WorkflowTerminated struct {
}

type AlphaConsumes struct {
}

type AlphaProduces struct {
}

type BetaProduces struct {
}

type AlphaBuilder struct {
}

func (a AlphaBuilder) Process(context BuilderContext) IData {
	return AlphaProduces{}
}

func (a AlphaBuilder) GetBuilderInfo() BuilderInfo {
	return BuilderInfo{
		Consumes: []IData{
			AlphaConsumes{},
		},
		Produces:  AlphaProduces{},
		Optionals: nil,
		Accesses:  nil,
	}
}

type BetaBuilder struct {
}

func (b BetaBuilder) Process(context BuilderContext) IData {
	return AlphaProduces{}
}

func (b BetaBuilder) GetBuilderInfo() BuilderInfo {
	return BuilderInfo{
		Consumes: []IData{
			AlphaProduces{},
		},
		Produces:  BetaProduces{},
		Optionals: nil,
		Accesses:  nil,
	}
}

type Terminator struct {
}

func (t Terminator) Process(context BuilderContext) IData {
	return WorkflowTerminated{}
}

func (t Terminator) GetBuilderInfo() BuilderInfo {
	return BuilderInfo{
		Consumes: []IData{
			AlphaProduces{},
			BetaProduces{},
		},
		Produces:  WorkflowTerminated{},
		Optionals: nil,
		Accesses:  nil,
	}
}

type TestWorkflow struct {
}

func (tW TestWorkflow) GetWorkflowMeta() WorkflowMeta {
	return WorkflowMeta{
		Builders: []IBuilder{
			AlphaBuilder{},
			BetaBuilder{},
			Terminator{},
		},
		TargetData: WorkflowTerminated{},
	}
}

type RedundantBuilderFailureWorklow struct {
}

func (rBFW RedundantBuilderFailureWorklow) GetWorkflowMeta() WorkflowMeta {
	return WorkflowMeta{
		Builders: []IBuilder{
			AlphaBuilder{},
			AlphaBuilder{},
			Terminator{},
		},
		TargetData: WorkflowTerminated{},
	}
}

type EmptyTargetDataWorkflow struct {
}

func (eTDW EmptyTargetDataWorkflow) GetWorkflowMeta() WorkflowMeta {
	return WorkflowMeta{
		Builders: []IBuilder{
			AlphaBuilder{},
			BetaBuilder{},
			Terminator{},
		},
	}
}

type MockStorage struct {
	store map[string]interface{}
}

func (ms *MockStorage) Read(key string) (interface{}, bool) {
	val, ok := ms.store[key]
	return val, ok
}

func (ms *MockStorage) Write(key string, val interface{}) {
	ms.store[key] = val
}

func TestRegisterWorkflow(t *testing.T) {
	type args struct {
		workflowKey string
		workflow    IWorkflow
	}
	tests := []struct {
		name               string
		args               args
		wantErr            bool
		skipHierarchyCheck bool
		wantHierarchy      [][]BuilderMeta
	}{
		{
			name: "Base_Successful_Test",
			args: args{
				workflowKey: "testWorkflow",
				workflow:    TestWorkflow{},
			},
			wantErr: false,
			wantHierarchy: [][]BuilderMeta{
				{
					{
						Name:      Name(AlphaBuilder{}),
						Consumes:  mapset.NewSet[string](Name(AlphaConsumes{})),
						Produces:  Name(AlphaProduces{}),
						Optionals: mapset.NewSet[string](),
						Accesses:  mapset.NewSet[string](),
					},
				},
				{
					{
						Name:      Name(BetaBuilder{}),
						Consumes:  mapset.NewSet[string](Name(AlphaProduces{})),
						Produces:  Name(BetaProduces{}),
						Optionals: mapset.NewSet[string](),
						Accesses:  mapset.NewSet[string](),
					},
				},
				{
					{
						Name:      Name(Terminator{}),
						Consumes:  mapset.NewSet[string](Name(AlphaProduces{}), Name(BetaProduces{})),
						Produces:  Name(WorkflowTerminated{}),
						Optionals: mapset.NewSet[string](),
						Accesses:  mapset.NewSet[string](),
					},
				},
			},
		},
		{
			name: "Redundant_Builder_Failure_Test",
			args: args{
				workflowKey: "redundantBuilderWorkflow",
				workflow:    RedundantBuilderFailureWorklow{},
			},
			wantErr:            true,
			skipHierarchyCheck: true,
		},
		{
			name: "Empty_Target_Data_Failure_Test",
			args: args{
				workflowKey: "emptyTargetDataWorkflow",
				workflow:    EmptyTargetDataWorkflow{},
			},
			wantErr:            true,
			skipHierarchyCheck: true,
		},
	}
	mockStorage := &MockStorage{
		store: make(map[string]interface{}),
	}
	InitRegistry(mockStorage)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RegisterWorkflow(tt.args.workflowKey, tt.args.workflow); (err != nil) != tt.wantErr {
				t.Errorf("RegisterWorkflow() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.skipHierarchyCheck {
				return
			}

			dataFlowInterface, _ := mockStorage.Read(tt.args.workflowKey)
			dataFlow := dataFlowInterface.(DataFlow)
			gotHierarchy := dataFlow.DependencyHierarchy
			wantHierarchy := tt.wantHierarchy

			compareHierarchies(wantHierarchy, gotHierarchy, t)
		})
	}
}

func compareHierarchies(wantHierarchy [][]BuilderMeta, gotHierarchy [][]BuilderMeta, t *testing.T) {
	for level := range gotHierarchy {
		for index, builder := range gotHierarchy[level] {
			if wantHierarchy[level][index].Name != builder.Name {
				t.Errorf("Builder Name %s does not match %s at %d %d", wantHierarchy[level][index].Name, builder.Name, level, index)
			}
			if !wantHierarchy[level][index].Consumes.Equal(builder.Consumes) {
				t.Errorf("Builder Consumes %s does not match %s at %d %d", wantHierarchy[level][index].Consumes, builder.Consumes, level, index)
			}
			if !wantHierarchy[level][index].Accesses.Equal(builder.Accesses) {
				t.Errorf("Builder Accesses %s does not match %s at %d %d", wantHierarchy[level][index].Accesses, builder.Accesses, level, index)
			}
			if !wantHierarchy[level][index].Optionals.Equal(builder.Optionals) {
				t.Errorf("Builder Optionals %s does not match %s at %d %d", wantHierarchy[level][index].Optionals, builder.Optionals, level, index)
			}
			if wantHierarchy[level][index].Produces != builder.Produces {
				t.Errorf("Builder Produces %s does not match %s at %d %d", wantHierarchy[level][index].Produces, builder.Produces, level, index)
			}
		}
	}
}
