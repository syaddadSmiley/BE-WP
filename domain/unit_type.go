package domain

import (
	"context"
)

const (
	CollectionUnitType = "unit_type"
)

type UnitType struct {
	ID           string `json:"id"`
	UnitTypeName string `json:"unit_type_name" form:"unit_type_name" binding:"required"`
}

type UnitTypeRepository interface {
	Create(c context.Context, unitType *UnitType) error
	GetById(c context.Context, id string) (UnitType, error)
}

type UnitTypeUsecase interface {
	Create(c context.Context, unitType *UnitType) error
	GetById(c context.Context, id string) (UnitType, error)
}
