package domain

import (
	"context"
)

const (
	CollectionProvince = "province"
)

type Province struct {
	Rajaongkir struct {
		Query  []string `json:"query"`
		Status struct {
			Code        int    `json:"code"`
			Description string `json:"description"`
		} `json:"status"`
		Results []ProvinceResult `json:"results"`
	} `json:"rajaongkir"`
}

type ProvinceResult struct {
	ProvinceID string `json:"province_id"`
	Province   string `json:"province"`
}

type ProvinceRepository interface {
	InsertProvince(c context.Context, province []ProvinceResult) ([]ProvinceResult, error)
	GetProvince(c context.Context) ([]ProvinceResult, error)
}

type ProvinceUsecase interface {
	InsertProvince(c context.Context, province []ProvinceResult) ([]ProvinceResult, error)
	GetProvince(c context.Context) ([]ProvinceResult, error)
}
