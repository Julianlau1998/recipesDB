package recipeIngredients

import (
	"database/sql"
	log "github.com/sirupsen/logrus"
	"recipes/models"
)

type Repository struct {
	dbClient *sql.DB
}

func NewRepository(dbClient *sql.DB) Repository {
	return Repository{dbClient: dbClient}
}

func (r *Repository) GetRecipeIngredients() ([]models.RecipeIngredients, error) {
	var recipeIngredients []models.RecipeIngredients
	query := `SELECT recipe_ID, ingredient_ID FROM recipe_ingredients`
	recipeIngredients, err := r.fetch(query)
	return recipeIngredients, err
}

func (r *Repository) PostRecipeIngredients(recipeId string, ingredientId string) error {
	recipeIngredientsQuery := `INSERT INTO recipe_ingredients (recipe_ID, ingredient_ID) VALUES ($1, $2)`
	_, err := r.dbClient.Exec(recipeIngredientsQuery, recipeId, ingredientId)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) fetch(query string) ([]models.RecipeIngredients, error) {
	rows, err := r.dbClient.Query(query)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Errorf("Datenbankverbindung konnte nicht korrekt geschlossen werden: %v", err)
		}
	}()
	result := make([]models.RecipeIngredients, 0)
	for rows.Next() {
		RecipeIngredientsDB := models.RecipeIngredientsDB{}
		err := rows.Scan(&RecipeIngredientsDB.Recipe_ID, &RecipeIngredientsDB.Ingredients_ID)
		if err != nil {
			if err == sql.ErrNoRows {
				continue
			}
			log.Infof("Fehler beim Lesen der Daten: %v", err)
			return result, err
		}
		result = append(result, RecipeIngredientsDB.GetRecipeIngredients())
	}
	return result, nil
}
