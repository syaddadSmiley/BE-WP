package repository

import (
	"context"
	"database/sql"
	"waroeng_pgn1/domain"
)

type addressesRepository struct {
	database   *sql.DB
	collection string
}

func NewAddressesRepository(db *sql.DB, collection string) domain.AddressesRepository {
	return &addressesRepository{
		database:   db,
		collection: collection,
	}
}

type Addresses struct {
	ID           string `json:"id"`
	IDUser       string `json:"id_user"`
	LabelAddress string `json:"label_address"`
	Address      string `json:"address"`
	City         string `json:"city"`
	Province     string `json:"province"`
	PostalCode   string `json:"postal_code"`
	Latitude     string `json:"latitude"`
	Longitude    string `json:"longitude"`
	Note         string `json:"note"`
	IsDefault    string `json:"is_default"`
}

func (ur *addressesRepository) Create(c context.Context, addresses *domain.Addresses) error {
	stmt, err := ur.database.Prepare(`INSERT INTO addresses (id, id_user, label_address, address, city, province, postal_code, latitude, longitude, note, is_default) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()
	_, err = stmt.Exec(addresses.ID, addresses.IDUser, addresses.LabelAddress, addresses.Address, addresses.City, addresses.Province, addresses.PostalCode, addresses.Latitude, addresses.Longitude, addresses.Note, addresses.IsDefault)
	if err != nil {
		return err
	}

	return nil
}

func (ur *addressesRepository) GetById(c context.Context, id string) (domain.Addresses, error) {
	var addresses domain.Addresses
	stmt, err := ur.database.Prepare(`SELECT id, id_user, label_address, address, city, province, postal_code, latitude, longitude, note, is_default FROM addresses WHERE id = ?`)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&addresses.ID, &addresses.IDUser, &addresses.LabelAddress, &addresses.Address, &addresses.City, &addresses.Province, &addresses.PostalCode, &addresses.Latitude, &addresses.Longitude, &addresses.Note, &addresses.IsDefault)
	if err != nil {
		return addresses, err
	}

	return addresses, nil
}

func (ur *addressesRepository) GetByIdUser(c context.Context, id string) ([]domain.Addresses, error) {
	var addresses []domain.Addresses
	stmt, err := ur.database.Prepare(`SELECT id, id_user, label_address, address, city, province, postal_code, latitude, longitude, note, is_default FROM addresses WHERE id_user = ?`)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		return addresses, err
	}

	for rows.Next() {
		var address domain.Addresses
		err := rows.Scan(&address.ID, &address.IDUser, &address.LabelAddress, &address.Address, &address.City, &address.Province, &address.PostalCode, &address.Latitude, &address.Longitude, &address.Note, &address.IsDefault)
		if err != nil {
			return addresses, err
		}

		addresses = append(addresses, address)
	}

	return addresses, nil
}

func (ur *addressesRepository) UpdateById(c context.Context, id string, addresses domain.Addresses) (domain.Addresses, error) {
	var addressesResult domain.Addresses
	stmt, err := ur.database.Prepare(`UPDATE addresses SET label_address = ?, address = ?, city = ?, province = ?, postal_code = ?, latitude = ?, longitude = ?, note = ?, is_default = ? WHERE id = ?`)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(addresses.LabelAddress, addresses.Address, addresses.City, addresses.Province, addresses.PostalCode, addresses.Latitude, addresses.Longitude, addresses.Note, addresses.IsDefault, id)
	if err != nil {
		return addressesResult, err
	}

	return addressesResult, nil
}
