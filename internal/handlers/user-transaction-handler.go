package handlers

import (
	"log"

	"github.com/allang-4779/financer/internal/constants"
	"github.com/allang-4779/financer/internal/service"
	"github.com/allang-4779/financer/internal/types"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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
		log.Print(err)
		if err != nil {
			context.JSON(400, gin.H{"message": "Error adding transaction: " + err.Error()})
			return
		}
		context.JSON(200, gin.H{"message": "Transaction added"})
		return
	} else {
		context.JSON(400, gin.H{"message": "Error adding transaction"})
		return
	}
}

func CreateBudget(context *gin.Context) {
	authCtx := context.MustGet(constants.CLAIMS).(jwt.MapClaims)
	username := authCtx["username"].(string)
	var budget types.BudgetRequest
	if context.ShouldBindJSON(&budget) == nil {
		err := service.CreateBudget(budget, username)
		if err != nil {
			context.JSON(400, gin.H{"message": "Error creating budget"})
			return
		}
		context.JSON(200, gin.H{"message": "Budget created"})
		return
	} else {
		context.JSON(400, gin.H{"message": "Error creating budget"})
		return
	}
}

func GetTransactions(context *gin.Context) {
	authCtx := context.MustGet(constants.CLAIMS).(jwt.MapClaims)
	username := authCtx["username"].(string)
	var account types.FetchRequest
	if context.ShouldBindJSON(&account) == nil {
		transactions, err := service.GetTransactions(account, username)
		if err != nil {
			context.JSON(400, gin.H{"message": "Error fetching transactions"})
			return
		}
		if len(transactions) == 0 {
			context.JSON(200, gin.H{"message": "No transactions found"})
			return
		}
		context.JSON(200, gin.H{"message": "Transactions fetched", "transactions": transactions})
		return
	} else {
		context.JSON(400, gin.H{"message": "Error fetching transactions"})
		return
	}

}

func GetUserAccounts(context *gin.Context) {
	authCtx := context.MustGet(constants.CLAIMS).(jwt.MapClaims)
	username := authCtx["username"].(string)

	accounts, err := service.GetUserAccounts(username)
	if err != nil {
		log.Println(err)
		context.JSON(400, gin.H{"message": "Error fetching accounts"})
		return
	}
	if len(accounts) == 0 {
		context.JSON(200, gin.H{"message": "No accounts found"})
		return
	}
	context.JSON(200, gin.H{"message": "Accounts fetched", "accounts": accounts})

}

func UpdateTransaction(context *gin.Context) {
	authCtx := context.MustGet(constants.CLAIMS).(jwt.MapClaims)
	username := authCtx["username"].(string)
	var transaction types.TransactionRequest
	if context.ShouldBindJSON(&transaction) == nil {
		err := service.UpdateTransaction(transaction, username)
		if err != nil {
			log.Println(err)
			context.JSON(400, gin.H{"message": "Error updating transaction:" + err.Error()})
			return
		}
		context.JSON(200, gin.H{"message": "Transaction updated"})
		return
	} else {
		log.Println("Error binding JSON:", context.Errors)
		context.JSON(400, gin.H{"message": "unable to bind json payload"})
		return
	}
}

func FilterTransactions(context *gin.Context) {
	authCtx := context.MustGet(constants.CLAIMS).(jwt.MapClaims)
	username := authCtx["username"].(string)
	var filter types.FilterRequest
	if context.ShouldBindJSON(&filter) == nil {
		transactions, err := service.FilterTransactions(filter, username)
		if err != nil {
			log.Println(err.Error())
			context.JSON(400, gin.H{"message": "Error filtering transactions"})
			return
		}
		if len(transactions) == 0 {
			context.JSON(200, gin.H{"message": "No transactions found"})
			return
		}
		context.JSON(200, gin.H{"message": "Transactions filtered", "transactions": transactions})
		return
	} else {
		context.JSON(400, gin.H{"message": "Error filtering transactions"})
		return
	}
}

func DeleteTransaction(context *gin.Context) {
	authCtx := context.MustGet(constants.CLAIMS).(jwt.MapClaims)
	username := authCtx["username"].(string)
	var request types.FetchRequest
	if context.ShouldBindJSON(&request) == nil {
		err := service.DeleteTransaction(uint(request.ID), username)
		if err != nil {
			log.Println(err)
			context.JSON(400, gin.H{"message": "Error deleting transaction"})
			return
		}
		context.JSON(200, gin.H{"message": "Transaction deleted"})
		return
	} else {
		context.JSON(400, gin.H{"message": "Error deleting transaction"})
		return
	}

}

func GetDashboardData(context *gin.Context) {
	authCtx := context.MustGet(constants.CLAIMS).(jwt.MapClaims)
	username := authCtx["username"].(string)
	var request types.FilterRequest
	if context.ShouldBindJSON(&request) == nil {
		data, err := service.GetDashboardData(request, username)
		if err != nil {
			log.Println(err)
			context.JSON(400, gin.H{"message": err.Error()})
			return
		}
		context.JSON(200, gin.H{"message": "Dashboard data fetched", "data": data})
		return
	} else {
		context.JSON(400, gin.H{"message": "Unable to bind payload"})
		return
	}
}

func GetBudgets(context *gin.Context) {
	authCtx := context.MustGet(constants.CLAIMS).(jwt.MapClaims)
	username := authCtx["username"].(string)
	request := types.FetchRequest{}
	if context.ShouldBindJSON(&request) == nil {
		returned, err := service.GetBudgets(request, username)
		response := types.UniversalResponse{}
		response.Body = returned
		response.Status = 200
		response.Message = "Budgets fetched successfully"
		if err != nil {
			log.Println(err)
			context.JSON(400, gin.H{"message": "Error fetching budgets"})
			return
		}
		context.JSON(200, gin.H{"message": "Budgets fetched", "budgets": response})
	}
}
