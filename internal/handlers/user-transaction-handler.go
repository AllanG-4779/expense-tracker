package handlers

import (
	"github.com/allang-4779/financer/internal/constants"
	"github.com/allang-4779/financer/internal/service"
	"github.com/allang-4779/financer/internal/types"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"log"
)

func ActivateAccount(context *gin.Context) {
	authCtx := context.MustGet(constants.CLAIMS).(jwt.MapClaims)
	username := authCtx["username"].(string)
	var account types.AccountRequest
	if context.ShouldBindJSON(&account) == nil {
		err := service.ActivateAccount(account, username)
		if err != nil {
			log.Println(err)
			context.JSON(400, gin.H{"message": "Error activating account"})
			return
		}
		context.JSON(200, gin.H{"message": "Account activated"})
		return
	} else {
		context.JSON(400, gin.H{"message": "Error activating account"})
		return
	}
}

func AddTransaction(context *gin.Context) {
	authCtx := context.MustGet(constants.CLAIMS).(jwt.MapClaims)
	username := authCtx["username"].(string)
	var transaction types.TransactionRequest
	if context.ShouldBindJSON(&transaction) == nil {
		err := service.AddTransaction(transaction, username)
		if err != nil {
			context.JSON(400, gin.H{"message": "Error adding transaction"})
			return
		}
		context.JSON(200, gin.H{"message": "Transaction added"})
		return
	} else {
		context.JSON(400, gin.H{"message": "Error adding transaction"})
		return
	}
}
