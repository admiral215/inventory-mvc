package repositories

import (
	"fmt"
	"gorm.io/gorm"
	"inventory-bee/conf"
	"inventory-bee/models"
)

type CategoryRepository interface {
	Create(category *models.Category) error
	Update(category *models.Category) error
	DeleteById(id uint) error
	FindById(id uint) (*models.Category, error)
	FindAllBySearch(search string) ([]models.Category, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository() CategoryRepository {
	return &categoryRepository{
		db: conf.GetDB(),
	}
}

func (c *categoryRepository) Create(category *models.Category) error {
	err := c.db.Create(category).Error
	return err
}

func (c *categoryRepository) Update(category *models.Category) error {
	return c.db.Save(category).Error
}

func (c *categoryRepository) DeleteById(id uint) error {
	result := c.db.Delete(&models.Category{}, id)
	if result.RowsAffected == 0 {
		return fmt.Errorf("category with id %d not found", id)
	}
	return result.Error
}

func (c *categoryRepository) FindById(id uint) (*models.Category, error) {
	var category models.Category
	err := c.db.First(&category, id).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (c *categoryRepository) FindAllBySearch(search string) ([]models.Category, error) {
	var categories []models.Category

	query := c.db
	if search != "" {
		query = query.Where("name ILIKE ?", "%"+search+"%")
	}

	err := query.Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}
