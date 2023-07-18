package usecase

import (
	"context"
	"time"

	"waroeng_pgn1/domain"
)

type unitTypeUsecase struct {
	unitTypeRepository domain.UnitTypeRepository
	contextTimeout     time.Duration
}

func NewUnitTypeUsecase(unitTypeRepository domain.UnitTypeRepository, timeout time.Duration) domain.UnitTypeUsecase {
	return &unitTypeUsecase{
		unitTypeRepository: unitTypeRepository,
		contextTimeout:     timeout,
	}
}

func (utu *unitTypeUsecase) Create(c context.Context, unitType *domain.UnitType) error {
	ctx, cancel := context.WithTimeout(c, utu.contextTimeout)
	defer cancel()
	return utu.unitTypeRepository.Create(ctx, unitType)
}

func (utu *unitTypeUsecase) GetById(c context.Context, unitTypeID string) (domain.UnitType, error) {
	ctx, cancel := context.WithTimeout(c, utu.contextTimeout)
	defer cancel()
	return utu.unitTypeRepository.GetById(ctx, unitTypeID)
}
