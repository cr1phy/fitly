package database

import (
	"database/sql"

	"github.com/cr1phy/fitly/internal/entity"
)

func GetAllProductsFromFilter(db *sql.DB, filter string) ([]entity.Product, error) {
	var products []entity.Product

	rows, err := db.Query(`SELECT * FROM products WHERE name || '%' || $1 || '%'`, filter)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if err := rows.Scan(&products); err != nil {
		return nil, err
	}

	return products, nil
}
