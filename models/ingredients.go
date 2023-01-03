package models

import (
	"database/sql"
	"recipes/utility"
)

type Ingredient struct {
	IngredientId string `json:"ingredientId"`
	Measurement  string `json:"measurement"`
	Ingredient   string `json:"ingredient"`
}

type IngredientsDB struct {
	IngredientId string
	Measurement  sql.NullString
	Ingredient   string
}

func (dbV *IngredientsDB) GetIngredients() (c Ingredient) {
	c.IngredientId = dbV.IngredientId
	c.Measurement = utility.GetStringValue(dbV.Measurement)
	c.Ingredient = dbV.Ingredient
	return c
}
