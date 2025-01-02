package services

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"inventory-bee/dto"
	"inventory-bee/models"
	"inventory-bee/repositories"
)

type CategoryService interface {
	Create(dto *dto.CategoryCreate) error
	GetById(categoryId uint) (*models.Category, error)
	GetAllBySearch(search string) ([]models.Category, error)
	Edit(dto *dto.CategoryUpdate) error
	DeleteById(categoryId uint) error
}

type categoryService struct {
	categoryRepository repositories.CategoryRepository
}

func NewCategoryService(categoryRepository repositories.CategoryRepository) CategoryService {
	return &categoryService{
		categoryRepository: categoryRepository,
	}
}

func (c *categoryService) Create(dto *dto.CategoryCreate) error {
	newCategory := models.Category{
		Name: dto.Name,
	}
	return c.categoryRepository.Create(&newCategory)
}

func (c *categoryService) Edit(dto *dto.CategoryUpdate) error {
	category, err := c.categoryRepository.FindById(dto.Id)
	if err != nil {
		return err
	}

	category.Name = dto.Name

	err = c.categoryRepository.Update(category)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("category not found")
		}
	}
	return nil
}

func (c *categoryService) GetById(categoryId uint) (*models.Category, error) {
	return c.categoryRepository.FindById(categoryId)
}

func (c *categoryService) GetAllBySearch(search string) ([]models.Category, error) {
	categories, err := c.categoryRepository.FindAllBySearch(search)
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (c *categoryService) DeleteById(categoryId uint) error {
	return c.categoryRepository.DeleteById(categoryId)
}
