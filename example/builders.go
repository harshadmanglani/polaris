package main

import (
	"fmt"

	"github.com/harshadmanglani/polaris/models"
)

type Validator struct {
}

func (v Validator) Process(context models.BuilderContext) models.IData {
	fmt.Println("In validator")
	// read from db
	// do some inventory checks, acquire locks (or perform the Auth phase for Auth Capture)
	// throw an error if checks fail
	return RequestValidated{}
}

func (v Validator) GetBuilderInfo() models.BuilderInfo {
	return models.BuilderInfo{
		Consumes: []models.IData{
			OrderRequest{},
		},
		Produces:  RequestValidated{},
		Optionals: nil,
		Accesses:  nil,
	}
}

type RiskChecker1 struct {
}

func (rC RiskChecker1) Process(context models.BuilderContext) models.IData {
	fmt.Println("In risk checker1")
	// read from db
	// do some risk checks with a downstream service
	// throw an error if checks fail
	return RiskCheck1Completed{}
}

func (rC RiskChecker1) GetBuilderInfo() models.BuilderInfo {
	return models.BuilderInfo{
		Consumes: []models.IData{
			OrderRequest{},
			RequestValidated{},
		},
		Produces:  RiskCheck1Completed{},
		Optionals: nil,
		Accesses:  nil,
	}
}

type RiskChecker2 struct {
}

func (rC RiskChecker2) Process(context models.BuilderContext) models.IData {
	fmt.Println("In risk checker2")
	// read from db
	// do some risk checks with a downstream service
	// throw an error if checks fail
	return RiskCheck1Completed{}
}

func (rC RiskChecker2) GetBuilderInfo() models.BuilderInfo {
	return models.BuilderInfo{
		Consumes: []models.IData{
			OrderRequest{},
			RequestValidated{},
		},
		Produces:  RiskCheck2Completed{},
		Optionals: nil,
		Accesses:  nil,
	}
}

type Persistor struct {
}

func (p Persistor) Process(context models.BuilderContext) models.IData {
	fmt.Println("In persistor")
	// store order in db
	// release locks and reduce stock count (or perform the Capture phase)
	return OrderInfo{}
}

func (p Persistor) GetBuilderInfo() models.BuilderInfo {
	return models.BuilderInfo{
		Consumes: []models.IData{
			OrderRequest{},
			RiskCheck1Completed{},
			RiskCheck2Completed{},
		},
		Produces:  OrderInfo{},
		Optionals: nil,
		Accesses:  nil,
	}
}

type PaymentInitializer struct {
}

func (pI PaymentInitializer) Process(context models.BuilderContext) models.IData {
	fmt.Println("In payment initializer")
	// call a PG
	return PaymentInitialized{}
}

func (pI PaymentInitializer) GetBuilderInfo() models.BuilderInfo {
	return models.BuilderInfo{
		Consumes: []models.IData{
			OrderInfo{},
		},
		Produces:  PaymentInitialized{},
		Optionals: nil,
		Accesses:  nil,
	}
}

type PaymentStatusUpdater struct {
}

func (pS PaymentStatusUpdater) Process(context models.BuilderContext) models.IData {
	fmt.Println("In payment status updater")
	// schedule warehouse ops if payment is successful
	return WarehouseOpsScheduled{}
}

func (pS PaymentStatusUpdater) GetBuilderInfo() models.BuilderInfo {
	return models.BuilderInfo{
		Consumes: []models.IData{
			PaymentInitialized{},
			PaymentStatus{},
		},
		Produces:  WarehouseOpsScheduled{},
		Optionals: nil,
		Accesses:  nil,
	}
}

type WarehouseStatusUpdater struct {
}

func (wS WarehouseStatusUpdater) Process(context models.BuilderContext) models.IData {
	fmt.Println("In warehouse status updater")
	// if order was delivered produce output, otherwise produce null and flow will stop
	return OrderDelivered{}
}

func (wS WarehouseStatusUpdater) GetBuilderInfo() models.BuilderInfo {
	return models.BuilderInfo{
		Consumes: []models.IData{
			WarehouseStatus{},
		},
		Produces:  OrderDelivered{},
		Optionals: nil,
		Accesses:  nil,
	}
}

type WorkflowTerminator struct {
}

func (wT WorkflowTerminator) Process(context models.BuilderContext) models.IData {
	fmt.Println("In workflow terminator")
	// schedule warehouse ops if payment is successful
	return WarehouseOpsScheduled{}
}

func (wT WorkflowTerminator) GetBuilderInfo() models.BuilderInfo {
	return models.BuilderInfo{
		Consumes: []models.IData{
			OrderDelivered{},
		},
		Produces:  WorkflowTerminated{},
		Optionals: nil,
		Accesses:  nil,
	}
}
