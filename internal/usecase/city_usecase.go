package usecase

import (
	"context"
	"time"

	"waroeng_pgn1/domain"
)

type cityUsecase struct {
	cityRepository domain.CityRepository
	contextTimeout time.Duration
}

func NewCityUsecase(cityRepository domain.CityRepository, timeout time.Duration) domain.CityUsecase {
	return &cityUsecase{
		cityRepository: cityRepository,
		contextTimeout: timeout,
	}
}

func (pu *cityUsecase) InsertCity(c context.Context, city []domain.CityResult) ([]domain.CityResult, error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	return pu.cityRepository.InsertCity(ctx, city)
}

func (pu *cityUsecase) GetCityByProvince(c context.Context, province string) ([]domain.CityResult, error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	return pu.cityRepository.GetCityByProvince(ctx, province)
}
