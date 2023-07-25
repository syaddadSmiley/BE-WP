package repository

import (
	"context"
	"database/sql"
	"errors"

	"waroeng_pgn1/domain"
)

type courierServiceRepository struct {
	database   *sql.DB
	collection string
}

func NewCourierServiceRepository(db *sql.DB, collection string) domain.CourierServiceRepository {
	return &courierServiceRepository{
		database:   db,
		collection: collection,
	}
}

func (csr *courierServiceRepository) Create(c context.Context, courierService *domain.CourierService) error {
	stmt, err := csr.database.Prepare(`INSERT INTO courier_service (id, id_courier, id_address, service_name, receipt_service_number, description, price_service, estimation_day) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()
	result, err := stmt.Exec(courierService.ID, courierService.IDCourier, courierService.IDAddress, courierService.ServiceName, courierService.ReceiptServiceNumber, courierService.Description, courierService.PriceService, courierService.EstimationDay)
	if err != nil {
		return err
	} else if result != nil {
		return nil
	}
	return errors.New("error while creating courier service")
}

func (csr *courierServiceRepository) GetById(c context.Context, id string) (domain.CourierService, error) {
	var courierService domain.CourierService
	stmt, err := csr.database.Prepare(`SELECT id, id_courier, id_address, service_name, receipt_service_number, description, price_service, estimation_day FROM courier_service WHERE id = ?`)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&courierService.ID, &courierService.IDCourier, &courierService.IDAddress, &courierService.ServiceName, &courierService.ReceiptServiceNumber, &courierService.Description, &courierService.PriceService, &courierService.EstimationDay)
	if err != nil {
		return courierService, err
	}
	return courierService, nil
}
