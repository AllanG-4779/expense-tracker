package handlers

import (
	"github.com/allang-4779/financer/internal/service"
	"github.com/allang-4779/financer/internal/types"
	"github.com/gin-gonic/gin"
	"log"
)

func CreateCategory(c *gin.Context) {

	var category types.CategoryRequest
	if c.ShouldBind(&category) == nil {
		service.CreateCategory(&category)
		c.JSON(200, gin.H{"message": "Category created"})
		return
	} else {
		c.JSON(400, gin.H{"message": "Error creating category"})
		return
	}

}

func GetCategories(c *gin.Context) {
	var category types.CategoryRequest

	if c.ShouldBindJSON(&category) == nil {
		categories, err := service.GetCategories(category)
		log.Printf("sending categories to client to be displayed %v", category)
		if err != nil {
			c.JSON(400, gin.H{"message": "Error fetching categories"})
			return
		}
		if len(categories) == 0 {
			c.JSON(200, gin.H{"message": "No categories found"})
			return
		}
		c.JSON(200, gin.H{"message": "Categories fetched", "categories": categories})
		return
	} else {
		c.JSON(400, gin.H{"message": "Error fetching categories"})
		return
	}

}

func UpdateCategory(c *gin.Context) {
	var category types.CategoryRequest
	if c.ShouldBindJSON(&category) == nil {
		category, err := service.UpdateCategory(&category)
		if err != nil {
			c.JSON(400, gin.H{"message": "Error updating category"})
			return
		}
		c.JSON(200, gin.H{"message": "Category updated", "category": category})
		return
	} else {
		c.JSON(400, gin.H{"message": "Error updating category"})
		return
	}
}

func DeleteCategory(c *gin.Context) {
	var category types.CategoryRequest
	if c.ShouldBindJSON(&category) == nil {
		err := service.DeleteCategory(category)
		if err != nil {
			c.JSON(400, gin.H{"message": "Error deleting category"})
			return
		}
		c.JSON(200, gin.H{"message": "Category deleted"})
		return
	} else {
		c.JSON(400, gin.H{"message": "Error deleting category"})
		return
	}
}
