package routes

import (
	"github.com/allang-4779/financer/internal/constants"
	"github.com/allang-4779/financer/internal/handlers"
	"github.com/allang-4779/financer/internal/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutesSetup(engine *gin.Engine) {
	matcher := engine.Group(constants.ApiEntryPointUsers)
	authenticated := engine.Group(constants.ApiEntryPointUsers)
	authenticated.Use(middleware.AuthMiddleware())
	matcher.POST(constants.ApiRegister, handlers.RegisterUser)
	matcher.POST(constants.ApiLogin, handlers.LoginUser)
	authenticated.GET(constants.ApiGetUser, handlers.GetUser)
	authenticated.PUT(constants.ApiGetUser, handlers.UpdateUser)
	authenticated.POST(constants.ActivateAccount, handlers.ActivateAccount)
	authenticated.POST(constants.AddTransaction, handlers.AddTransaction)
	authenticated.POST(constants.CreateBudget, handlers.CreateBudget)
	authenticated.GET(constants.GetTransactions, handlers.GetTransactions)
	authenticated.GET(constants.GetAccounts, handlers.GetUserAccounts)
}
