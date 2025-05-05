package polaris

type IDataStore interface {
	WriteDataFlow(key string, dataFlow *DataFlow)
	ReadDataFlow(key string) (*DataFlow, bool)
	WriteDataSet(key string, dataSet *DataSet)
	ReadDataSet(key string) (*DataSet, bool)
}
