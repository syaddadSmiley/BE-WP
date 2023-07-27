package usecase

import (
	"context"
	"time"

	"waroeng_pgn1/domain"
)

type paymentUsecase struct {
	paymentRepository domain.PaymentRepository
	contextTimeout    time.Duration
}

func NewPaymentUsecase(paymentRepository domain.PaymentRepository, timeout time.Duration) domain.PaymentUsecase {
	return &paymentUsecase{
		paymentRepository: paymentRepository,
		contextTimeout:    timeout,
	}
}

func (pu *paymentUsecase) Create(c context.Context, payment *domain.Payment) error {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()

	err := pu.paymentRepository.Create(ctx, payment)
	if err != nil {
		return err
	}

	return nil
}

func (pu *paymentUsecase) GetById(c context.Context, id string) (domain.Payment, error) {

	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()

	payment, err := pu.paymentRepository.GetById(ctx, id)
	if err != nil {
		return payment, err
	}

	return payment, nil
}

func (pu *paymentUsecase) GetByIdOrder(c context.Context, id string) (domain.Payment, error) {

	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()

	payment, err := pu.paymentRepository.GetByIdOrder(ctx, id)
	if err != nil {
		return payment, err
	}

	return payment, nil
}

func (pu *paymentUsecase) UpdateById(c context.Context, id string, payment domain.Payment) (domain.Payment, error) {

	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()

	payment, err := pu.paymentRepository.UpdateById(ctx, id, payment)
	if err != nil {
		return payment, err
	}

	return payment, nil
}
