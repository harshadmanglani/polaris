package polaris

import "fmt"

type MetaDataManager struct {
	builders              map[string]IBuilder
	builderMetaMap        map[string]BuilderMeta
	producedToProducerMap map[string]BuilderMeta // TODO: make this []BuilderMeta when ResolutionSpec is implemented
}

func newMetaDataManager(workflow IWorkflow) MetaDataManager {
	metaDataManager := MetaDataManager{
		builders:              make(map[string]IBuilder),
		builderMetaMap:        make(map[string]BuilderMeta),
		producedToProducerMap: make(map[string]BuilderMeta),
	}

	for _, b := range workflow.GetWorkflowMeta().Builders {
		metaDataManager.registerBuilder(b)
	}
	return metaDataManager
}

func (m *MetaDataManager) registerBuilder(builder IBuilder) error {
	builderMeta := newBuilderMeta(builder)
	if _, ok := m.builderMetaMap[builderMeta.Name]; ok {
		sugar.Errorf("Builder %s already exists", builderMeta.Name)
		return fmt.Errorf("BUILDER_ALREADY_EXISTS")
	}
	m.builders[builderMeta.Name] = builder
	m.builderMetaMap[builderMeta.Name] = builderMeta
	m.producedToProducerMap[builderMeta.Produces] = builderMeta
	return nil
}
