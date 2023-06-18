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
	IDCategory         string `json:"id_category"`
	IDUnitType         string `json:"id_unit_type"`
	NameProduct        string `json:"name_product" form:"name_product" binding:"required"`
	Price              string `json:"price" form:"price" binding:"required"`
	DescriptionProduct string `json:"description_product" form:"description_product" binding:"required"`
	Discount           string `json:"discount"`
	SoldAmount         string `json:"sold_amount"`
	Stock              string `json:"stock" form:"stock" binding:"required"`
	Location           string `json:"location"`
}

type ProductRepository interface {
	Create(c context.Context, product *Product) error
	GetById(c context.Context, id string) (Product, error)
	GetAll(c context.Context) ([]Product, error)
}

type ProductUsecase interface {
	Create(c context.Context, product *Product) error
	GetById(c context.Context, id string) (Product, error)
	GetAll(c context.Context) ([]Product, error)
}
