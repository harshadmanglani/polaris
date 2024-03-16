---
sidebar_position: 5
---
# Data Store

Workflows need to be stored to a database. Whether you're using a key-value store or RDBMS, you need to implement the `IDataStore` interface and pass it to the `polaris.InitRegistry` method.

```go
type IDataStore interface {
  Write(key string, value interface{})
  Read(key string) (interface{}, bool)
}
```

Once the interface is implemented, data stores must be set with polaris.
```go
polaris.InitRegistry(dataStore)
```

## In-memory data store
**DO NOT DO THIS IN PRODUCTION!**

While getting set up with polaris, you can start with a simple in-memory data store
```go
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

mockStorage := &mockStorage{
  store: make(map[string]interface{}),
}

polaris.InitRegistry(dataStore)
```