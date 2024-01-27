package polaris

import (
	"reflect"

	mapset "github.com/deckarep/golang-set/v2"
)

type IBuilder interface {
	GetBuilderInfo() BuilderInfo
	Process(BuilderContext) IData
}

type BuilderContext struct {
	DataSet     DataSet
	ContextData map[string]any
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
	PtrType   reflect.Type
}

func (meta *BuilderMeta) DeepCopy() BuilderMeta {
	return BuilderMeta{
		Consumes:  meta.Consumes,
		Optionals: meta.Optionals,
		Accesses:  meta.Accesses,
		Produces:  meta.Produces,
		Name:      meta.Name,
		Rank:      meta.Rank,
	}
}

func (meta *BuilderMeta) EffectiveConsumes() mapset.Set[string] {
	if meta.Optionals != nil {
		return meta.Consumes.Union(meta.Optionals)
	}
	return meta.Consumes
}
