---
sidebar_position: 3
---
# Usage

## Creating Workflows

### Defining a Workflow
A Workflow is a sequence of Builders that will perform some work. Let's take the example of a cab ride workflow. Essentially, for a cab ride workflow, builders (units of work) could be:
    - User initiating a request
    - Cabbie match
    - Cabbie reaches source
    - Ride starts
    - Cabbie reaches destination
    - User makes payment
    - Ride ends


Workflows must implement the `IWorkflow` interface.
```go
type IWorkflow interface {
	GetWorkflowMeta() WorkflowMeta
}
```
```go
type CabRideWorkflow struct {
}

func (cr CabRideWorkflow) GetWorkflowMeta() WorkflowMeta {
	return WorkflowMeta{
		Builders: []IBuilder{
                    UserInitiation{},
                    CabbieMatching{},
                    CabbieArrivalAtSource{},
                    CabDepartureFromSource{},
                    CabArrivalAtDest{},
                    UserPayment{},
                    RideEnds{}
		},
		TargetData: WorkflowTerminated{},
	}
}
```
You don't have to sequentially define the builders in order of execution. Polaris will figure it out. However, **you should if you can. It helps readability.**

### Defining a Builder

A Builder is a unit of work in the workflow. Builders must implement the `IBuilder` interface.

```go
type IBuilder interface {
	GetBuilderInfo() BuilderInfo
	Process(BuilderContext) IData
}
```

Following the same example, for the first unit of work in a cab ride workflow:
```go
var database Database
var cabbieHttpClient CabbieHttpClient 

type UserInitiation struct {
}

func (uI UserInitiation) GetBuilderInfo() BuilderInfo {
    return BuilderInfo{
        Consumes: []IData{
            UserInitiationRequest{},
        },
        Produces:  UserInitiationResponse{},
        Optionals: nil,
        Accesses:  nil,
    }
}

func (uI UserInitiation) Process(context BuilderContext) IData {
    userInitReq := context.get(UserInitiationRequest{})
    database.save(userInitReq)

    cabbieResponse := cabbieHttpClient.request(RideRequest{
        userId: userInitReq.userId,
        source: userInitReq.source,
        dest: userInitReq.dest
    })

    return UserInitiationResponse{
        success: true,
        etaForCabbie: cabbieResponse.eta
    }
}
```

### Defining a Data

A Data is a struct that holds the data that will be consumed and/or produced by steps in your workflow.
These objects must implement the `IData` interface, which basically means that they should be a `struct`.

For a user initiating a cab ride request, this is what the initial `Data` might look like.

```go
type UserInitiationRequest struct{
    userId string
    source string
    dest string
}
```

## Storing Workflows

Workflows need to be stored to a database. Whether you're using a key-value store or RDBMS, you need to implement the `IDataStore` interface and pass it to the `polaris.InitRegistry` method.

```go
type IDataStore interface {
	Write(key string, value interface{})
	Read(key string) (interface{}, bool)
}
```

## Running Workflows

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