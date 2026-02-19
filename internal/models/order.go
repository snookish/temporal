package models

type OrderStatus string

const (
	OrderStatusPending                   OrderStatus = "PENDING"
	OrderStatusPaymentFailed             OrderStatus = "PAYMENT_FAILED"
	OrderStatusPaymentSucceeded          OrderStatus = "PAYMENT_SUCCEEDED"
	OrderStatusReserveInventoryFailed    OrderStatus = "RESERVE_INVENTORY_FAILED"
	OrderStatusReserveInventorySucceeded OrderStatus = "RESERVE_INVENTORY_SUCCEEDED"
)

type CompensationStep string

var (
	CompensationStepRefund CompensationStep = "REFUND_PAYMENT"
)

type ProcessOrderCommand struct {
	OrderID       string
	CustomerID    string
	Items         []OrderItem
	Amount        Money
	PaymentMethod PaymentMethod
}

type Money struct {
	Currency string
	Amount   float64
}

type OrderItem struct {
	Name      string
	ProductID string
	Quantity  uint
	UnitPrice Money
}

type OrderState struct {
	OrderID           string
	ShipmentID        string
	Status            OrderStatus
	CompensationSteps []CompensationStep
}

func (o *OrderState) Is(status OrderStatus) bool {
	return o.Status == status
}

func (o *OrderState) IsPaid() bool {
	return o.Is(OrderStatusPaymentSucceeded)
}

func (o *OrderState) IsPending() bool {
	return o.Is(OrderStatusPending)
}
