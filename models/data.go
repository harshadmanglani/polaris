package models

import (
	"reflect"
	"strings"
)

type IData interface {
}

// type DataMeta struct {
// 	Data        string
// 	GeneratedBy string
// }

type DataDelta struct {
	Delta []IData
}

type DataSet struct {
	AvailableData map[string]IData
}

func StructsToMap(data ...interface{}) map[string]bool {
	names := make(map[string]bool, len(data))
	for _, d := range data {
		names[Name(d)] = true
	}
	return names
}

var structToNameMapping = make(map[reflect.Type]string)

func Name(strucc interface{}) string {
	t := reflect.TypeOf(strucc)

	if name, ok := structToNameMapping[t]; ok {
		return strings.ToUpper(name)
	}

	name := strings.ToUpper(camelToSnake(t.Name()))
	structToNameMapping[t] = name

	return name
}

func camelToSnake(s string) string {
	var result []rune
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result = append(result, '_')
		}
		result = append(result, r)
	}
	return strings.ToLower(string(result))
}
