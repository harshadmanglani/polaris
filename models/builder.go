package models

type IBuilder interface {
	GetBuilderMeta() BuilderMeta
	Process(BuilderContext) IData
}

type BuilderMeta struct {
	Consumes  map[string]bool
	Optionals map[string]bool
	Accesses  map[string]bool
	Produces  string
	Name      string
	Rank      int
}
