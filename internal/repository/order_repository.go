package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"waroeng_pgn1/domain"
)

type orderRepository struct {
	database   *sql.DB
	collection string
}

func NewOrderRepository(db *sql.DB, collection string) domain.OrderRepository {
	return &orderRepository{
		database:   db,
		collection: collection,
	}
}

func (or *orderRepository) Create(c context.Context, order *domain.Order) error {
	stmt, err := or.database.Prepare(`INSERT INTO waroeng_pgn1.order (id, id_user, id_courier_service, id_address, total_price, current_status_order, is_refund) VALUES (?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()
	result, err := stmt.Exec(order.ID, order.IDUser, order.IDCourierService, order.IDAddress, order.TotalPrice, order.CurrentStatusOrder, order.IsRefund)
	if err != nil {
		return err
	} else if result != nil {
		return nil
	}
	return errors.New("error while creating order")
}

func (or *orderRepository) CreateOrderStatus(c context.Context, order *domain.OrderStatus) error {
	stmt, err := or.database.Prepare(`INSERT INTO order_status (id, id_order, status) VALUES (?, ?, ?)`)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()
	result, err := stmt.Exec(order.ID, order.IDOrder, order.Status)
	if err != nil {
		return err
	} else if result != nil {
		return nil
	}
	return errors.New("error while creating order status")
}

func (or *orderRepository) CreateOrderItem(c context.Context, order *domain.OrderItems) error {
	stmt, err := or.database.Prepare(`INSERT INTO order_items (id, id_order, id_product, quantity, price, description) VALUES (?, ?, ?, ?, ?, ?)`)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()
	result, err := stmt.Exec(order.ID, order.IDOrder, order.IDProduct, order.Quantity, order.Price, order.Description)
	if err != nil {
		return err
	} else if result != nil {
		return nil
	}
	return errors.New("error while creating order item")
}

func (or *orderRepository) GetById(c context.Context, id string) (domain.Order, error) {
	var order domain.Order
	stmt, err := or.database.Prepare(`SELECT id, id_user, id_courier_service, id_address, total_price, current_status_order, is_refund FROM waroeng_pgn1.order WHERE id = ?`)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&order.ID, &order.IDUser, &order.IDCourierService, &order.IDAddress, &order.TotalPrice, &order.CurrentStatusOrder, &order.IsRefund)
	if err != nil {
		return order, err
	} else if order.ID == "" {
		return order, errors.New("order not found")
	}

	return order, nil
}

func (or *orderRepository) GetByIdUser(c context.Context, id string) ([]domain.Order, error) {
	var orders []domain.Order
	fmt.Println("id", id)
	stmt, err := or.database.Prepare(`SELECT id, id_user, id_courier_service, id_address, total_price, current_status_order, is_refund FROM waroeng_pgn1.order WHERE id_user = ?`)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()
	rows, err := stmt.Query(id)
	if err != nil {
		return orders, err
	}

	for rows.Next() {
		var order domain.Order
		err = rows.Scan(&order.ID, &order.IDUser, &order.IDCourierService, &order.IDAddress, &order.TotalPrice, &order.CurrentStatusOrder, &order.IsRefund)
		if err != nil {
			return orders, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

func (or *orderRepository) GetOrderItemsByIdOrder(c context.Context, orders []domain.Order) ([]domain.Order, error) {
	for i, order := range orders {
		stmt, err := or.database.Prepare(`SELECT oi.id, p.name_product, oi.quantity, oi.price, oi.id_product, oi.id_order, oi.description FROM order_items oi INNER JOIN product p ON oi.id_product = p.id  WHERE id_order = ?`)
		if err != nil {
			panic(err)
		}

		defer stmt.Close()

		rows, err := stmt.Query(order.ID)
		if err != nil {
			return orders, err
		}

		for rows.Next() {
			var orderItem domain.OrderItems
			err = rows.Scan(&orderItem.ID, &orderItem.Name, &orderItem.Quantity, &orderItem.Price, &orderItem.IDProduct, &orderItem.IDOrder, &orderItem.Description)
			if err != nil {
				return orders, err
			}
			orders[i].OrderItems = append(orders[i].OrderItems, orderItem)
		}
	}

	return orders, nil
}

func (or *orderRepository) UpdateById(c context.Context, id string, order domain.Order) (domain.Order, error) {
	stmt, err := or.database.Prepare(`UPDATE waroeng_pgn1.order SET id_user = ?, id_courier_service = ?, id_address = ?, total_price = ?, current_status_order = ?, is_refund = ? WHERE id = ?`)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()
	result, err := stmt.Exec(order.IDUser, order.IDCourierService, order.IDAddress, order.TotalPrice, order.CurrentStatusOrder, order.IsRefund, id)
	if err != nil {
		return order, err
	} else if result != nil {
		return order, nil
	}
	return order, errors.New("error while updating order")
}
