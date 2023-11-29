package core

import "github.com/harshadmanglani/polaris/models"

type DataSetAccessor struct {
	DataSet models.DataSet
}

func createNew(dataSet models.DataSet) DataSetAccessor {
	return DataSetAccessor{DataSet: dataSet}
}

func (dsa DataSetAccessor) add(data models.IData) models.DataSet {
	dsa.DataSet.AvailableData[models.DataToString(data)] = data
	return dsa.DataSet
}

func (dsa DataSetAccessor) get(data models.IData) any {
	return dsa.DataSet.AvailableData[models.DataToString(data)]
}

func (dsa DataSetAccessor) getAccessibleData(data models.IData, builder models.IBuilder) any {
	return dsa.DataSet.AvailableData[models.DataToString(data)]
}
