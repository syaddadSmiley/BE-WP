package domain

import (
	"context"
)

const (
	CollectionProduct = "product"
)

type Product struct {
	ID                 string `json:"id"`
	IDPriceType        string `json:"id_price_type"`
	IDUnitType         string `json:"id_category"`
	NameProduct        string `json:"name_product"`
	Price              string `json:"price"`
	DescriptionProduct string `json:"description_product"`
	Discount           string `json:"discount"`
	SoldAmount         string `json:"sold_amount"`
	Stock              string `json:"stock"`
	StockType          string `json:"stock_type"`
	Location           string `json:"location"`
}

type ProductRepository interface {
	Create(c context.Context, product *Product) error
	GetById(c context.Context, id string) (Product, error)
}

type ProductUsecase interface {
	Create(c context.Context, product *Product) error
	GetById(c context.Context, id string) (Product, error)
}
