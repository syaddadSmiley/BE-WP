package domain

import (
	"context"
)

const (
	CollectionCourierService = "courier_service"
)

type CourierService struct {
	ID                   string `json:"id"`
	IDCourier            string `json:"id_courier"`
	IDAddress            string `json:"id_address"`
	ServiceName          string `json:"service_name"`
	ReceiptServiceNumber string `json:"receipt_service_number"`
	Description          string `json:"description"`
	PriceService         string `json:"price_service"`
	EstimationDay        string `json:"estimation_day"`
}

type CourierServiceRepository interface {
	Create(c context.Context, courierService *CourierService) error
	GetById(c context.Context, id string) (CourierService, error)
}

type CourierServiceUsecase interface {
	Create(c context.Context, courierService *CourierService) error
	GetById(c context.Context, id string) (CourierService, error)
}
