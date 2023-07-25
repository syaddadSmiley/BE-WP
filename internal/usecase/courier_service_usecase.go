package usecase

import (
	"context"
	"time"

	"waroeng_pgn1/domain"
)

type courierServiceUsecase struct {
	courierServiceRepository domain.CourierServiceRepository
	contextTimeout           time.Duration
}

func NewCourierServiceUsecase(courierServiceRepository domain.CourierServiceRepository, timeout time.Duration) domain.CourierServiceUsecase {
	return &courierServiceUsecase{
		courierServiceRepository: courierServiceRepository,
		contextTimeout:           timeout,
	}
}

func (csu *courierServiceUsecase) Create(c context.Context, courierService *domain.CourierService) error {
	ctx, cancel := context.WithTimeout(c, csu.contextTimeout)
	defer cancel()
	return csu.courierServiceRepository.Create(ctx, courierService)
}

func (csu *courierServiceUsecase) GetById(c context.Context, id string) (domain.CourierService, error) {
	ctx, cancel := context.WithTimeout(c, csu.contextTimeout)
	defer cancel()
	return csu.courierServiceRepository.GetById(ctx, id)
}
