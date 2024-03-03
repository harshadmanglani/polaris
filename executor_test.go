package polaris

import (
	"reflect"
	"testing"
)

func TestExecutor_Sequential(t *testing.T) {
	type fields struct {
		Before func(builder reflect.Type, delta []IData)
		After  func(builder reflect.Type, produced IData)
	}
	type args struct {
		workflowKey string
		workflowId  string
		data        []IData
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		workflowKey string
		workflow    IWorkflow
		want        DataExecutionResponse
		wantErr     bool
	}{
		{
			name: "Base_Successful_Test",
			args: args{
				workflowKey: "testWorkflow",
				workflowId:  "1",
				data: []IData{
					alphaConsumes{},
				},
			},
			want: DataExecutionResponse{
				Responses: map[string]IData{
					Name(alphaProduces{}):      alphaProduces{},
					Name(betaProduces{}):       betaProduces{},
					Name(workflowTerminated{}): workflowTerminated{},
				},
			},
			workflowKey: "testWorkflow",
			workflow:    testWorkflow{},
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockStorage := &mockStorage{
				store: make(map[string]interface{}),
			}
			InitRegistry(mockStorage)
			RegisterWorkflow(tt.workflowKey, tt.workflow)
			e := &Executor{
				Before: tt.fields.Before,
				After:  tt.fields.After,
			}
			got, err := e.Sequential(tt.args.workflowKey, tt.args.workflowId, tt.args.data...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Executor.Sequential() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Executor.Sequential() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExecutor_Parallel(t *testing.T) {
	type fields struct {
		Before func(builder reflect.Type, delta []IData)
		After  func(builder reflect.Type, produced IData)
	}
	type args struct {
		workflowKey string
		workflowId  string
		data        []IData
	}
	tests := []struct {
		name        string
		fields      fields
		workflowKey string
		workflow    IWorkflow
		args        args
		want        DataExecutionResponse
		wantErr     bool
	}{
		{
			name: "Base_Successful_Test",
			args: args{
				workflowKey: "testWorkflow",
				workflowId:  "1",
				data: []IData{
					alphaConsumes{},
				},
			},
			want: DataExecutionResponse{
				Responses: map[string]IData{
					Name(alphaProduces{}):      alphaProduces{},
					Name(betaProduces{}):       betaProduces{},
					Name(workflowTerminated{}): workflowTerminated{},
				},
			},
			workflowKey: "testWorkflow",
			workflow:    testWorkflow{},
			wantErr:     false,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockStorage := &mockStorage{
				store: make(map[string]interface{}),
			}
			InitRegistry(mockStorage)
			RegisterWorkflow(tt.workflowKey, tt.workflow)
			e := &Executor{
				Before: tt.fields.Before,
				After:  tt.fields.After,
			}
			got, err := e.Parallel(tt.args.workflowKey, tt.args.workflowId, tt.args.data...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Executor.Parallel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Executor.Parallel() = %v, want %v", got, tt.want)
			}
		})
	}
}
