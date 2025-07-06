package models

import "errors"

const (
	FRUIT     ProductCategory = "fruit"     // фрукты
	VEGETABLE ProductCategory = "vegetable" // овощи
	MEAT      ProductCategory = "meat"      // мясо и мясные изделия
	FISH      ProductCategory = "fish"      // рыба и морепродукты
	DAIRY     ProductCategory = "dairy"     // молочные и сырные продукты
	SAUCE     ProductCategory = "sauce"     // соусы, приправы-жидкости
	GRAIN     ProductCategory = "grain"     // хлеб, крупы, макароны
	OIL       ProductCategory = "oil"       // растительные и сливочные масла
	BEVERAGE  ProductCategory = "beverage"  // напитки и соки
	SNACK     ProductCategory = "snack"     // снеки и орехи
	SPICE     ProductCategory = "spice"     // соль, сахар, специи, сухие приправы
)

type ProductCategory string

type Product struct {
	Id          int
	Category    ProductCategory `gorm:"embedded"`
	Name        string
	Description string
	Calories    float64
	Fats        float64
	Protein     float64
	Carbs       float64
}

func ValidateProductCategory(category ProductCategory) error {
	switch category {
	case FRUIT, VEGETABLE, MEAT, FISH, DAIRY, SAUCE, GRAIN, OIL, BEVERAGE, SNACK, SPICE:
		return nil
	default:
		return errors.New("invalid product's category")
	}
}
