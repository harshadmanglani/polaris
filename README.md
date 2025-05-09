<div align="center">
  <a href="https://harshadmanglani.github.io/polaris">
    <img src="docs/docs/assets/polaris-header-dark.svg"/>
  </a>

  <p>
     An extremely light weight workflow orchestrator for Golang.
  </p>
</div>
<br/>
Polaris is an advanced workflow orchestration engine designed to help you create, store, and run complex workflows efficiently. Whether your workflows span multiple request scopes or are self-contained, Polaris can handle them all. This documentation will guide you through setting up Polaris in your Go application, creating your first workflow, and understanding the key concepts behind its operation.

## Getting Started with Polaris
Install Go
Ensure you have Go installed on your system. You can check your version using the following command:

```bash
go version
```
This should return your installed Go version. For these tutorials, we used Go 1.21.5.

Install Polaris
If you are creating a new project using Polaris, start by creating a new directory and switching to it:

```bash
mkdir polaris-project
cd polaris-project
```

Initialize a Go project in that directory:

```bash
go mod init polaris-project
```

Finally, install Polaris with the following command:

```bash
go get github.com/harshadmanglani/polaris
```
## Polaris Concepts
### What is Polaris?
Polaris helps you create, store, and run workflows efficiently. Key features include:

Automatic sequence determination for steps.
Implicit pausing when out of new data.
Support for Redis, Postgres, and timeouts at both Builder and Workflow levels.
Workflows and Builders
A workflow is a series of multiple steps that can be executed sequentially or in parallel. Builders are the units of work within a workflow. They must implement the IBuilder interface.

```go
type IBuilder interface {
    GetBuilderInfo() BuilderInfo
    Process(context BuilderContext) IData
}
```
For example, for a cab ride workflow, builders might include:

- UserInitiation
- CabbieMatching
- CabDepartureFromSource
- CabArrivalAtDest
- UserPayment
- RideEnds
- Setting Up Your Go App
- Service Startup
- Implement the IDataStore interface and initialize - Polaris with it. Register your workflow(s) with Polaris, and then initialize your executor:

```go
import (
    "github.com/harshadmanglani/polaris"
)

func main() {
    // Implement IDataStore interface and initialize Polaris
    polaris.Initialize("your-data-store")
    
    // Register your workflow(s) with Polaris
    polaris.RegisterWorkflow(CabRideWorkflow{})
    
    // Initialize your executor
    executor := polaris.Executor{}
}
```
Runtime Operations
At runtime, you must:

Accept the request.
Generate a unique identifier for this request (uniqueWorkflowId).
Pass your request data (ensuring it is part of the Consumes for the Builder you want to run).
Creating Workflows
Create Your First Workflow
For a cab ride workflow, builders could include:

```go
type CabRideWorkflow struct {}
func (cr CabRideWorkflow) GetWorkflowMeta() WorkflowMeta {
    return WorkflowMeta{
        Builders: []IBuilder{
            UserInitiation{},
            CabbieMatching{},
            // Add other builders as needed
        },
        TargetData: WorkflowTerminated{},
    }
}
```

You don't have to sequentially define the builders in order of execution. Polaris will figure it out. However, you should if you can for readability.

### Deployment and Versioning
Workflow Versioning
Deploying new code requires backward compatibility unless 100% downtime is guaranteed. This involves:

Deploying a version of code that is backward compatible for older non-terminal workflows while newer ones execute on the new code.
Once older workflows complete, deploy to clean up stale code.

  <div align="center">
  <p>
    <a href="https://harshadmanglani.github.io/polaris/get-started/">Installation</a> | <a href="https://harshadmanglani.github.io/polaris/usage/">Documentation</a> | <a href="https://x.com/polaris_golang">Twitter</a> | <a href="https://github.com/flipkart-incubator/databuilderframework">Inspiration</a>
  </p>
</div>
