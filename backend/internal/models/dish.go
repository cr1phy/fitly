package models

import (
	"errors"

	"gorm.io/gorm"
)

type DishCategory string

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
	DRINK     DishCategory = "drink"
)

type Ingredient struct {
	gorm.Model
	ID        uint    `gorm:"primaryKey"`
	DishID    uint    `gorm:"not null"`
	ProductID uint    `gorm:"not null"`
	Product   Product `gorm:"foreignKey:ProductID"`
	Amount    float64 `gorm:"not null;check:amount > 0"`
	Unit      Unit    `gorm:"not null;default:'g'"`

	// Кулинарная обработка
	Preparation string `gorm:"comment:'способ подготовки'"`
	IsOptional  bool   `gorm:"default:false"`
	Notes       string `gorm:"comment:'заметки пользователя'"`
}

type Dish struct {
	gorm.Model
	ID          uint         `gorm:"primaryKey"`
	Name        string       `gorm:"not null;index"`
	Category    DishCategory `gorm:"not null;index"`
	Description string

	// Кулинарная информация
	CookingTime     int    `gorm:"comment:'время готовки в минутах'"`
	PreparationTime int    `gorm:"comment:'время подготовки в минутах'"`
	Servings        int    `gorm:"default:1;check:servings > 0"`
	Instructions    string `gorm:"type:text"`

	// Связи
	Ingredients []Ingredient `gorm:"foreignKey:DishID"`

	// Пользовательские данные
	UserID   *uint   `gorm:"index"`
	IsPublic bool    `gorm:"default:false"`
	Rating   float64 `gorm:"default:0;check:rating >= 0 AND rating <= 5"`
}

// Суммарная пищевая ценность блюда
type DishNutrition struct {
	TotalCalories float64
	TotalFats     float64
	TotalProtein  float64
	TotalCarbs    float64
	TotalWeight   float64
	PerServing    NutritionInfo
}

// Валидация категории блюда
func ValidateDishCategory(category DishCategory) error {
	switch category {
	case SALAD, SOUP, MAIN, DESSERT, SANDWICH, WRAP, PIZZA, PASTA, BURGER, BREAKFAST, DRINK:
		return nil
	default:
		return errors.New("invalid dish category")
	}
}

// Методы для удобства
func (d *Dish) TotalTime() int {
	return d.CookingTime + d.PreparationTime
}

func (d *Dish) IsQuickMeal() bool {
	return d.TotalTime() <= 30 // быстрое блюдо - до 30 минут
}

func (d *Dish) HasInstructions() bool {
	return d.Instructions != ""
}

// Рассчитать общую пищевую ценность блюда
func (d *Dish) CalculateTotalNutrition() (*DishNutrition, error) {
	if len(d.Ingredients) == 0 {
		return nil, errors.New("dish has no ingredients")
	}

	totalNutrition := &DishNutrition{}

	for _, ingredient := range d.Ingredients {
		// Пропускаем опциональные ингредиенты без пищевой ценности
		if ingredient.IsOptional && !ingredient.Product.HasNutritionInfo() {
			continue
		}

		nutrition, err := ingredient.Product.CalculateNutrition(ingredient.Amount, ingredient.Unit)
		if err != nil {
			// Если не можем рассчитать для какого-то ингредиента, пропускаем
			continue
		}

		totalNutrition.TotalCalories += nutrition.Calories
		totalNutrition.TotalFats += nutrition.Fats
		totalNutrition.TotalProtein += nutrition.Protein
		totalNutrition.TotalCarbs += nutrition.Carbs
		totalNutrition.TotalWeight += nutrition.Weight
	}

	// Рассчитываем на порцию
	if d.Servings > 0 {
		totalNutrition.PerServing = NutritionInfo{
			Calories: totalNutrition.TotalCalories / float64(d.Servings),
			Fats:     totalNutrition.TotalFats / float64(d.Servings),
			Protein:  totalNutrition.TotalProtein / float64(d.Servings),
			Carbs:    totalNutrition.TotalCarbs / float64(d.Servings),
			Weight:   totalNutrition.TotalWeight / float64(d.Servings),
		}
	}

	return totalNutrition, nil
}

// Рассчитать пищевую ценность на 100г готового блюда
func (d *Dish) CalculateNutritionPer100g() (*NutritionInfo, error) {
	totalNutrition, err := d.CalculateTotalNutrition()
	if err != nil {
		return nil, err
	}

	if totalNutrition.TotalWeight == 0 {
		return nil, errors.New("cannot calculate nutrition per 100g: total weight is zero")
	}

	factor := 100.0 / totalNutrition.TotalWeight

	return &NutritionInfo{
		Calories: totalNutrition.TotalCalories * factor,
		Fats:     totalNutrition.TotalFats * factor,
		Protein:  totalNutrition.TotalProtein * factor,
		Carbs:    totalNutrition.TotalCarbs * factor,
		Weight:   100.0,
	}, nil
}

// Найти ингредиенты определенной категории
func (d *Dish) GetIngredientsByCategory(category ProductCategory) []Ingredient {
	var ingredients []Ingredient
	for _, ingredient := range d.Ingredients {
		if ingredient.Product.Category == category {
			ingredients = append(ingredients, ingredient)
		}
	}
	return ingredients
}

// Проверить, является ли блюдо вегетарианским
func (d *Dish) IsVegetarian() bool {
	for _, ingredient := range d.Ingredients {
		if ingredient.Product.Category == MEAT || ingredient.Product.Category == FISH {
			return false
		}
	}
	return true
}

// Проверить, является ли блюдо веганским
func (d *Dish) IsVegan() bool {
	for _, ingredient := range d.Ingredients {
		if ingredient.Product.Category == MEAT ||
			ingredient.Product.Category == FISH ||
			ingredient.Product.Category == DAIRY ||
			ingredient.Product.Category == EGG {
			return false
		}
	}
	return true
}

// Проверить, содержит ли блюдо глютен
func (d *Dish) IsGlutenFree() bool {
	for _, ingredient := range d.Ingredients {
		if ingredient.Product.Category == GRAIN && !ingredient.Product.IsGlutenFree {
			return false
		}
	}
	return true
}

// Получить основные ингредиенты (не опциональные)
func (d *Dish) GetEssentialIngredients() []Ingredient {
	var essentialIngredients []Ingredient
	for _, ingredient := range d.Ingredients {
		if !ingredient.IsOptional {
			essentialIngredients = append(essentialIngredients, ingredient)
		}
	}
	return essentialIngredients
}

// Оценить сложность приготовления блюда
func (d *Dish) GetComplexityLevel() string {
	// Базовая сложность основана на количестве ингредиентов
	ingredientCount := len(d.GetEssentialIngredients())
	totalTime := d.TotalTime()

	complexityScore := 0

	// Фактор времени
	if totalTime > 120 {
		complexityScore += 3
	} else if totalTime > 60 {
		complexityScore += 2
	} else if totalTime > 30 {
		complexityScore += 1
	}

	// Фактор количества ингредиентов
	if ingredientCount > 15 {
		complexityScore += 3
	} else if ingredientCount > 10 {
		complexityScore += 2
	} else if ingredientCount > 5 {
		complexityScore += 1
	}

	// Фактор сложности обработки
	for _, ingredient := range d.Ingredients {
		if ingredient.Preparation != "" {
			complexityScore += 1
		}
	}

	// Определяем уровень сложности
	if complexityScore <= 2 {
		return "Простое"
	} else if complexityScore <= 5 {
		return "Среднее"
	} else if complexityScore <= 8 {
		return "Сложное"
	} else {
		return "Очень сложное"
	}
}

// Рассчитать примерную стоимость блюда (если есть данные о ценах)
func (d *Dish) EstimateCost(productPrices map[uint]float64) (float64, error) {
	if len(productPrices) == 0 {
		return 0, errors.New("no price information available")
	}

	totalCost := 0.0

	for _, ingredient := range d.Ingredients {
		if price, exists := productPrices[ingredient.ProductID]; exists {
			// Примерный расчет стоимости на основе веса
			weightInGrams, err := ingredient.Product.convertToGrams(ingredient.Amount, ingredient.Unit)
			if err != nil {
				continue // пропускаем ингредиенты с неизвестным весом
			}

			// Предполагаем, что цена указана за 100г
			ingredientCost := (weightInGrams / 100.0) * price
			totalCost += ingredientCost
		}
	}

	return totalCost, nil
}
