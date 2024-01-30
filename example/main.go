package main

import (
	"fmt"
	"reflect"
	"time"

	"github.com/harshadmanglani/polaris"
)

type OrderRequest struct {
	ProductId int
	Qty       int
	UserId    string
	AddressId string
}

type RequestValidated struct {
}

type RiskCheck1Completed struct {
}

type RiskCheck2Completed struct {
}

type OrderInfo struct {
	OrderId     string
	TotalAmount int
}

type PaymentInitialized struct {
	PaymentUrl string
	ExpiresAt  time.Time
}

type PaymentStatus struct {
	OrderId   string
	PaymentId string
	Status    string
}

type WarehouseOpsScheduled struct {
	EtaInHours int
	Status     string
}

type WarehouseStatus struct {
	OrderId    string
	EtaInHours int
	Status     string
}

type OrderDelivered struct {
	OrderId     string
	DeliveredAt time.Time
}

type WorkflowTerminated struct {
}

type OmsWorkflow struct {
}

func (omsW OmsWorkflow) GetWorkflowMeta() polaris.WorkflowMeta {
	return polaris.WorkflowMeta{
		// the ordering of builders below is for readability, the framework will evaluate the exec graph irrespective of it
		Builders: []polaris.IBuilder{
			Validator{},
			RiskChecker1{}, RiskChecker2{},
			Persistor{},
			PaymentInitializer{},
			PaymentStatusUpdater{},
			WarehouseStatusUpdater{},
			WorkflowTerminator{},
		},
		TargetData: WorkflowTerminated{},
	}
}

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
	return RiskCheck2Completed{}
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
		Produces: OrderDelivered{},
		Optionals: []polaris.IData{
			OrderInfo{},
		},
		Accesses: nil,
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

func main() {
	polaris.RegisterWorkflow("OMSWORKFLOW", OmsWorkflow{})

	e := polaris.Executor{
		Before: func(builder reflect.Type, delta []polaris.IData) {
			fmt.Println("Before execution")
		},
		After: func(builder reflect.Type, produced polaris.IData) {
			fmt.Println("After execution")
		},
	}
	responses := e.Run("OMSWORKFLOW", "someUniqueId", OrderRequest{
		ProductId: 12,
		Qty:       1,
		UserId:    "abcd",
		AddressId: "1234",
	})
	fmt.Println(responses)
}
