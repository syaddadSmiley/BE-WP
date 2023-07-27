package repository

import (
	"context"
	"database/sql"

	"waroeng_pgn1/domain"
)

type paymentRepository struct {
	database   *sql.DB
	collection string
}

func NewPaymentRepository(db *sql.DB, collection string) domain.PaymentRepository {
	return &paymentRepository{
		database:   db,
		collection: collection,
	}
}

func (pr *paymentRepository) Create(c context.Context, payment *domain.Payment) error {
	stmt, err := pr.database.Prepare(`INSERT INTO payment (id, id_order, status, metode_pembayaran) VALUES (?, ?, ?, ?)`)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(payment.ID, payment.IDOrder, payment.Status, payment.MetodePembayaran)
	if err != nil {
		return err
	}

	return nil
}

func (pr *paymentRepository) GetById(c context.Context, id string) (domain.Payment, error) {
	var payment domain.Payment
	stmt, err := pr.database.Prepare(`SELECT id, id_order, status, metode_pembayaran FROM payment WHERE id = ?`)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		return payment, err
	}

	for rows.Next() {
		err = rows.Scan(&payment.ID, &payment.IDOrder, &payment.Status, &payment.MetodePembayaran)
		if err != nil {
			return payment, err
		}
	}

	return payment, nil
}

func (pr *paymentRepository) GetByIdOrder(c context.Context, id string) (domain.Payment, error) {
	var payment domain.Payment
	stmt, err := pr.database.Prepare(`SELECT id, id_order, status, metode_pembayaran FROM payment WHERE id_order = ?`)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		return payment, err
	}

	for rows.Next() {
		err = rows.Scan(&payment.ID, &payment.IDOrder, &payment.Status, &payment.MetodePembayaran)
		if err != nil {
			return payment, err
		}
	}

	return payment, nil
}

func (pr *paymentRepository) UpdateById(c context.Context, id string, payment domain.Payment) (domain.Payment, error) {
	stmt, err := pr.database.Prepare(`UPDATE payment SET id_order = ?, status = ?, metode_pembayaran = ? WHERE id = ?`)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(payment.IDOrder, payment.Status, payment.MetodePembayaran, id)
	if err != nil {
		return payment, err
	}

	return payment, nil
}
