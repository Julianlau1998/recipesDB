package ingredients

import (
	uuid "github.com/nu7hatch/gouuid"
	log "github.com/sirupsen/logrus"
	"recipes/models"
	"recipes/recipeIngredients"
)

/*import (
	"github.com/nu7hatch/gouuid"
	"recipes/models"
)*/

type Service struct {
	ingredientsRepo          Repository
	recipeIngredientsService recipeIngredients.Service
}

func NewService(ingredientsRepositoty Repository, recipeIngredients recipeIngredients.Service) Service {
	return Service{
		ingredientsRepo:          ingredientsRepositoty,
		recipeIngredientsService: recipeIngredients,
	}
}

func (s *Service) GetIngredientsByRecipe(recipeID string) ([]models.Ingredient, error) {
	allIngredients, err := s.ingredientsRepo.GetIngredientsByRecipe(recipeID)
	if err != nil {
		log.Warningf("IngredientService.GetIngredientByRecipe: Unable to Load ingredient by recipe: %v", err)
		return allIngredients, err
	}
	/* for _, ingredient := range allIngredients {
	 if ingredient.IngredientId == recipeIngredientsId {
		 recipeIngredients = append(recipeIngredients, ingredient)
	 }
	} */
	return allIngredients, err
}

func (s *Service) PostIngredients(recipe *models.Recipe) error {
	ingredients := recipe.Ingredients
	var ingredientsModelArray []models.Ingredient
	var ingredientsModel models.Ingredient
	var recipeIngredients models.RecipeIngredients
	var recipeIngredientsArray []models.RecipeIngredients

	for _, ingredient := range ingredients {
		id, err := uuid.NewV4()
		if err != nil {
			return err
		}
		ingredientsModel.Ingredient = ingredient.Ingredient
		ingredientsModel.IngredientId = id.String()
		ingredientsModelArray = append(ingredientsModelArray, ingredientsModel)

		recipeIngredients.Ingredients_ID = id.String()
		recipeIngredients.Recipe_ID = recipe.Id
		recipeIngredientsArray = append(recipeIngredientsArray, recipeIngredients)
		/*s.recipeIngredientsService.PostRecipeIngredients(recipe.Id, ingredientsModel.IngredientId)*/
	}
	return s.ingredientsRepo.PostIngredients(ingredientsModelArray, recipeIngredientsArray)
}

func (s *Service) UpdateIngredient(ingredient *models.Ingredient) error {
	err := s.ingredientsRepo.UpdateIngredient(ingredient)
	if err != nil {
		log.Warningf("IngredientsService.UpdateIngredient Ingredient could not get updated: %v", err)
		return err
	}
	return nil
}

func (s *Service) DeleteIngredients(recipeID string) error {
	return s.ingredientsRepo.DeleteIngredients(recipeID)
}

func (s *Service) DeleteById(ingredientID string) error {
	err := s.ingredientsRepo.DeleteIngredientById(ingredientID)
	if err != nil {
		log.Warningf("IngredientsService.DeleteById Ingredients could not get deleted by ID: %v", err)
		return err
	}
	return nil
}
