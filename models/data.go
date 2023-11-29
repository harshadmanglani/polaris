package models

import "reflect"

type IData interface {
	GetDataMeta() DataMeta
}

type DataMeta struct {
	Data        string
	GeneratedBy string
}

type DataDelta struct {
	Delta []IData
}

type DataSet struct {
	AvailableData map[string]IData
}

func DataArrayToMap(data ...interface{}) map[string]bool {
	names := make(map[string]bool, len(data))
	for _, d := range data {
		names[DataToString(d)] = true
	}
	return names
}

func DataToString(data interface{}) string {
	name := reflect.TypeOf(data).String()
	return name
}
