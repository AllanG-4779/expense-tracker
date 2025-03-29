package service

import (
	"errors"
	"github.com/allang-4779/financer/internal/models"
	"github.com/allang-4779/financer/internal/repository"
	"github.com/allang-4779/financer/internal/types"
)

func CreateCategory(category *types.CategoryRequest) error {
	if _, err := repository.GetCategoryByName(category.Name); err == nil {
		return errors.New("category already exists")
	}
	if category.Type != "expense" && category.Type != "income" {
		return errors.New("category type must be either expense or income")
	}
	err := repository.InsertCategory(category)
	if err != nil {
		return err
	}
	return nil
}

func GetCategories(category types.CategoryRequest) ([]models.Category, error) {
	return repository.GetCategories(category)
}

func UpdateCategory(category *types.CategoryRequest) (*models.Category, error) {
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
	data, err := repository.GetCategoryByName(category.Name)
	if err != nil {
		return errors.New("could not get category to delete")
	}
	return repository.DeleteCategory(*data)
}
