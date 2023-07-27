package domain

import (
	"context"
)

const (
	CollectionPayment = "payment"
)

type Payment struct {
	ID               string `json:"id"`
	IDOrder          string `json:"id_order"`
	Status           string `json:"status"`
	MetodePembayaran string `json:"metode_pembayaran"`
}

type MidtransTransactionStatusRespone struct {
	TransactionID     string `json:"transaction_id"`
	TransactionStatus string `json:"transaction_status"`
	PaymentType       string `json:"payment_type"`
}

type PaymentRepository interface {
	Create(c context.Context, payment *Payment) error
	GetById(c context.Context, id string) (Payment, error)
	GetByIdOrder(c context.Context, id string) (Payment, error)
	UpdateById(c context.Context, id string, payment Payment) (Payment, error)
}

type PaymentUsecase interface {
	Create(c context.Context, payment *Payment) error
	GetById(c context.Context, id string) (Payment, error)
	GetByIdOrder(c context.Context, id string) (Payment, error)
	UpdateById(c context.Context, id string, payment Payment) (Payment, error)
}
