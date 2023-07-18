package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"waroeng_pgn1/domain"
)

type unitTypeRepository struct {
	database   *sql.DB
	collection string
}

func NewUnitTypeRepository(db *sql.DB, collection string) domain.UnitTypeRepository {
	return &unitTypeRepository{
		database:   db,
		collection: collection,
	}
}

func (utr *unitTypeRepository) Create(c context.Context, unitType *domain.UnitType) error {
	stmt, err := utr.database.Prepare(`INSERT INTO unit_type (id, unit_type_name) VALUES (?, ?)`)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()
	result, err := stmt.Exec(unitType.ID, unitType.UnitTypeName)
	if err != nil {
		return err
	} else if result != nil {
		return nil
	}
	return errors.New("error while creating unit type")
}

func (utr *unitTypeRepository) GetById(c context.Context, id string) (domain.UnitType, error) {
	var unitType domain.UnitType
	stmt, err := utr.database.Prepare(`SELECT id, unit_type_name FROM unit_type WHERE id = ?`)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()
	row := stmt.QueryRow(id)
	err = row.Scan(&unitType.ID, &unitType.UnitTypeName)
	if err != nil {
		fmt.Println(err)
		return unitType, err
	}
	return unitType, nil
}
