package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/harshadmanglani/polaris/core"
	"github.com/harshadmanglani/polaris/models"
)

type OmsWorkflow struct {
}

func (omsW OmsWorkflow) GetWorkflowMeta() models.WorkflowMeta {
	return models.WorkflowMeta{
		// the ordering of builders below is for readability, the framework will evaluate the exec graph irrespective of it
		Builders: []models.IBuilder{
			Validator{},
			RiskChecker1{}, RiskChecker2{}, // perform these in parallel
			Persistor{},
			PaymentInitializer{},
			PaymentStatusUpdater{},
			WarehouseStatusUpdater{},
			WorkflowTerminator{},
		},
		TargetData: WorkflowTerminated{},
	}
}

func createOrder(w http.ResponseWriter, r *http.Request) {
	var orderRequest OrderRequest
	json.NewDecoder(r.Body).Decode(&orderRequest)
	response := executor.Run(&omsDataFlow, orderRequest)
	json.NewEncoder(w).Encode(response.Responses[models.Name(PaymentInitialized{})])
}

func handlePaymentsCb(w http.ResponseWriter, r *http.Request) {
	var paymentStatus PaymentStatus
	json.NewDecoder(r.Body).Decode(&paymentStatus)
	response := executor.Run(&omsDataFlow, paymentStatus)
	json.NewEncoder(w).Encode(response.Responses[models.Name(WarehouseOpsScheduled{})])
}

func handleOrderCb(w http.ResponseWriter, r *http.Request) {
	var warehouseStatus WarehouseStatus
	json.NewDecoder(r.Body).Decode(&warehouseStatus)
	response := executor.Run(&omsDataFlow, warehouseStatus)
	json.NewEncoder(w).Encode(response.Responses[models.Name(OrderDelivered{})])
}

var (
	omsDataFlow models.DataFlow
	executor    core.Executor
)

func init() {
	omsDataFlow = core.RegisterWorkflow(OmsWorkflow{})
	executor = core.Executor{
		Before:  func() { fmt.Printf("Builder X is about to be executed") },
		After:   func() { fmt.Printf("Builder X executed successfully and generated data Y") },
		OnError: func() { fmt.Printf("Builder X errored with stack trace Z") },
	}
}

func main() {
	http.HandleFunc("/order", createOrder)
	http.HandleFunc("/payments/callback", handlePaymentsCb)
	http.HandleFunc("/order/callback", handleOrderCb)

	port := 8080
	fmt.Printf("Server is running on :%d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
