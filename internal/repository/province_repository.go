package repository

import (
	"context"
	"database/sql"

	"waroeng_pgn1/domain"
)

type provinceRepository struct {
	database   *sql.DB
	collection string
}

func NewProvinceRepository(db *sql.DB, collection string) domain.ProvinceRepository {
	return &provinceRepository{
		database:   db,
		collection: collection,
	}
}

func (pr *provinceRepository) InsertProvince(ctx context.Context, province []domain.ProvinceResult) ([]domain.ProvinceResult, error) {
	var provinces []domain.ProvinceResult
	stmt, err := pr.database.Prepare(`INSERT INTO province (province_id, province) VALUES (?, ?)`)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	for _, v := range province {
		result, err := stmt.Exec(v.ProvinceID, v.Province)
		if err != nil {
			return provinces, err
		} else if result != nil {
			provinces = append(provinces, v)
		}
	}

	return provinces, nil
}

func (pr *provinceRepository) GetProvince(ctx context.Context) ([]domain.ProvinceResult, error) {
	var provinces []domain.ProvinceResult
	stmt, err := pr.database.Prepare(`SELECT province_id, province FROM province`)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return provinces, err
	}

	for rows.Next() {
		var province domain.ProvinceResult
		err := rows.Scan(&province.ProvinceID, &province.Province)
		if err != nil {
			return provinces, err
		}
		provinces = append(provinces, province)
	}

	return provinces, nil
}
