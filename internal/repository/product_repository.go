package repository

import (
	"context"
	"database/sql"
	"errors"

	"waroeng_pgn1/domain"
)

type productRepository struct {
	database   *sql.DB
	collection string
}

func NewProductRepository(db *sql.DB, collection string) domain.ProductRepository {
	return &productRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *productRepository) Create(c context.Context, product *domain.Product) error {
	stmt, err := ur.database.Prepare(`INSERT INTO product (id, id_price_type, id_unit_type, name_product, price, description_product, discount, sold_amount, stock,  location) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()
	result, err := stmt.Exec(product.ID, product.IDPriceType, product.IDUnitType, product.NameProduct, product.Price, product.DescriptionProduct, product.Discount, product.SoldAmount, product.Stock, product.Location)
	if err != nil {
		return err
	} else if result != nil {
		return nil
	}
	return errors.New("error while creating product")
}

func (ur *productRepository) GetById(c context.Context, id string) (domain.Product, error) {
	var product domain.Product
	stmt, err := ur.database.Prepare(`SELECT id, id_price_type, id_unit_type, name_product, price, description_product, discount, sold_amount, stock, location FROM product WHERE id = ?`)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&product.ID, &product.IDPriceType, &product.IDUnitType, &product.NameProduct, &product.Price, &product.DescriptionProduct, &product.Discount, &product.SoldAmount, &product.Stock, &product.Location)
	if err != nil {
		return product, err
	} else if product.ID == "" {
		return product, errors.New("product not found")
	}

	return product, nil
}

func (ur *productRepository) GetAll(c context.Context) ([]domain.Product, error) {
	var products []domain.Product
	stmt, err := ur.database.Prepare(`SELECT id, id_price_type, id_unit_type, name_product, price, description_product, discount, sold_amount, stock, location FROM product`)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return products, err
	}

	for rows.Next() {
		var product domain.Product
		err = rows.Scan(&product.ID, &product.IDPriceType, &product.IDUnitType, &product.NameProduct, &product.Price, &product.DescriptionProduct, &product.Discount, &product.SoldAmount, &product.Stock, &product.Location)
		if err != nil {
			return products, err
		}

		products = append(products, product)
	}

	return products, nil
}
