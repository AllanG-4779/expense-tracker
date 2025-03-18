package repository

import (
	"github.com/allang-4779/financer/internal/database"
	"github.com/allang-4779/financer/internal/types"
	"log"
)

func InsertCategory(category *types.CategoryRequest) {
	_, err := database.DB.NamedExec(database.CreateCategory, category)
	if err != nil {
		log.Panic(err)
	}
}

func GetCategories(request types.CategoryRequest) ([]types.CategoryRequest, error) {
	var categories []types.CategoryRequest
	size := request.Size
	offset := request.Size * (request.Page)
	err := database.DB.Select(&categories, database.GetCategories, size, offset)
	if err != nil {
		log.Panic(err)
	}

	return categories, nil
}
func GetCategoryByName(name string) (*types.CategoryRequest, error) {
	var category types.CategoryRequest
	err := database.DB.Get(&category, database.GetCategory, name)
	if err != nil {
		log.Panic(err)
	}
	return &category, nil
}
func UpdateCategory(categoryData types.CategoryRequest) error {

	_, err := database.DB.NamedExec(database.UpdateCategory, categoryData)
	return err

}
func DeleteCategory(category types.CategoryRequest) error {
	_, err := database.DB.NamedExec(database.DeleteCategory, category)
	return err
}
