package main

import "github.com/harshadmanglani/polaris/models"

type Validator struct {
}

func (v Validator) Process(context models.BuilderContext) models.IData {
	// read from db
	// do some inventory checks, acquire locks (or perform the Auth phase for Auth Capture)
	// throw an error if checks fail
	return RequestValidated{}
}

func (v Validator) GetBuilderMeta() models.BuilderMeta {
	return models.BuilderMeta{
		Consumes: models.StructsToMap([]models.IData{
			OrderRequest{},
		}),
		Produces:  models.Name(RequestValidated{}),
		Optionals: nil,
		Accesses:  nil,
		Name:      models.Name(v),
	}
}

type RiskChecker struct {
}

func (rC RiskChecker) Process(context models.BuilderContext) models.IData {
	// read from db
	// do some risk checks with a downstream service
	// throw an error if checks fail
	return RiskCheckCompleted{}
}

func (rC RiskChecker) GetBuilderMeta() models.BuilderMeta {
	return models.BuilderMeta{
		Consumes: models.StructsToMap([]models.IData{
			OrderRequest{},
			RequestValidated{},
		}),
		Produces:  models.Name(RiskCheckCompleted{}),
		Optionals: nil,
		Accesses:  nil,
		Name:      models.Name(rC),
	}
}

type Persistor struct {
}

func (p Persistor) Process(context models.BuilderContext) models.IData {
	// store order in db
	// release locks and reduce stock count (or perform the Capture phase)
	return OrderInfo{}
}

func (p Persistor) GetBuilderMeta() models.BuilderMeta {
	return models.BuilderMeta{
		Consumes: models.StructsToMap([]models.IData{
			OrderRequest{},
			RiskCheckCompleted{},
		}),
		Produces:  models.Name(OrderInfo{}),
		Optionals: nil,
		Accesses:  nil,
		Name:      models.Name(p),
	}
}

type PaymentInitializer struct {
}

func (pI PaymentInitializer) Process(context models.BuilderContext) models.IData {
	// call a PG
	return PaymentInitialized{}
}

func (pI PaymentInitializer) GetBuilderMeta() models.BuilderMeta {
	return models.BuilderMeta{
		Consumes: models.StructsToMap([]models.IData{
			OrderInfo{},
		}),
		Produces:  models.Name(PaymentInitializer{}),
		Optionals: nil,
		Accesses:  nil,
		Name:      models.Name(pI),
	}
}

type PaymentStatusUpdater struct {
}

func (pS PaymentStatusUpdater) Process(context models.BuilderContext) models.IData {
	// schedule warehouse ops if payment is successful
	return WarehouseOpsScheduled{}
}

func (pS PaymentStatusUpdater) GetBuilderMeta() models.BuilderMeta {
	return models.BuilderMeta{
		Consumes: models.StructsToMap([]models.IData{
			PaymentInitialized{},
			PaymentStatus{},
		}),
		Produces:  models.Name(WarehouseOpsScheduled{}),
		Optionals: nil,
		Accesses:  nil,
		Name:      models.Name(pS),
	}
}

type WarehouseStatusUpdater struct {
}

func (wS WarehouseStatusUpdater) Process(context models.BuilderContext) models.IData {
	// schedule warehouse ops if payment is successful
	return WarehouseOpsScheduled{}
}

func (wS WarehouseStatusUpdater) GetBuilderMeta() models.BuilderMeta {
	return models.BuilderMeta{
		Consumes: models.StructsToMap([]models.IData{
			PaymentInitialized{},
			PaymentStatus{},
		}),
		Produces:  models.Name(WarehouseOpsScheduled{}),
		Optionals: nil,
		Accesses:  nil,
		Name:      models.Name(wS),
	}
}
