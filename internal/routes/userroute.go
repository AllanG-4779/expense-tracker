package routes

import (
	"github.com/allang-4779/financer/internal/constants"
	"github.com/allang-4779/financer/internal/handlers"
	"github.com/gin-gonic/gin"
)

func UserRoutesSetup(engine *gin.Engine){
	matcher := engine.Group(constants.API_ENTRY_POINT_USERS)
	{
		matcher.GET(constants.API_REGISTER,  handlers.LoginUser)
	}


}