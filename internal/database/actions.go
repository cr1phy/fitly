package database

import (
	"database/sql"

	"github.com/cr1phy/fitly/internal/entity"
)

func GetAllProductsFromFilter(db *sql.DB, filter string) ([]entity.Product, error) {
	var products []entity.Product

	rows, err := db.Query(GET_PRODUCTS, filter)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if err := rows.Scan(&products); err != nil {
		return nil, err
	}

	return products, nil
}

func CreateProduct(db *sql.DB, data entity.Product) (int, error) {
	var id int

	result, err := db.Query(INSERT_PRODUCT_WITH_RETURNING, data.Category, data.Name, data.Description, data.Calories, data.Fats, data.Protein, data.Carbs)
	if err != nil {
		return -1, err
	}
	defer result.Close()
	result.Scan(&id)

	return id, nil
}
