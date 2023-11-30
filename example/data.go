package main

import "time"

type OrderRequest struct {
	ProductId int
	Qty       int
	UserId    string
	AddressId string
}

type RequestValidated struct {
}

type RiskCheckCompleted struct {
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

type WarehouseStatusUpdateRequest struct {
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
