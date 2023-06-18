package usecase

import (
	"context"
	"time"

	"waroeng_pgn1/domain"
)

type addressesUsecase struct {
	addressesRepository domain.AddressesRepository
	contextTimeout      time.Duration
}

func NewAddressesUsecase(addressesRepository domain.AddressesRepository, timeout time.Duration) domain.AddressesUsecase {
	return &addressesUsecase{
		addressesRepository: addressesRepository,
		contextTimeout:      timeout,
	}
}

func (au *addressesUsecase) Create(c context.Context, addresses *domain.Addresses) error {
	ctx, cancel := context.WithTimeout(c, au.contextTimeout)
	defer cancel()
	return au.addressesRepository.Create(ctx, addresses)
}

func (au *addressesUsecase) GetById(c context.Context, addressesID string) (domain.Addresses, error) {
	ctx, cancel := context.WithTimeout(c, au.contextTimeout)
	defer cancel()
	return au.addressesRepository.GetById(ctx, addressesID)
}

func (au *addressesUsecase) GetByIdUser(c context.Context, addressesID string) ([]domain.Addresses, error) {
	ctx, cancel := context.WithTimeout(c, au.contextTimeout)
	defer cancel()
	return au.addressesRepository.GetByIdUser(ctx, addressesID)
}

func (au *addressesUsecase) UpdateById(c context.Context, addressesID string, addresses domain.Addresses) (domain.Addresses, error) {
	ctx, cancel := context.WithTimeout(c, au.contextTimeout)
	defer cancel()
	return au.addressesRepository.UpdateById(ctx, addressesID, addresses)
}
