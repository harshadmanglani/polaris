---
sidebar_position: 3
---
# Builder

An actor that consumes a bunch of data and produces another data. It has the following meta associated with it:
    * **Name** - Name of the builder.
    * **Consumes** - A set of data that the builder consumes.
    * **Produces** - Data that the builder produces.
    * **Optionals** - Data that the builder can optionally consume; one possible use case for this: if a builder wants to be re-run on demand with the same set of consumable data already present, add an optional data in the Builder and restart the workflow by passing an instance of the optional data in the data-delta
    * **Access** - Data that the builder will just access and has no effect on the sequence of execution of builders
    * **BuilderContext** - A wrapper to access the data given to that builder to process.

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
  // save the request in a database (different from Polaris storing workflows in `IDataStore`)
  database.save(userInitReq)

  // call another service to place a request, and wait for the response
  cabbieResponse := cabbieHttpClient.request(RideRequest{
    userId: userInitReq.userId,
    source: userInitReq.source,
    dest: userInitReq.dest
  })

  // once done, return the `Produces` of the data
  return UserInitiationResponse{
    success: true,
    etaForCabbie: cabbieResponse.eta
  }
}
```
