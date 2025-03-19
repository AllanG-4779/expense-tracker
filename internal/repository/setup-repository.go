package repository

import (
	"github.com/allang-4779/financer/internal/database"
	"github.com/allang-4779/financer/internal/models"
	"github.com/allang-4779/financer/internal/types"
	"log"
)

func InsertCategory(category *types.CategoryRequest) {
	var categoryCreate = models.Category{
		Name:        category.Name,
		Icon:        category.Icon,
		Type:        category.Type,
		Description: category.Description,
	}
	err := database.DB.Create(&categoryCreate)
	if err != nil {
		log.Panic(err)
	}
}

func GetCategories(request types.CategoryRequest) ([]models.Category, error) {
	var categories []models.Category
	database.DB.Limit(request.Size).Offset(request.Page * request.Size).Find(&categories)
	return categories, nil
}
func GetCategoryByName(name string) (*models.Category, error) {
	var category models.Category
	err := database.DB.Where(&models.Category{Name: name}).First(&category)
	if err != nil {
		log.Panic(err)
	}
	return &category, nil
}
func UpdateCategory(categoryData models.Category) error {
	err := database.DB.Updates(categoryData)
	return err.Error

}
func DeleteCategory(category models.Category) error {
	err := database.DB.Delete(&category)
	return err.Error
}
