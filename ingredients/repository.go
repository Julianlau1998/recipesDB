package ingredients

import (
	"database/sql"
	log "github.com/sirupsen/logrus"
	"recipes/models"
	"recipes/utility"
)

type Repository struct {
	dbClient *sql.DB
}

func NewRepository(dbClient *sql.DB) Repository {
	return Repository{dbClient: dbClient}
}

func (r *Repository) GetIngredientsByRecipe(recipeID string) ([]models.Ingredient, error) {
	var ingredients []models.Ingredient
	query := `SELECT ingredients.uuid, ingredients.ingredient, ingredients.measurement FROM recipe_ingredients JOIN ingredients ON ingredients.uuid = recipe_ingredients.ingredient_ID WHERE recipe_ingredients.recipe_ID = $1`
	ingredients, err := r.fetch(query, recipeID)
	if err != nil {
		return ingredients, err
	}
	return ingredients, nil
}

func (r *Repository) PostIngredients(ingredients []models.Ingredient, recipeIngredients []models.RecipeIngredients) error {
	return utility.Transact(r.dbClient, func(tx *sql.Tx) error {
		for _, ingredient := range ingredients {
			ingredientsQuery := `INSERT INTO ingredients (uuid, ingredient, measurement) VALUES ($1, $2, $3)`
			_, err := r.dbClient.Exec(ingredientsQuery, ingredient.IngredientId, ingredient.Ingredient, ingredient.Measurement)
			if err != nil {
				return err
			}
		}
		for _, recipeIngredient := range recipeIngredients {
			recipeIngredientsQuery := `INSERT INTO recipe_ingredients (recipe_ID, ingredient_ID) VALUES ($1, $2)`
			_, err := r.dbClient.Exec(recipeIngredientsQuery, recipeIngredient.Recipe_ID, recipeIngredient.Ingredients_ID)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *Repository) UpdateIngredient(ingredient *models.Ingredient) error {
	query := "UPDATE ingredients SET ingredient=$1 WHERE uuid = $2"
	_, err := r.dbClient.Exec(query, ingredient.Ingredient, ingredient.IngredientId)
	return err
}

func (r *Repository) DeleteIngredients(recipeID string) error {
	query := `DELETE ingredients.* FROM recipe_ingredients JOIN ingredients ON ingredients.uuid = recipe_ingredients.ingredient_ID WHERE recipe_ingredients.recipe_ID = $1`
	_, err := r.dbClient.Exec(query, recipeID)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteIngredientById(ingredientID string) error {
	query := `DELETE FROM ingredients WHERE uuid = $1`
	_, err := r.dbClient.Exec(query, ingredientID)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) fetch(query string, recipeID string) ([]models.Ingredient, error) {
	result := make([]models.Ingredient, 0)
	rows, err := r.dbClient.Query(query, recipeID)
	if err != nil {
		return result, err
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Errorf("Datenbankverbindung konnte nicht korrekt geschlossen werden: %v", err)
		}
	}()
	for rows.Next() {
		ingredientsDB := models.IngredientsDB{}
		err := rows.Scan(&ingredientsDB.IngredientId, &ingredientsDB.Ingredient, &ingredientsDB.Measurement)
		if err != nil {
			if err == sql.ErrNoRows {
				continue
			}
			log.Infof("Fehler beim Lesen der Daten: %v", err)
			return result, err
		}
		result = append(result, ingredientsDB.GetIngredients())
	}
	return result, nil
}
