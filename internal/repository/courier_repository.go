package repository

import (
	"context"
	"database/sql"
	"errors"

	"waroeng_pgn1/domain"
)

type courierRepository struct {
	database   *sql.DB
	collection string
}

func NewCourierRepository(db *sql.DB, collection string) domain.CourierRepository {
	return &courierRepository{
		database:   db,
		collection: collection,
	}
}

func (cr *courierRepository) Create(c context.Context, courier *domain.Courier) error {
	stmt, err := cr.database.Prepare(`INSERT INTO courier (id, courier_name, is_available) VALUES (?, ?, 1)`)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()
	result, err := stmt.Exec(courier.ID, courier.CourierName)
	if err != nil {
		return err
	} else if result != nil {
		return nil
	}
	return errors.New("error while creating courier")
}

func (cr *courierRepository) GetAll(c context.Context) ([]domain.Courier, error) {
	var couriers []domain.Courier
	stmt, err := cr.database.Prepare(`SELECT id, courier_name FROM courier WHERE is_available = 1`)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return couriers, err
	}
	defer rows.Close()

	for rows.Next() {
		var courier domain.Courier
		err := rows.Scan(&courier.ID, &courier.CourierName)
		if err != nil {
			return couriers, err
		}
		couriers = append(couriers, courier)
	}

	return couriers, nil
}

func (cr *courierRepository) GetIdCityByName(c context.Context, name string) (string, error) {
	var idCity string
	stmt, err := cr.database.Prepare(`SELECT city_id FROM city WHERE city_name = ?`)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()
	row := stmt.QueryRow(name)
	err = row.Scan(&idCity)
	if err != nil {
		return idCity, err
	}
	return idCity, nil
}
