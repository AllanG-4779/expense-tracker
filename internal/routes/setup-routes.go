package routes

import (
	"github.com/allang-4779/financer/internal/constants"
	"github.com/allang-4779/financer/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(engine *gin.Engine) {
	setup := engine.Group(constants.SetupEntryPoint)
	setup.POST(constants.SetupCategoryEndpoint, handlers.CreateCategory)
	setup.POST(constants.SetupCategoryEndpointGET, handlers.GetCategories)
	setup.PUT(constants.SetupCategoryEndpoint, handlers.UpdateCategory)
	setup.DELETE(constants.SetupCategoryEndpoint, handlers.DeleteCategory)
}
