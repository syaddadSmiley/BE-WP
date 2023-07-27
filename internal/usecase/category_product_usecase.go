package usecase

import (
	"context"
	"time"

	"waroeng_pgn1/domain"
)

type categoryUsecase struct {
	categoryProductRepository domain.CategoryProductRepository
	contextTimeout            time.Duration
}

func NewCategoryUsecase(categoryProductRepository domain.CategoryProductRepository, timeout time.Duration) domain.CategoryProductUsecase {
	return &categoryUsecase{
		categoryProductRepository: categoryProductRepository,
		contextTimeout:            timeout,
	}
}

func (cu *categoryUsecase) Create(c context.Context, category *domain.CategoryProduct) error {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.categoryProductRepository.Create(ctx, category)
}

func (cu *categoryUsecase) GetById(c context.Context, id string) (domain.CategoryProduct, error) {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.categoryProductRepository.GetById(ctx, id)
}

func (cu *categoryUsecase) GetAll(c context.Context) ([]domain.CategoryProduct, error) {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.categoryProductRepository.GetAll(ctx)
}
