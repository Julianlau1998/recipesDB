package categories

import (
	uuid "github.com/nu7hatch/gouuid"
	"recipes/models"
)

type Service struct {
	categoryRepo Repository
}

func NewService(categoryRepository Repository) Service {
	return Service{categoryRepo: categoryRepository}
}

func (s *Service) GetCategories() ([]models.Category, error) {
	return s.categoryRepo.GetCategories()
}

func (s *Service) GetById(id string) (models.Category, error) {
	return s.categoryRepo.GetCategoryById(id)
}

func (s *Service) postCategory(category *models.Category) (*models.Category, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return category, err
	}
	category.Id = id.String()
	return s.categoryRepo.PostCategory(category)
}

func (s *Service) updateCategory(id string, category *models.Category) (models.Category, error) {
	category.Id = id
	return s.categoryRepo.updateCategory(category)
}

func (s *Service) deleteCategory(id string) (models.Category, error) {
	var category models.Category
	category.Id = id
	return s.categoryRepo.deleteCategory(category)
}
