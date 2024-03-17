---
sidebar_position: 4
---
# Usage

From the root of your Go module, run:
```
go get github.com/harshadmanglani/polaris
```
## Adding Polaris to your Go app

Assuming you've already read up on <a href="/polaris/concepts/polaris">concepts</a>, let's get started

We'll break integration in two parts:
- Things to do when your Go app is starting up
- Things to do at runtime

### Service Startup

This is where you want to 
1. Implement the <a href="/polaris/concepts/datastore">`IDataStore`</a> interface and initialize Polaris with it
2. Register your <a href="/polaris/concepts/workflow">workflow(s)</a> with Polaris
3. Initialize your executor

```go
var dataStore polaris.IDataStore
var executor polaris.Executor

void init(){
  dataStore := SomeDataStoreImpl{}
  polaris.InitRegistry(dataStore)

  polaris.RegisterWorkflow("alphaWorkflowKey", AlphaWorkflow{})

  executor := polaris.Executor{
    Before: func(builder reflect.Type, delta []IData) {
      fmt.Printf("Builder %s is about to be run with new data %v\n", builder, delta)
    }
    After: func(builder reflect.Type, produced IData) {
      fmt.Printf("Builder %s produced %s\n", builder, produced)
    }
  }
}
```
### Runtime

This is where you process requests on your server by handing them over to polaris.

At runtime, you must
1. Accept the request
2. Generate a unique identifier for this request (`uniqueWorkflowId`)
3. Pass your request data (ensuring that it would be part of the <a href="/polaris/concepts/builder#:~:text=of%20the%20builder.-,Consumes,-%2D%20A%20set%20of">`Consumes`</a> for the <a href="/polaris/concepts/builder">`Builder`</a> you want to run)

```go
void main(){
    http.HandleFunc("/request", RequestHandler)

    fmt.Println("Server running at port 8080...")
    http.ListenAndServe(":8080", nil)
}

func RequestHandler(w http.ResponseWriter, r *http.Request) {
    var alpha AlphaConsumes
    err := json.NewDecoder(r.Body).Decode(&alpha)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    uniqueWorkflowId := alpha.Id
    result := executor.Sequential("alphaWorkflowKey", uniqueWorkflowId, alpha)

    // fetch expected data from the result
    alphaProcessed, ok := result.Get(AlphaProcessed{})
    if !ok {
        w.WriteHeader(http.StatusInternalServerError)
    }

    // do something with alphaProcessed
    w.WriteHeader(http.StatusOK)
}
```

## Limitations
### Workflow versioning
1. Unless you can afford a 100% downtime ensuring all active workflows move into a terminal state, deploying new code requires ensuring backward compatibility.
2. What this means is - you'll need to a deploy a version of code that is backward compatible for older non terminal workflows while newer ones will execute on the new code.
3. Once the older workflows have completed, a deployment to clean up stale code will be required.

## How does the framework perform at scale?
The framework itself has extremely low overhead. Since execution graphs are generated pre-runtime, all the orchestrator will do at runtime is use the graph and available data to run whichever builders can be run. 

## Use cases
1. You have multi-step workflow executions where each step is dependent on data generated from previous steps.
2. Executions can span one request scope or multiple scopes.
3. Your systems works with reusable components that can be combined in different ways to generate different end-results.
4. Your workflows can pause, resume or even restart from the beginning.
