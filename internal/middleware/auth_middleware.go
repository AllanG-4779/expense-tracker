package middleware

import (
	"log"
	"strings"

	"github.com/allang-4779/financer/internal/security"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc{
	return func(context *gin.Context){
		// Do some authentication here
		authorization := context.GetHeader("Authorization")
		if (len(authorization)==0 || strings.Split(authorization, " ")[0] != "Bearer"){
			context.JSON(401, gin.H{"error": "missing required authorization header"})
			context.Abort()
			return
		}

		claims, err := security.VerifyToken(strings.Split(authorization, " ")[1])
		if err != nil {
			log.Print(err)
			context.JSON(401, gin.H{"error": "Invalid authorization token"})
			context.Abort()
			return
		}
		context.Set("claims", claims)
		context.Next()


	}
}