package models

const (
	SALAD     DishCategory = "salad"
	SOUP      DishCategory = "soup"
	MAIN      DishCategory = "main"
	DESSERT   DishCategory = "dessert"
	SANDWICH  DishCategory = "sandwich"
	WRAP      DishCategory = "wrap"
	PIZZA     DishCategory = "pizza"
	PASTA     DishCategory = "pasta"
	BURGER    DishCategory = "burger"
	BREAKFAST DishCategory = "breakfast"
)

type DishCategory string
type Ingredient struct {
	product *Product
	weight  float64
}

type Dish struct {
	id          uint         `gorm:"primaryKey"`
	category    DishCategory `gorm:"embedded"`
	name        string
	description string
	ingredients []Ingredient `gorm:"embedded"`
}
