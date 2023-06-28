package repository

import (
	"context"
	"database/sql"

	"waroeng_pgn1/domain"
)

type cityRepository struct {
	database   *sql.DB
	collection string
}

func NewCityRepository(db *sql.DB, collection string) domain.CityRepository {
	return &cityRepository{
		database:   db,
		collection: collection,
	}
}

func (pr *cityRepository) InsertCity(ctx context.Context, city []domain.CityResult) ([]domain.CityResult, error) {
	var cities []domain.CityResult
	stmt, err := pr.database.Prepare(`INSERT INTO city (city_id, province_id, city_name, type, postal_code) VALUES (?, ?, ?, ?, ?)`)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	for _, v := range city {
		result, err := stmt.Exec(v.CityID, v.ProvinceID, v.CityName, v.Type, v.PostalCode)
		if err != nil {
			return cities, err
		} else if result != nil {
			cities = append(cities, v)
		}
	}

	return cities, nil
}

func (pr *cityRepository) GetCityByProvince(ctx context.Context, provinceID string) ([]domain.CityResult, error) {
	var cities []domain.CityResult
	stmt, err := pr.database.Prepare(`SELECT city_id, province_id, city_name, type, postal_code FROM city WHERE province_id = ?`)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()
	rows, err := stmt.Query(provinceID)
	if err != nil {
		return cities, err
	}

	for rows.Next() {
		var city domain.CityResult
		err := rows.Scan(&city.CityID, &city.ProvinceID, &city.CityName, &city.Type, &city.PostalCode)
		if err != nil {
			return cities, err
		}
		cities = append(cities, city)
	}

	return cities, nil
}
