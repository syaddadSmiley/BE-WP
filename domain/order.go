package domain

import (
	"context"
)

const (
	CollectionOrder = "order"
)

type Order struct {
	ID                 string       `json:"id"`
	IDUser             string       `json:"id_user"`
	IDCourierService   string       `json:"id_courier_service"`
	IDAddress          string       `json:"id_address"`
	TotalPrice         string       `json:"total_price"`
	TaxPrice           string       `json:"tax_price"`
	OrderItems         []OrderItems `json:"order_items"`
	CurrentStatusOrder string       `json:"current_status_order"`
	IsRefund           bool         `json:"is_refund"`
}

type OrderItems struct {
	ID        string `json:"id"`
	IDOrder   string `json:"id_order"`
	Name      string `json:"name"`
	IDProduct string `json:"id_product"`
	Quantity  string `json:"quantity"`
	Price     string `json:"price"`
}

// type OrderRequest struct {

type MidtransResponse struct {
	Token       string `json:"token"`
	RedirectURL string `json:"redirect_url"`
}

type OrderRepository interface {
	Create(c context.Context, product *Order) error
	CreateOrderItem(c context.Context, orderItem *OrderItems) error
	GetById(c context.Context, id string) (Order, error)
	GetByIdUser(c context.Context, id string) ([]Order, error)
	UpdateById(c context.Context, id string, order Order) (Order, error)
}

type OrderUsecase interface {
	Create(c context.Context, product *Order) error
	CreateOrderItem(c context.Context, orderItem *OrderItems) error
	GetById(c context.Context, id string) (Order, error)
	GetByIdUser(c context.Context, id string) ([]Order, error)
	UpdateById(c context.Context, id string, order Order) (Order, error)
}
