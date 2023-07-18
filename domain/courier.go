package domain

import (
	"context"
)

const (
	CollectionCourier = "courier"
)

type Courier struct {
	ID          string `json:"id"`
	CourierName string `json:"courier_name"`
	IsAvailable bool   `json:"is_available"`
}

type CostRequest struct {
	City    string `json:"city"`
	Weight  int    `json:"weight"`
	Courier string `json:"courier"`
}

type CostResponse struct {
	Rajaongkir struct {
		Query struct {
			Origin      string `json:"origin"`
			Destination string `json:"destination"`
			Weight      int    `json:"weight"`
			Courier     string `json:"courier"`
		} `json:"query"`
		Status struct {
			Code        int    `json:"code"`
			Description string `json:"description"`
		} `json:"status"`
		OriginDetails struct {
			CityID     string `json:"city_id"`
			Province   string `json:"province"`
			ProvinceID string `json:"province_id"`
			Type       string `json:"type"`
			CityName   string `json:"city_name"`
			PostalCode string `json:"postal_code"`
		} `json:"origin_details"`
		DestinationDetails struct {
			CityID     string `json:"city_id"`
			Province   string `json:"province"`
			ProvinceID string `json:"province_id"`
			Type       string `json:"type"`
			CityName   string `json:"city_name"`
			PostalCode string `json:"postal_code"`
		} `json:"destination_details"`
		Results []CostResult `json:"results"`
	} `json:"rajaongkir"`
}

type CostResult struct {
	Code  string `json:"code"`
	Name  string `json:"name"`
	Costs []Cost `json:"costs"`
}

type Cost struct {
	Service     string       `json:"service"`
	Description string       `json:"description"`
	Cost        []CostDetail `json:"cost"`
}

type CostDetail struct {
	Value int    `json:"value"`
	Etd   string `json:"etd"`
	Note  string `json:"note"`
}

type CourierRepository interface {
	Create(c context.Context, courier *Courier) error
	GetAll(c context.Context) ([]Courier, error)
	GetIdCityByName(c context.Context, name string) (string, error)
}

type CourierUsecase interface {
	Create(c context.Context, courier *Courier) error
	GetAll(c context.Context) ([]Courier, error)
	GetIdCityByName(c context.Context, name string) (string, error)
}
