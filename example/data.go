package example

import "time"

type OrderRequest struct {
	ProductId int
	Qty       int
	UserId    string
	AddressId string
}

type OrderRequestValidated struct {
}

type RiskCheckCompleted struct {
}

type OrderInfo struct {
	OrderId     string
	TotalAmount int
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

type WarehouseStatusUpdateRequest struct {
	OrderId    string
	EtaInHours int
	Status     string
}

type OrderDelivered struct {
	OrderId string
	Time    time.Time
}

type WorkflowTerminated struct {
}
