package models

type PaymentMethod string

const (
	PaymentMethodUPI  PaymentMethod = "UPI"
	PaymentMethodCard PaymentMethod = "CARD"
)

type TransactionStatus string

const (
	TransactionStatusFailed  TransactionStatus = "FAILED"
	TransactionStatusSuccess TransactionStatus = "SUCCESS"
)

type ProcessPaymentCommand struct {
	OrderID       string
	Amount        Money
	PaymentMethod PaymentMethod
}

type ProcessPaymentResponse struct {
	TransactionID string
	Status        TransactionStatus
}
