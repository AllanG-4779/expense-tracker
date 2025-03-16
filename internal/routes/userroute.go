package routes

import (
	"github.com/allang-4779/financer/internal/constants"
	"github.com/allang-4779/financer/internal/handlers"
	"github.com/allang-4779/financer/internal/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutesSetup(engine *gin.Engine) {
	matcher := engine.Group(constants.API_ENTRY_POINT_USERS)
	authenticated := engine.Group(constants.API_ENTRY_POINT_USERS)
	authenticated.Use(middleware.AuthMiddleware())
	matcher.POST(constants.API_REGISTER, handlers.RegisterUser)
	matcher.POST(constants.API_LOGIN, handlers.LoginUser)
	authenticated.GET(constants.API_GET_USER, handlers.GetUser)
	authenticated.PUT(constants.API_GET_USER, handlers.UpdateUser)

}
