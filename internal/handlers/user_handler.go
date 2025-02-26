package handlers

import (
	"github.com/allang-4779/financer/internal/service"
	"github.com/allang-4779/financer/internal/types"
	"github.com/gin-gonic/gin"
)

func RegisterUser(context *gin.Context) {
	var user types.UserRegistration

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}else if err := service.RegisterUser(&user); err != nil {
		context.JSON(400, gin.H{"error": "Could not register user"})
		return;
	}
	context.JSON(201, gin.H{"message": "User registered successfully"})
	
}
