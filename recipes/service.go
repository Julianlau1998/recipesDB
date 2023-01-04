package recipes

import (
	"fmt"
	uuid "github.com/nu7hatch/gouuid"
	log "github.com/sirupsen/logrus"
	"recipes/categories"
	"recipes/ingredients"
	"recipes/models"
	"recipes/recipeIngredients"
)

type Service struct {
	recipeRepo               Repository
	ingredientsService       ingredients.Service
	recipeIngredientsService recipeIngredients.Service
	categoryService          categories.Service
}

func NewService(
	recipeRepository Repository,
	ingredientsService ingredients.Service,
	recipeIngredientsService recipeIngredients.Service,
	categoryService categories.Service,
) Service {
	return Service{
		recipeRepo:               recipeRepository,
		ingredientsService:       ingredientsService,
		recipeIngredientsService: recipeIngredientsService,
		categoryService:          categoryService,
	}
}

func (s *Service) getRecipes(offset int64, randomisation int64) ([]models.Recipe, error) {
	recipes, err := s.recipeRepo.getRecipes(offset, randomisation)
	if err != nil {
		return nil, err
	}

	/*for index, recipe := range recipes {
		allIngredients, err := s.ingredientsService.GetIngredientsByRecipe(recipe.Id)
		if err != nil {
			return nil, err
		}
		recipes[index].Ingredients = allIngredients
	}*/

	return recipes, err
}

func (s *Service) getByCategory(category string, offset int64, randomisation int64) ([]models.Recipe, error) {
	recipes, err := s.recipeRepo.getByCategory(category, offset, randomisation)
	if err != nil {
		return nil, err
	}

	/*for index, recipe := range recipes {
		allIngredients, err := s.ingredientsService.GetIngredientsByRecipe(recipe.Id)
		if err != nil {
			return nil, err
		}
		recipes[index].Ingredients = allIngredients
	}*/

	return recipes, err
}

func (s *Service) GetRandom() ([]models.Recipe, error) {
	recipes, err := s.recipeRepo.GetRandom()
	if err != nil {
		return nil, err
	}

	for index, recipe := range recipes {
		allIngredients, err := s.ingredientsService.GetIngredientsByRecipe(recipe.Id)
		if err != nil {
			return nil, err
		}
		recipes[index].Ingredients = allIngredients
	}

	category, err := s.categoryService.GetById(recipes[0].Category)
	if err != nil {
		log.Warningf("RecipeService.UpdateRecipe: Could note Get Category: %v", err)
		return recipes, err
	}
	recipes[0].Category = category.Title

	return recipes, err
}

func (s *Service) getRecipeById(id string) (models.Recipe, error) {
	recipe, err := s.recipeRepo.getRecipeById(id)
	if err != nil {
		return recipe, err
	}

	allIngredients, err := s.ingredientsService.GetIngredientsByRecipe(recipe.Id)
	if err != nil {
		return recipe, err
	}
	recipe.Ingredients = allIngredients

	category, err := s.categoryService.GetById(recipe.Category)
	if err != nil {
		log.Warningf("RecipeService.UpdateRecipe: Could note Get Category: %v", err)
		return recipe, err
	}
	recipe.Category = category.Title

	return recipe, err
}

func (s *Service) PostRecipe(recipe *models.Recipe) error {
	id, err := uuid.NewV4()
	if err != nil {
		return err
	}
	recipe.Id = id.String()
	err = s.recipeRepo.postRecipe(recipe)
	if err != nil {
		return err
	}

	allIngredients := s.ingredientsService.PostIngredients(recipe)
	fmt.Print(allIngredients)
	//recipeIngredients, _ := s.recipeIngredientsService.PostRecipeIngredients()

	return err
}

func (s *Service) UpdateRecipe(id string, recipe *models.Recipe) error {
	recipe.Id = id
	oldRecipe, err := s.getRecipeById(recipe.Id)
	if err != nil {
		return err
	}
	for _, oldIngredient := range oldRecipe.Ingredients {
		ingredientFound := false
		for index, newIngredient := range recipe.Ingredients {
			if oldIngredient.IngredientId == newIngredient.IngredientId {
				ingredientFound = true
				if oldIngredient.Ingredient != newIngredient.Ingredient {
					_ = s.ingredientsService.UpdateIngredient(&recipe.Ingredients[index])
				}
			}
		}
		if ingredientFound == false {
			_ = s.ingredientsService.DeleteById(oldIngredient.IngredientId)
		}
	}
	if err := s.recipeRepo.updateRecipe(recipe); err != nil {
		log.Warningf("RecipeService.UpdateRecipe: Could note update Recipe: %v", err)
		return err
	}

	return nil
}

func (s *Service) deleteRecipe(id string) (models.Recipe, error) {
	var recipe models.Recipe
	recipe.Id = id
	err := s.ingredientsService.DeleteIngredients(id)
	if err != nil {
		return recipe, err
	}
	return s.recipeRepo.deleteRecipe(recipe)
}
