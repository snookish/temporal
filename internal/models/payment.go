package models

type PaymentMethod string

const (
	PaymentMethodUPI  PaymentMethod = "UPI"
	PaymentMethodCard PaymentMethod = "CARD"
)
