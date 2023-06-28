package usecase

import (
	"context"
	"time"

	"waroeng_pgn1/domain"
)

type provinceUsecase struct {
	provinceRepository domain.ProvinceRepository
	contextTimeout     time.Duration
}

func NewProvinceUsecase(provinceRepository domain.ProvinceRepository, timeout time.Duration) domain.ProvinceUsecase {
	return &provinceUsecase{
		provinceRepository: provinceRepository,
		contextTimeout:     timeout,
	}
}

func (pu *provinceUsecase) InsertProvince(c context.Context, province []domain.ProvinceResult) ([]domain.ProvinceResult, error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	return pu.provinceRepository.InsertProvince(ctx, province)
}

func (pu *provinceUsecase) GetProvince(c context.Context) ([]domain.ProvinceResult, error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	return pu.provinceRepository.GetProvince(ctx)
}
