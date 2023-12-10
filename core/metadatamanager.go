package core

import (
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/harshadmanglani/polaris/models"
)

type MetaDataManager struct {
	builders              map[string]models.IBuilder
	builderMetaMap        map[string]models.BuilderMeta
	producedToProducerMap map[string][]models.BuilderMeta
	consumesMeta          map[string]mapset.Set[models.BuilderMeta]
	optionalsMeta         map[string]mapset.Set[models.BuilderMeta]
	accessesMeta          map[string]mapset.Set[models.BuilderMeta]
}

func buildSet(data []models.IData) mapset.Set[string] {
	set := mapset.NewSet[string]()
	for _, d := range data {
		set.Add(models.Name(d))
	}
	return set
}

func newBuilderMeta(builder models.IBuilder) models.BuilderMeta {
	builderInfo := builder.GetBuilderInfo()
	return models.BuilderMeta{
		Consumes:  buildSet(builderInfo.Consumes),
		Optionals: buildSet(builderInfo.Optionals),
		Accesses:  buildSet(builderInfo.Accesses),
		Produces:  models.Name(builderInfo.Produces),
		Name:      models.Name(builder),
	}
}

func (m *MetaDataManager) register(builder models.IBuilder) {
	builderMeta := newBuilderMeta(builder)
	if _, ok := m.builderMetaMap[builderMeta.Name]; ok {
		panic("Builder already exists")
	}
	m.builders[builderMeta.Name] = builder
	m.builderMetaMap[builderMeta.Name] = builderMeta
	m.producedToProducerMap[builderMeta.Produces] = append(m.producedToProducerMap[builderMeta.Produces], builderMeta)
	for _, c := range builder.GetBuilderInfo().Consumes {
		m.consumesMeta[models.Name(c)].Add(builderMeta)
	}
	for _, o := range builder.GetBuilderInfo().Optionals {
		m.optionalsMeta[models.Name(o)].Add(builderMeta)
	}
	for _, a := range builder.GetBuilderInfo().Accesses {
		m.accessesMeta[models.Name(a)].Add(builderMeta)
	}
}

func (m *MetaDataManager) getMetaForProducerOf(data string) []models.BuilderMeta {
	return m.producedToProducerMap[data]
}
