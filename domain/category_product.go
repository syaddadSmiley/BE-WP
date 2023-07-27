package domain

import (
	"context"
)

const (
	CollectionCategory = "category_product"
)

type CategoryProduct struct {
	ID           string `json:"id"`
	NameCategory string `json:"name_category"`
	TypeCategory string `json:"type_category"`
}

type CategoryProductRepository interface {
	Create(c context.Context, category *CategoryProduct) error
	GetById(c context.Context, id string) (CategoryProduct, error)
	GetAll(c context.Context) ([]CategoryProduct, error)
}

type CategoryProductUsecase interface {
	Create(c context.Context, category *CategoryProduct) error
	GetById(c context.Context, id string) (CategoryProduct, error)
	GetAll(c context.Context) ([]CategoryProduct, error)
}
