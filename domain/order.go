package domain

import (
	"context"
)

const (
	CollectionOrder = "order"
)

type Order struct {
	ID           string `json:"id"`
	IDUser       string `json:"id_user"`
	IDCourier    string `json:"id_courier"`
	IDAddress    string `json:"id_address"`
	TotalPrice   string `json:"total_price"`
	StatusOrder  string `json:"status_order"`
	DeliveryType string `json:"delivery_type"`
}

type OrderRepository interface {
	Create(c context.Context, product *Order) error
	GetById(c context.Context, id string) (Order, error)
	GetByIdUser(c context.Context, id string) ([]Order, error)
	UpdateById(c context.Context, id string, order Order) (Order, error)
}

type OrderUsecase interface {
	Create(c context.Context, product *Order) error
	GetById(c context.Context, id string) (Order, error)
	GetByIdUser(c context.Context, id string) ([]Order, error)
	UpdateById(c context.Context, id string, order Order) (Order, error)
}
