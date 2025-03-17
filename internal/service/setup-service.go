package service

import (
	"github.com/allang-4779/financer/internal/repository"
	"github.com/allang-4779/financer/internal/types"
)

func CreateCategory(category *types.CategoryRequest) {
	repository.InsertCategory(category)
}

func GetCategories(category types.CategoryRequest) ([]types.CategoryRequest, error) {
	return repository.GetCategories(category)
}
