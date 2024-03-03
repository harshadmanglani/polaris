package polaris

import (
	"testing"

	mapset "github.com/deckarep/golang-set/v2"
)

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
				workflow:    testWorkflow{},
			},
			wantErr: false,
			wantHierarchy: [][]BuilderMeta{
				{
					{
						Name:      Name(alphaBuilder{}),
						Consumes:  mapset.NewSet[string](Name(alphaConsumes{})),
						Produces:  Name(alphaProduces{}),
						Optionals: mapset.NewSet[string](),
						Accesses:  mapset.NewSet[string](),
					},
				},
				{
					{
						Name:      Name(betaBuilder{}),
						Consumes:  mapset.NewSet[string](Name(alphaProduces{})),
						Produces:  Name(betaProduces{}),
						Optionals: mapset.NewSet[string](),
						Accesses:  mapset.NewSet[string](),
					},
				},
				{
					{
						Name:      Name(terminator{}),
						Consumes:  mapset.NewSet[string](Name(alphaProduces{}), Name(betaProduces{})),
						Produces:  Name(workflowTerminated{}),
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
				workflow:    redundantBuilderFailureWorklow{},
			},
			wantErr:            true,
			skipHierarchyCheck: true,
		},
		{
			name: "Empty_Target_Data_Failure_Test",
			args: args{
				workflowKey: "emptyTargetDataWorkflow",
				workflow:    emptyTargetDataWorkflow{},
			},
			wantErr:            true,
			skipHierarchyCheck: true,
		},
	}
	mockStorage := &mockStorage{
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
