package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

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

func (pr *productRepository) Create(c context.Context, product *domain.Product) error {
	stmt, err := pr.database.Prepare(`INSERT INTO product (id, unit_type_value, id_category, id_unit_type, name_product, price, description_product, discount, sold_amount, stock,  location) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()
	result, err := stmt.Exec(product.ID, product.UnitTypeValue, product.IDCategory, product.IDUnitType, product.NameProduct, product.Price, product.DescriptionProduct, product.Discount, product.SoldAmount, product.Stock, product.Location)
	if err != nil {
		return err
	} else if result != nil {
		return nil
	}
	return errors.New("error while creating product")
}

func (pr *productRepository) GetById(c context.Context, id string) (domain.Product, error) {
	var product domain.Product
	stmt, err := pr.database.Prepare(`SELECT id, unit_type_value, id_unit_type, name_product, price, description_product, discount, sold_amount, stock, location FROM product WHERE id = ?`)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&product.ID, &product.UnitTypeValue, &product.IDUnitType, &product.NameProduct, &product.Price, &product.DescriptionProduct, &product.Discount, &product.SoldAmount, &product.Stock, &product.Location)
	if err != nil {
		return product, err
	} else if product.ID == "" {
		return product, errors.New("product not found")
	}

	return product, nil
}

func (pr *productRepository) GetAll(c context.Context) ([]domain.ProductResponse, error) {
	var products []domain.ProductResponse
	stmt, err := pr.database.Prepare(`SELECT 
	p.id, 
	p.unit_type_value, 
	cp.name_category,
	p.name_product, 
	p.price, 
	p.description_product, 
	p.discount, p.sold_amount, p.stock, p.location, ut.unit_type_name FROM product p INNER JOIN unit_type ut ON p.id_unit_type = ut.id RIGHT JOIN category_product cp ON cp.id = p.id_category`)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return products, err
	}

	for rows.Next() {
		var product domain.ProductResponse
		err = rows.Scan(&product.ID, &product.UnitTypeValue, &product.NameCategory, &product.NameProduct, &product.Price, &product.DescriptionProduct, &product.Discount, &product.SoldAmount, &product.Stock, &product.Location, &product.UnitTypeValue)
		if err != nil {
			return products, err
		}

		products = append(products, product)
	}

	return products, nil
}

func (pr *productRepository) GetAllByCity(c context.Context, city string) ([]domain.ProductResponse, error) {
	fmt.Println("city", city)
	var products []domain.ProductResponse
	stmt, err := pr.database.Prepare(`SELECT 
	p.id, 
	p.unit_type_value, 
	cp.type_category,
	cp.name_category,
	p.name_product, 
	p.price, 
	p.description_product, 
	p.discount, p.sold_amount, p.stock, p.location, ut.unit_type_name FROM product p INNER JOIN unit_type ut ON p.id_unit_type = ut.id RIGHT JOIN category_product cp ON cp.id = p.id_category WHERE p.location = ?;`)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()
	rows, err := stmt.Query(city)
	if err != nil {
		return products, err
	}

	for rows.Next() {
		var product domain.ProductResponse
		err = rows.Scan(&product.ID, &product.UnitTypeValue, &product.TypeCategory, &product.NameCategory, &product.NameProduct, &product.Price, &product.DescriptionProduct, &product.Discount, &product.SoldAmount, &product.Stock, &product.Location, &product.UnitTypeName)
		if err != nil {
			return products, err
		}

		products = append(products, product)
	}

	return products, nil
}

func (pr *productRepository) GetAllByCategory(c context.Context, category string, city string) ([]domain.ProductResponse, error) {
	var products []domain.ProductResponse
	stmt, err := pr.database.Prepare(`SELECT 
	p.id, 
	p.unit_type_value, 
	cp.type_category,
	cp.name_category,
	p.name_product, 
	p.price, 
	p.description_product, 
	p.discount, p.sold_amount, p.stock, p.location, ut.unit_type_name FROM product p INNER JOIN unit_type ut ON p.id_unit_type = ut.id RIGHT JOIN category_product cp ON cp.id = p.id_category WHERE p.location = ? AND cp.type_category = ?;`)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()
	rows, err := stmt.Query(city, category)
	if err != nil {
		return products, err
	}

	for rows.Next() {
		var product domain.ProductResponse
		err = rows.Scan(&product.ID, &product.UnitTypeValue, &product.TypeCategory, &product.NameCategory, &product.NameProduct, &product.Price, &product.DescriptionProduct, &product.Discount, &product.SoldAmount, &product.Stock, &product.Location, &product.UnitTypeName)
		if err != nil {
			return products, err
		}

		products = append(products, product)
	}

	return products, nil
}
