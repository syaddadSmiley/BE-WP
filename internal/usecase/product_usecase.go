package usecase

import (
	"context"
	"time"

	"waroeng_pgn1/domain"
)

type productUsecase struct {
	productRepository domain.ProductRepository
	contextTimeout    time.Duration
}

func NewProductUsecase(productRepository domain.ProductRepository, timeout time.Duration) domain.ProductUsecase {
	return &productUsecase{
		productRepository: productRepository,
		contextTimeout:    timeout,
	}
}

func (pu *productUsecase) Create(c context.Context, product *domain.Product) error {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	return pu.productRepository.Create(ctx, product)
}

func (pu *productUsecase) GetById(c context.Context, productID string) (domain.Product, error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	return pu.productRepository.GetById(ctx, productID)
}

// func (tu *ProductUsecase) FetchByUserID(c context.Context, userID string) ([]domain.Task, error) {
// 	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
// 	defer cancel()
// 	return tu.productRepository.FetchByUserID(ctx, userID)
// }
