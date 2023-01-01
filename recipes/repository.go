package recipes

import (
	"database/sql"
	"fmt"
	"recipes/models"
)

type Repository struct {
	dbClient *sql.DB
}

func NewRepository(dbClient *sql.DB) Repository {
	return Repository{dbClient: dbClient}
}

func (r *Repository) getRecipes(offset int64) ([]models.Recipe, error) {
	var recipes []models.Recipe
	query := `SELECT uuid, title, description, duration, category_ID, image, country, is_vegetarian, is_vegan, createdDate FROM recipes LIMIT 20 OFFSET $1`
	recipes, err := r.fetch(query, "", offset)
	return recipes, err
}

func (r *Repository) getByCategory(category string, offset int64) ([]models.Recipe, error) {
	var recipes []models.Recipe
	query := `SELECT uuid, title, description, duration, category_ID, image, country, is_vegetarian, is_vegan, createdDate FROM recipes WHERE category_ID = $1 LIMIT 20 OFFSET $2`
	recipes, err := r.fetch(query, category, offset)
	return recipes, err
}

func (r *Repository) GetRandom() ([]models.Recipe, error) {
	var recipes []models.Recipe
	query := `SELECT uuid, title, description, duration, category_ID, image, country, is_vegetarian, is_vegan, createdDate FROM recipes order by random() limit $1`
	recipes, err := r.fetch(query, "", 1)
	return recipes, err
}

func (r *Repository) getRecipeById(id string) (models.Recipe, error) {
	var recipe models.Recipe

	query := `SELECT uuid, title, description, duration, category_ID, image, country, is_vegetarian, is_vegan, createdDate FROM recipes WHERE uuid = $1`
	recipe, err := r.getOne(query, id)

	return recipe, err
}

func (r *Repository) postRecipe(recipe *models.Recipe) error {
	if recipe.Category == "" {
		recipe.Category = "1"
	}
	query := `INSERT INTO recipes (uuid, title, description, duration, category_id, image, country, is_vegetarian, is_vegan, createdDate) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, CURRENT_TIMESTAMP)`
	_, err := r.dbClient.Exec(query, recipe.Id, recipe.Title, recipe.Description, recipe.Duration, recipe.Category, recipe.Image, recipe.Country, recipe.IsVegetarian, recipe.IsVegan)

	return err
}

func (r *Repository) updateRecipe(recipe *models.Recipe) error {
	query := `UPDATE recipes SET title = $1, description = $2, duration = $3, category_id = $4, country = $5, is_vegetarian = $6, is_vegan = $7,
				WHERE uuid = ?`
	_, err := r.dbClient.Exec(query, recipe.Title, recipe.Description, recipe.Duration, recipe.Category, recipe.Country, recipe.IsVegetarian, recipe.IsVegan, recipe.Id)

	return err
}

func (r *Repository) deleteRecipe(recipe models.Recipe) (models.Recipe, error) {
	query := `DELETE FROM recipes where uuid = ?`
	_, err := r.dbClient.Exec(query, recipe.Id)
	if err != nil {
		return recipe, err
	}

	return recipe, nil
}

func (r *Repository) fetch(query string, category string, offset int64) ([]models.Recipe, error) {
	var rows *sql.Rows
	var err error
	if len(category) > 0 {
		rows, err = r.dbClient.Query(query, category, offset)
	} else {
		rows, err = r.dbClient.Query(query, offset)
	}
	if err != nil {
		return nil, err
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			fmt.Printf("Datenbankverbindung konnte nicht korrekt geschlossen werden: %v", err)
		}
	}()
	result := make([]models.Recipe, 0)
	for rows.Next() {
		recipeDB := models.RecipeDB{}
		err := rows.Scan(&recipeDB.Id, &recipeDB.Title, &recipeDB.Description, &recipeDB.Duration, &recipeDB.Category, &recipeDB.Image, &recipeDB.Country, &recipeDB.IsVegetarian, &recipeDB.IsVegan, &recipeDB.CreatedDate)
		if err != nil {
			if err == sql.ErrNoRows {
				continue
			}
			fmt.Printf("Fehler beim Lesen der Daten: %v", err)
			return result, err
		}
		result = append(result, recipeDB.GetRecipe())
	}
	return result, nil
}

/*func (r *Repository) getMultiple(query string, id string) ([]models.Recipe, error) {
	rows, err := r.dbClient.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			fmt.Printf("Datenbankverbindung konnte nicht korrekt geschlossen werden: %v", err)
		}
	}()
	result := make([]models.Recipe, 0)
	for rows.Next() {
		recipeDB := models.RecipeDB{}
		err := rows.Scan(&recipeDB.Id, &recipeDB.Title, &recipeDB.Description, &recipeDB.Duration, &recipeDB.Category)
		if err != nil {
			if err == sql.ErrNoRows {
				continue
			}
			fmt.Printf("Fehler beim Lesen der Daten: %v", err)
			return result, err
		}
		result = append(result, recipeDB.GetRecipe())
	}
	return result, nil
}*/

func (r *Repository) getOne(query string, id string) (models.Recipe, error) {
	recipeDB := models.RecipeDB{}
	err := r.dbClient.QueryRow(query, id).Scan(&recipeDB.Id, &recipeDB.Title, &recipeDB.Description, &recipeDB.Duration,
		&recipeDB.Category, &recipeDB.Image, &recipeDB.Country, &recipeDB.IsVegetarian, &recipeDB.IsVegan, &recipeDB.CreatedDate)
	if err != nil && err != sql.ErrNoRows {
		fmt.Printf("Fehler beim Lesen der Daten: %v", err)
	}
	return recipeDB.GetRecipe(), err
}
