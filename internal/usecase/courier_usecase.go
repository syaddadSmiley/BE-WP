package usecase

import (
	"context"
	"time"

	"waroeng_pgn1/domain"
)

type courierUsecase struct {
	courierRepository domain.CourierRepository
	contextTimeout    time.Duration
}

func NewCourierUsecase(courierRepository domain.CourierRepository, timeout time.Duration) domain.CourierUsecase {
	return &courierUsecase{
		courierRepository: courierRepository,
		contextTimeout:    timeout,
	}
}

func (cu *courierUsecase) Create(c context.Context, courier *domain.Courier) error {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.courierRepository.Create(ctx, courier)
}

func (cu *courierUsecase) GetAll(c context.Context) ([]domain.Courier, error) {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.courierRepository.GetAll(ctx)
}

func (cu *courierUsecase) GetIdCityByName(c context.Context, name string) (string, error) {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.courierRepository.GetIdCityByName(ctx, name)
}
