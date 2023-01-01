package categories

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

func (r *Repository) GetCategories() ([]models.Category, error) {
	var categories []models.Category
	query := `SELECT uuid, title, color, createdDate FROM categories`
	categories, err := r.fetch(query)
	return categories, err
}

func (r *Repository) GetCategoryById(id string) (models.Category, error) {
	var category models.Category

	query := `SELECT uuid, title, color, createdDate FROM categories WHERE uuid = $1`
	category, err := r.getOne(query, id)
	return category, err
}

func (r *Repository) PostCategory(category *models.Category) (*models.Category, error) {
	statement := `INSERT INTO categories (uuid, title, color, createdDate) VALUES ($1, $2, $3, CURRENT_TIMESTAMP)`
	_, err := r.dbClient.Exec(statement, category.Id, category.Title, category.Color)
	return category, err
}

func (r *Repository) updateCategory(category *models.Category) (models.Category, error) {
	query := `UPDATE categories SET title = $1, color = $2 WHERE uuid = $3`
	_, err := r.dbClient.Exec(query, category.Title, category.Color, category.Id)

	return *category, err
}

func (r *Repository) deleteCategory(category models.Category) (models.Category, error) {
	query := `DELETE FROM categories WHERE uuid = $1`
	_, err := r.dbClient.Exec(query, category.Id)
	return category, err
}

func (r *Repository) fetch(query string) ([]models.Category, error) {
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
	result := make([]models.Category, 0)
	for rows.Next() {
		categoryDB := models.CategoryDB{}
		err := rows.Scan(&categoryDB.Id, &categoryDB.Title, &categoryDB.Color, &categoryDB.CreatedDate)
		if err != nil {
			if err == sql.ErrNoRows {
				continue
			}
			log.Infof("Fehler beim Lesen der Daten: %v", err)
			return result, err
		}
		result = append(result, categoryDB.GetCategory())
	}
	return result, nil
}

func (r *Repository) getOne(query string, id string) (models.Category, error) {
	categoryDB := models.CategoryDB{}
	err := r.dbClient.QueryRow(query, id).Scan(&categoryDB.Id, &categoryDB.Title, &categoryDB.Color, &categoryDB.CreatedDate)
	if err != nil && err != sql.ErrNoRows {
		log.Infof("Fehler beim Lesen der Daten: %v", err)
	}
	return categoryDB.GetCategory(), err
}
