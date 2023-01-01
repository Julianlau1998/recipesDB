package models

type RecipeIngredients struct {
	Recipe_ID      string `json: recipe_id`
	Ingredients_ID string `json: ingredients_id`
}

type RecipeIngredientsDB struct {
	Recipe_ID      string
	Ingredients_ID string
}

func (dbV *RecipeIngredientsDB) GetRecipeIngredients() (c RecipeIngredients) {
	c.Recipe_ID = dbV.Recipe_ID
	c.Ingredients_ID = dbV.Ingredients_ID
	return c
}
