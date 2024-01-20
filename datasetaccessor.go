package polaris

type DataSetAccessor struct {
	DataSet DataSet
}

func createNew(dataSet DataSet) DataSetAccessor {
	return DataSetAccessor{DataSet: dataSet}
}

func (dsa DataSetAccessor) add(data IData) DataSet {
	dsa.DataSet.AvailableData[Name(data)] = data
	return dsa.DataSet
}

func (dsa DataSetAccessor) get(data IData) any {
	return dsa.DataSet.AvailableData[Name(data)]
}

func (dsa DataSetAccessor) getAccessibleData(data IData, builder IBuilder) any {
	return dsa.DataSet.AvailableData[Name(data)]
}
