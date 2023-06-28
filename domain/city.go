package domain

import (
	"context"
)

const (
	CollectionCity = "city"
)

type City struct {
	Rajaongkir struct {
		Query  []string `json:"query"`
		Status struct {
			Code        int    `json:"code"`
			Description string `json:"description"`
		} `json:"status"`
		Results []CityResult `json:"results"`
	} `json:"rajaongkir"`
}

type CityResult struct {
	CityID     string `json:"city_id"`
	ProvinceID string `json:"province_id"`
	Province   string `json:"province"`
	Type       string `json:"type"`
	CityName   string `json:"city_name"`
	PostalCode string `json:"postal_code"`
}

type CityRepository interface {
	InsertCity(c context.Context, city []CityResult) ([]CityResult, error)
	GetCityByProvince(c context.Context, province string) ([]CityResult, error)
}

type CityUsecase interface {
	InsertCity(c context.Context, city []CityResult) ([]CityResult, error)
	GetCityByProvince(c context.Context, province string) ([]CityResult, error)
}
