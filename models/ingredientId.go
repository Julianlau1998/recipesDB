package models

type IngredientId struct {
	RecipeId     string `json:"recipeId"`
	IngredientId string `json:"ingredientId"`
}

type IngredientIdDB struct {
	RecipeId     string
	IngredientId string
}

func (dbV *IngredientIdDB) GetIngredientsId() (c IngredientId) {
	c.RecipeId = dbV.RecipeId
	c.IngredientId = dbV.IngredientId
	return c
}
