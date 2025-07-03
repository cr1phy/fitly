package database

const (
	GET_PRODUCTS                  = `SELECT * FROM products WHERE name LIKE '%' || $1 || '%'`
	INSERT_PRODUCT_WITH_RETURNING = `INSERT INTO products (category, name, description, calories, fats, protein, carbs) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
)
