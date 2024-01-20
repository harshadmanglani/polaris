package main

import (
	"fmt"

	"github.com/harshadmanglani/polaris"
)

type Validator struct {
}

func (v Validator) Process(context polaris.BuilderContext) polaris.IData {
	fmt.Println("In validator")
	// read from db
	// do some inventory checks, acquire locks (or perform the Auth phase for Auth Capture)
	// throw an error if checks fail
	return RequestValidated{}
}

func (v Validator) GetBuilderInfo() polaris.BuilderInfo {
	return polaris.BuilderInfo{
		Consumes: []polaris.IData{
			OrderRequest{},
		},
		Produces:  RequestValidated{},
		Optionals: nil,
		Accesses:  nil,
	}
}

type RiskChecker1 struct {
}

func (rC RiskChecker1) Process(context polaris.BuilderContext) polaris.IData {
	fmt.Println("In risk checker1")
	// read from db
	// do some risk checks with a downstream service
	// throw an error if checks fail
	return RiskCheck1Completed{}
}

func (rC RiskChecker1) GetBuilderInfo() polaris.BuilderInfo {
	return polaris.BuilderInfo{
		Consumes: []polaris.IData{
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

func (rC RiskChecker2) Process(context polaris.BuilderContext) polaris.IData {
	fmt.Println("In risk checker2")
	// read from db
	// do some risk checks with a downstream service
	// throw an error if checks fail
	return RiskCheck1Completed{}
}

func (rC RiskChecker2) GetBuilderInfo() polaris.BuilderInfo {
	return polaris.BuilderInfo{
		Consumes: []polaris.IData{
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

func (p Persistor) Process(context polaris.BuilderContext) polaris.IData {
	fmt.Println("In persistor")
	// store order in db
	// release locks and reduce stock count (or perform the Capture phase)
	return OrderInfo{}
}

func (p Persistor) GetBuilderInfo() polaris.BuilderInfo {
	return polaris.BuilderInfo{
		Consumes: []polaris.IData{
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

func (pI PaymentInitializer) Process(context polaris.BuilderContext) polaris.IData {
	fmt.Println("In payment initializer")
	// call a PG
	return PaymentInitialized{}
}

func (pI PaymentInitializer) GetBuilderInfo() polaris.BuilderInfo {
	return polaris.BuilderInfo{
		Consumes: []polaris.IData{
			OrderInfo{},
		},
		Produces:  PaymentInitialized{},
		Optionals: nil,
		Accesses:  nil,
	}
}

type PaymentStatusUpdater struct {
}

func (pS PaymentStatusUpdater) Process(context polaris.BuilderContext) polaris.IData {
	fmt.Println("In payment status updater")
	// schedule warehouse ops if payment is successful
	return WarehouseOpsScheduled{}
}

func (pS PaymentStatusUpdater) GetBuilderInfo() polaris.BuilderInfo {
	return polaris.BuilderInfo{
		Consumes: []polaris.IData{
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

func (wS WarehouseStatusUpdater) Process(context polaris.BuilderContext) polaris.IData {
	fmt.Println("In warehouse status updater")
	// if order was delivered produce output, otherwise produce null and flow will stop
	return OrderDelivered{}
}

func (wS WarehouseStatusUpdater) GetBuilderInfo() polaris.BuilderInfo {
	return polaris.BuilderInfo{
		Consumes: []polaris.IData{
			WarehouseStatus{},
		},
		Produces:  OrderDelivered{},
		Optionals: nil,
		Accesses:  nil,
	}
}

type WorkflowTerminator struct {
}

func (wT WorkflowTerminator) Process(context polaris.BuilderContext) polaris.IData {
	fmt.Println("In workflow terminator")
	// schedule warehouse ops if payment is successful
	return WarehouseOpsScheduled{}
}

func (wT WorkflowTerminator) GetBuilderInfo() polaris.BuilderInfo {
	return polaris.BuilderInfo{
		Consumes: []polaris.IData{
			OrderDelivered{},
		},
		Produces:  WorkflowTerminated{},
		Optionals: nil,
		Accesses:  nil,
	}
}
