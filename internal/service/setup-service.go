package service

import (
	"errors"
	"github.com/allang-4779/financer/internal/repository"
	"github.com/allang-4779/financer/internal/types"
)

func CreateCategory(category *types.CategoryRequest) {
	repository.InsertCategory(category)
}

func GetCategories(category types.CategoryRequest) ([]types.CategoryRequest, error) {
	return repository.GetCategories(category)
}

func UpdateCategory(category *types.CategoryRequest) (*types.CategoryRequest, error) {
	data, err := repository.GetCategoryByName(category.Name)
	if err != nil {
		return nil, errors.New("could not get category to update")
	}
	if category.Description != "" {
		data.Description = category.Description
	}
	if category.Name != "" {
		data.Name = category.Name
	}
	if category.Icon != "" {
		data.Icon = category.Icon
	}
	updateErr := repository.UpdateCategory(*data)
	if updateErr != nil {
		return nil, errors.New("could not update category")
	}
	return repository.GetCategoryByName(data.Name)
}
func DeleteCategory(category types.CategoryRequest) error {
	return repository.DeleteCategory(category)
}
