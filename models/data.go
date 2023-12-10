package models

import (
	"reflect"
	"strings"
)

type IData interface {
}

type DataDelta struct {
	Delta []IData
}

type DataSet struct {
	AvailableData map[string]IData
}

type DataExecutionResponse struct {
	Responses map[string]IData
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
