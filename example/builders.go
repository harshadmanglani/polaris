package example

import "github.com/harshadmanglani/polaris/models"

type Validator struct {
}

func (v Validator) Process(context models.BuilderContext) models.IData {
	// read from db
	// do some inventory checks
	// throw an error if checks fail
	return OrderRequestValidated{}
}

func (v Validator) GetBuilderMeta() models.BuilderMeta {
	consumes := [...]models.IData{OrderRequest{}}
	produces := OrderRequestValidated{}
	return models.BuilderMeta{
		Consumes:  models.StructsToMap(consumes),
		Produces:  models.Name(produces),
		Optionals: nil,
		Accesses:  nil,
		Name:      models.Name("RequestValidator"), // or models.Name(Validator{})
	}
}

type RiskChecker struct {
}

func (rc RiskChecker) Process(context models.BuilderContext) models.IData {
	// read from db
	// do some risk checks with a downstream service
	// throw an error if checks fail
	return RiskCheckCompleted{}
}

func (rc RiskChecker) GetBuilderMeta() models.BuilderMeta {
	consumes := [...]models.IData{OrderRequest{}, OrderRequestValidated{}}
	produces := RiskCheckCompleted{}
	return models.BuilderMeta{
		Consumes:  models.StructsToMap(consumes),
		Produces:  models.Name(produces),
		Optionals: nil,
		Accesses:  nil,
		Name:      models.Name("RiskChecker"), // or models.Name(Validator{})
	}
}
