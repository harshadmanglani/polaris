package polaris

type IDataStore interface {
	Write(key string, value interface{})
	Read(key string) (interface{}, bool)
}
