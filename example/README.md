## A simple Order Management System with Polaris

(B) - a builder - a piece of code that does some tasks
(D) - data required as inputs to builders, or produces as an output
The flow looks like this:


When a customer tries to place an order
POST /order ->
```
          OrderRequest (D)-----------
               |              |     |
Level 0:    Validator (B)     |     |
               |              |     |
         RequestValidated (D) |     |
               |              |     |
Level 1:        RiskChecker (B)     |
                      |             |
            RiskCheckCompleted (D)  |
                      |             |
Level 2:               Persistor (B)
                            |
                        OrderInfo (D)
                            |
Level 3:                PaymentInitializer (B)
                            |
                        PaymentInitialized (D)
```
<- Response

When the Payment Gateway gives a callback for a status update
POST /payments/callback ->
```
                        PaymentStatus (D)
                            |
Level 4:                PaymentStatusUpdater (B)
                            |
                        WarehouseOpsScheduled (D)
```
<- Response (this is for demonstration purposes, I would recommend handling callbacks asynchronously on production)

When the warehouse gives a callback for a status update on delivery
POST /warehouse/callback ->
```
                        WarehouseStatusUpdateRequest (D)
                            |
Level 5:                WarehouseStatusUpdater (B)-------------------------------
                            |                                                   |
                        null (if it is an intermediate update)          OrderDelivered (D)
                                                                                |
                                                                        WorkflowTerminator (B)
                                                                                |
                                                                        WorkflowTerminated (D)
```
<- Response