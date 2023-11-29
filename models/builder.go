package models

type IBuilder interface {
	GetBuilderMeta() BuilderMeta
	Process() IData
}

type BuilderMeta struct {
	Consumes  map[string]bool
	Optionals map[string]bool
	Accesses  map[string]bool
	Produces  string
	Name      string
	Rank      int
}

func CreateNew(consumes []IData,
	optionals []IData,
	accesses []IData,
	produces IData,
	builder IBuilder) BuilderMeta {
	return BuilderMeta{
		Consumes:  DataArrayToMap(consumes),
		Optionals: DataArrayToMap(optionals),
		Accesses:  DataArrayToMap(accesses),
		Produces:  DataToString(produces),
		Name:      DataToString(builder),
	}
}
