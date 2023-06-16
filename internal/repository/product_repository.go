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
	stmt, err := ur.database.Prepare(`INSERT INTO products (id, id_price_type, id_unit_type, name_product, price, description_product, discount, sold_amount, stock, stock_type, location) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()
	result, err := stmt.Exec(product.ID, product.IDPriceType, product.IDUnitType, product.NameProduct, product.Price, product.DescriptionProduct, product.Discount, product.SoldAmount, product.Stock, product.StockType, product.Location)
	if err != nil {
		return err
	} else if result != nil {
		return nil
	}
	return errors.New("error while creating product")
}

func (ur *productRepository) GetById(c context.Context, id string) (domain.Product, error) {
	var product domain.Product
	return product, nil
}
