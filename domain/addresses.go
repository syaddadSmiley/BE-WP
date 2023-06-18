package domain

import (
	"context"
)

const (
	CollectionAddresses = "addresses"
)

type Addresses struct {
	ID           string `json:"id"`
	IDUser       string `json:"id_user"`
	LabelAddress string `json:"label_address"`
	Address      string `json:"address"`
	City         string `json:"city"`
	Province     string `json:"province"`
	PostalCode   string `json:"postal_code"`
	Latitude     string `json:"latitude"`
	Longitude    string `json:"longitude"`
	Note         string `json:"note"`
	IsDefault    string `json:"is_default"`
}

type AddressesRepository interface {
	Create(c context.Context, product *Addresses) error
	GetById(c context.Context, id string) (Addresses, error)
	GetByIdUser(c context.Context, id string) ([]Addresses, error)
	UpdateById(c context.Context, id string, addresses Addresses) (Addresses, error)
}

type AddressesUsecase interface {
	Create(c context.Context, product *Addresses) error
	GetById(c context.Context, id string) (Addresses, error)
	GetByIdUser(c context.Context, id string) ([]Addresses, error)
	UpdateById(c context.Context, id string, addresses Addresses) (Addresses, error)
}
