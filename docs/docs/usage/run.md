---
sidebar_position: 3
---
# Running Workflows

```go
polaris.InitRegistry(dataStore)
polaris.RegisterWorkflow(workflowKey, workflow)

executor := polaris.Executor{
	Before: func(builder reflect.Type, delta []IData) {
        fmt.Printf("Builder %s is about to be run with new data %v\n", builder, delta)
    }
	After: func(builder reflect.Type, produced IData) {
        fmt.Printf("Builder %s produced %s\n", builder, produced)
    }
}

response, err := executor.Sequential(workflowKey, workflowId, dataDelta)

response, err := executor.Parallel(workflowKey, workflowId, dataDelta)
```