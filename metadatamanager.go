package polaris

import (
	mapset "github.com/deckarep/golang-set/v2"
)

type MetaDataManager struct {
	builders              map[string]IBuilder
	builderMetaMap        map[string]BuilderMeta
	producedToProducerMap map[string]BuilderMeta // TODO: make this []BuilderMeta when ResolutionSpec is implemented
	// When are the below used?
	consumesMeta  map[string]mapset.Set[BuilderMeta]
	optionalsMeta map[string]mapset.Set[BuilderMeta]
	accessesMeta  map[string]mapset.Set[BuilderMeta]
}

func buildSet(data []IData) mapset.Set[string] {
	set := mapset.NewSet[string]()
	for _, d := range data {
		set.Add(Name(d))
	}
	return set
}

func newBuilderMeta(builder IBuilder) BuilderMeta {
	builderInfo := builder.GetBuilderInfo()
	return BuilderMeta{
		Consumes:  buildSet(builderInfo.Consumes),
		Optionals: buildSet(builderInfo.Optionals),
		Accesses:  buildSet(builderInfo.Accesses),
		Produces:  Name(builderInfo.Produces),
		Name:      Name(builder),
	}
}

func (m *MetaDataManager) register(builder IBuilder) {
	builderMeta := newBuilderMeta(builder)
	if _, ok := m.builderMetaMap[builderMeta.Name]; ok {
		panic("Builder already exists")
	}
	m.builders[builderMeta.Name] = builder
	m.builderMetaMap[builderMeta.Name] = builderMeta
	m.producedToProducerMap[builderMeta.Produces] = m.producedToProducerMap[builderMeta.Produces]
	for _, c := range builder.GetBuilderInfo().Consumes {
		m.consumesMeta[Name(c)].Add(builderMeta)
	}
	for _, o := range builder.GetBuilderInfo().Optionals {
		m.optionalsMeta[Name(o)].Add(builderMeta)
	}
	for _, a := range builder.GetBuilderInfo().Accesses {
		m.accessesMeta[Name(a)].Add(builderMeta)
	}
}

func (m *MetaDataManager) getMetaForProducerOf(data string) BuilderMeta {
	return m.producedToProducerMap[data]
}
