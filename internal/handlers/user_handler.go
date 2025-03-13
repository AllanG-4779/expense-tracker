package handlers

import (
	"log"

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


func LoginUser(context *gin.Context) {
	var user types.LoginRequest

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(400, gin.H{"error": err})
		return
	}
	token, err := service.LoginUser(user.Username, user.Password)
	if (err != nil){
		log.Println(err)
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	context.JSON(201, token)
	
}

func GetUser(context *gin.Context) {
	params := context.Params;
	log.Println(params)
	gmail := params.ByName("email")
	log.Println(gmail)
	user, err := service.GetUser(gmail)
	if (err != nil){
		context.JSON(400, gin.H{"error": err.Error()})
	}
	context.JSON(200, gin.H{"message": "User retrieved successfully", "user": user})

}
