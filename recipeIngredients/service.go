package recipeIngredients

import (
	"recipes/models"
)

/*import (
	"recipes/models"
)*/

type Service struct {
	recipeIngredientsRepo Repository
}

func NewService(recipeIngredientsRepository Repository) Service {
	return Service{recipeIngredientsRepo: recipeIngredientsRepository}
}

func (s *Service) GetRecipeIngredients() ([]models.RecipeIngredients, error) {
	return s.recipeIngredientsRepo.GetRecipeIngredients()
}

func (s *Service) PostRecipeIngredients(recipeId string, ingredientId string) error {
	return s.recipeIngredientsRepo.PostRecipeIngredients(recipeId, ingredientId)
}
