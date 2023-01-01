package models

type Ingredient struct {
	IngredientId string `json:"ingredientId"`
	Ingredient   string `json: ingredient`
}

type IngredientsDB struct {
	IngredientId string
	Ingredient   string
}

func (dbV *IngredientsDB) GetIngredients() (c Ingredient) {
	c.IngredientId = dbV.IngredientId
	c.Ingredient = dbV.Ingredient
	return c
}
