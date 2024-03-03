package polaris

import (
	"log"
	"reflect"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"go.uber.org/zap"
)

var sugar *zap.SugaredLogger

func init() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}

	sugar = logger.Sugar()
	defer logger.Sync()
}

type IWorkflow interface {
	GetWorkflowMeta() WorkflowMeta
}

type WorkflowMeta struct {
	Builders   []IBuilder
	TargetData IData
}

type IData interface {
}

type DataSet struct {
	AvailableData map[string]IData
}

type DataExecutionResponse struct {
	Responses map[string]IData
	Error     error
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

type IBuilder interface {
	GetBuilderInfo() BuilderInfo
	Process(BuilderContext) IData
}

type BuilderContext struct {
	DataSet DataSet
}

type BuilderInfo struct {
	Consumes  []IData
	Optionals []IData
	Accesses  []IData
	Produces  IData
}

type BuilderMeta struct {
	Consumes  mapset.Set[string]
	Optionals mapset.Set[string]
	Accesses  mapset.Set[string]
	Produces  string
	Name      string
	Rank      int
	Type      reflect.Type
}

func (meta *BuilderMeta) EffectiveConsumes() mapset.Set[string] {
	if meta.Optionals != nil {
		return meta.Consumes.Union(meta.Optionals)
	}
	return meta.Consumes
}
