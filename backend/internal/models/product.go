package models

import (
	"errors"
)

// Категории продуктов
type ProductCategory string

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
	EGG       ProductCategory = "egg"       // яйца
	SWEET     ProductCategory = "sweet"     // сладости, десерты
	FROZEN    ProductCategory = "frozen"    // замороженные продукты
	CANNED    ProductCategory = "canned"    // консервы
)

// Типы продуктов по степени обработки
type ProductType string

const (
	RAW_INGREDIENT ProductType = "raw"       // сырые ингредиенты
	READY_PRODUCT  ProductType = "ready"     // готовые продукты
	SEMI_FINISHED  ProductType = "semi"      // полуфабрикаты
	COMPOSITE      ProductType = "composite" // составные продукты
)

type Unit string

const (
	GRAM       Unit = "g"
	KILOGRAM   Unit = "kg"
	LITER      Unit = "l"
	MILLILITER Unit = "ml"
	PIECE      Unit = "pcs"
	TABLESPOON Unit = "tbsp"
	TEASPOON   Unit = "tsp"
	CUP        Unit = "cup"
	PACKAGE    Unit = "pack"
	BOTTLE     Unit = "bottle"
	CAN        Unit = "can"
	SLICE      Unit = "slice"
	BUNCH      Unit = "bunch" // пучок (зелень)
	CLOVE      Unit = "clove" // зубчик (чеснок)
)

type Product struct {
	ID           uint            `json:"id" gorm:"primaryKey"`
	Name         string          `json:"name" gorm:"not null;index"`
	Description  string          `json:"description" gorm:"type:text"`
	Category     ProductCategory `json:"category" gorm:"not null;index"`
	Type         ProductType     `json:"type" gorm:"not null;default:'raw'"`
	Brand        string          `json:"brand" gorm:"index"`
	ImageURL     string          `json:"image_url" gorm:"index"`
	Calories     *float64        `json:"calories" gorm:"not null;column:calories"`
	Fats         *float64        `json:"fats" gorm:"not null;column:fats"`
	Protein      *float64        `json:"protein" gorm:"not null;column:protein"`
	Carbs        *float64        `json:"carbs" gorm:"not null;column:carbs"`
	IsOrganic    bool            `json:"is_organic" gorm:"default:false"`
	IsVegetarian bool            `json:"is_vegetarian" gorm:"default:false"`
	IsVegan      bool            `json:"is_vegan" gorm:"default:false"`
	IsGlutenFree bool            `json:"is_gluten_free" gorm:"default:false"`
}

// Альтернативные названия для поиска
type ProductAlias struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	ProductID uint   `json:"product_id" gorm:"not null"`
	Alias     string `json:"alias" gorm:"not null;index"`
}

// Пищевая ценность для конкретного количества
type NutritionInfo struct {
	Calories float64
	Fats     float64
	Protein  float64
	Carbs    float64
	Weight   float64 // вес в граммах
}

// Коэффициенты для перевода единиц в граммы
var unitToGramMultiplier = map[Unit]float64{
	GRAM:       1.0,
	KILOGRAM:   1000.0,
	LITER:      1000.0, // приблизительно для воды
	MILLILITER: 1.0,    // приблизительно для воды
	PIECE:      0.0,    // зависит от продукта
	TABLESPOON: 15.0,   // приблизительно
	TEASPOON:   5.0,    // приблизительно
	CUP:        240.0,  // приблизительно
	SLICE:      0.0,    // зависит от продукта
	BUNCH:      0.0,    // зависит от продукта
	CLOVE:      3.0,    // средний зубчик чеснока
}

func ValidateProductCategory(category ProductCategory) error {
	switch category {
	case FRUIT, VEGETABLE, MEAT, FISH, DAIRY, SAUCE, GRAIN, OIL, BEVERAGE, SNACK, SPICE, EGG, SWEET, FROZEN, CANNED:
		return nil
	default:
		return errors.New("invalid product category")
	}
}

func ValidateProductType(productType ProductType) error {
	switch productType {
	case RAW_INGREDIENT, READY_PRODUCT, SEMI_FINISHED, COMPOSITE:
		return nil
	default:
		return errors.New("invalid product type")
	}
}

func (p *Product) HasNutritionInfo() bool {
	return p.Calories != nil && p.Fats != nil && p.Protein != nil && p.Carbs != nil
}

func (p *Product) IsReadyToEat() bool {
	return p.Type == READY_PRODUCT
}

// Рассчитать пищевую ценность для конкретного количества
func (p *Product) CalculateNutrition(amount float64, unit Unit) (*NutritionInfo, error) {
	if !p.HasNutritionInfo() {
		return nil, errors.New("nutrition information not available for this product")
	}

	// Конвертируем в граммы
	weightInGrams, err := p.convertToGrams(amount, unit)
	if err != nil {
		return nil, err
	}

	// Рассчитываем пищевую ценность пропорционально
	factor := weightInGrams / 100.0 // пищевая ценность указана на 100г

	return &NutritionInfo{
		Calories: *p.Calories * factor,
		Fats:     *p.Fats * factor,
		Protein:  *p.Protein * factor,
		Carbs:    *p.Carbs * factor,
		Weight:   weightInGrams,
	}, nil
}

// Конвертация в граммы
func (p *Product) convertToGrams(amount float64, unit Unit) (float64, error) {
	multiplier, exists := unitToGramMultiplier[unit]
	if !exists {
		return 0, errors.New("unsupported unit")
	}

	// Для единиц, которые зависят от продукта
	if multiplier == 0.0 {
		return p.getProductSpecificWeight(amount, unit)
	}

	return amount * multiplier, nil
}

// Вес для продукто-специфичных единиц
func (p *Product) getProductSpecificWeight(amount float64, unit Unit) (float64, error) {
	switch unit {
	case PIECE:
		return p.getPieceWeight(amount)
	case SLICE:
		return p.getSliceWeight(amount)
	case BUNCH:
		return p.getBunchWeight(amount)
	default:
		return 0, errors.New("unsupported product-specific unit")
	}
}

// Примерный вес одной штуки разных продуктов
func (p *Product) getPieceWeight(amount float64) (float64, error) {
	averageWeights := map[ProductCategory]float64{
		EGG:       60.0,  // среднее яйцо
		FRUIT:     150.0, // среднее яблоко
		VEGETABLE: 100.0, // средний помидор
	}

	if weight, exists := averageWeights[p.Category]; exists {
		return amount * weight, nil
	}

	return 0, errors.New("cannot determine piece weight for this product category")
}

// Примерный вес одного кусочка
func (p *Product) getSliceWeight(amount float64) (float64, error) {
	averageSliceWeights := map[ProductCategory]float64{
		GRAIN: 25.0, // кусочек хлеба
		DAIRY: 20.0, // кусочек сыра
	}

	if weight, exists := averageSliceWeights[p.Category]; exists {
		return amount * weight, nil
	}

	return 0, errors.New("cannot determine slice weight for this product category")
}

// Примерный вес пучка зелени
func (p *Product) getBunchWeight(amount float64) (float64, error) {
	if p.Category == VEGETABLE {
		return amount * 50.0, nil // средний пучок зелени
	}

	return 0, errors.New("bunch weight only applicable to vegetables")
}
