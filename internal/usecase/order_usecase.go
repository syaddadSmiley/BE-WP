package usecase

import (
	"context"
	"time"

	"waroeng_pgn1/domain"
)

type orderUsecase struct {
	orderRepository domain.OrderRepository
	contextTimeout  time.Duration
}

func NewOrderUsecase(orderRepository domain.OrderRepository, timeout time.Duration) domain.OrderUsecase {
	return &orderUsecase{
		orderRepository: orderRepository,
		contextTimeout:  timeout,
	}
}

func (ou *orderUsecase) Create(c context.Context, order *domain.Order) error {
	ctx, cancel := context.WithTimeout(c, ou.contextTimeout)
	defer cancel()
	return ou.orderRepository.Create(ctx, order)
}

func (ou *orderUsecase) CreateOrderStatus(c context.Context, orderStatus *domain.OrderStatus) error {
	ctx, cancel := context.WithTimeout(c, ou.contextTimeout)
	defer cancel()
	return ou.orderRepository.CreateOrderStatus(ctx, orderStatus)
}

func (ou *orderUsecase) CreateOrderItem(c context.Context, orderItem *domain.OrderItems) error {
	ctx, cancel := context.WithTimeout(c, ou.contextTimeout)
	defer cancel()
	return ou.orderRepository.CreateOrderItem(ctx, orderItem)
}

func (ou *orderUsecase) GetById(c context.Context, orderID string) (domain.Order, error) {
	ctx, cancel := context.WithTimeout(c, ou.contextTimeout)
	defer cancel()
	return ou.orderRepository.GetById(ctx, orderID)
}

func (ou *orderUsecase) GetByIdUser(c context.Context, userID string) ([]domain.Order, error) {
	ctx, cancel := context.WithTimeout(c, ou.contextTimeout)
	defer cancel()
	return ou.orderRepository.GetByIdUser(ctx, userID)
}

func (ou *orderUsecase) GetOrderItemsByIdOrder(c context.Context, orders []domain.Order) ([]domain.Order, error) {
	ctx, cancel := context.WithTimeout(c, ou.contextTimeout)
	defer cancel()
	return ou.orderRepository.GetOrderItemsByIdOrder(ctx, orders)
}

func (ou *orderUsecase) UpdateById(c context.Context, orderID string, order domain.Order) (domain.Order, error) {
	ctx, cancel := context.WithTimeout(c, ou.contextTimeout)
	defer cancel()
	return ou.orderRepository.UpdateById(ctx, orderID, order)
}
