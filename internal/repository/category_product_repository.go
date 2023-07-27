package repository

import (
	"context"
	"database/sql"
	"errors"

	"waroeng_pgn1/domain"
)

type categoryProductRepository struct {
	database   *sql.DB
	collection string
}

func NewCategoryProductRepository(db *sql.DB, collection string) domain.CategoryProductRepository {
	return &categoryProductRepository{
		database:   db,
		collection: collection,
	}
}

func (cpr *categoryProductRepository) Create(c context.Context, category *domain.CategoryProduct) error {
	stmt, err := cpr.database.Prepare(`INSERT INTO category_product (id, name_category, type_category) VALUES (?, ?, ?)`)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()
	result, err := stmt.Exec(category.ID, category.NameCategory, category.TypeCategory)
	if err != nil {
		return err
	} else if result != nil {
		return nil
	}
	return errors.New("error while creating category")
}

func (cpr *categoryProductRepository) GetById(c context.Context, id string) (domain.CategoryProduct, error) {
	var category domain.CategoryProduct
	stmt, err := cpr.database.Prepare(`SELECT id, name_category, type_category FROM category_product WHERE id = ?`)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()
	result, err := stmt.Query(id)
	if err != nil {
		return category, err
	} else if result != nil {
		for result.Next() {
			err := result.Scan(&category.ID, &category.NameCategory, &category.TypeCategory)
			if err != nil {
				return category, err
			}
		}
		return category, nil
	}
	return category, errors.New("error while getting category")
}

func (cpr *categoryProductRepository) GetAll(c context.Context) ([]domain.CategoryProduct, error) {
	var category []domain.CategoryProduct
	stmt, err := cpr.database.Prepare(`SELECT id, name_category, type_category FROM category_product`)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()
	result, err := stmt.Query()
	if err != nil {
		return category, err
	} else if result != nil {
		for result.Next() {
			var cat domain.CategoryProduct
			err := result.Scan(&cat.ID, &cat.NameCategory, &cat.TypeCategory)
			if err != nil {
				return category, err
			}
			category = append(category, cat)
		}
		return category, nil
	}
	return category, errors.New("error while getting category")
}
