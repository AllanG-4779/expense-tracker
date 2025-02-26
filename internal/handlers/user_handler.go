package handlers

import (
	"github.com/allang-4779/financer/internal/types"
	"github.com/gin-gonic/gin"
)

func LoginUser(context *gin.Context){

   data := types.UserLoginResponse{
	Message: "User login successful",
	Status: 200,
	Successful: true,
	Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VyX25hbWUiOiJhbGxhIiwidXNlcl9lbWFpbCI6ImFsYUBnbWFpbC5jb20iLCJleHBpcmVkX3N0YW5kYXJkX2Zyb20iOiIyMDIxLTAzLTAxVDEwOjMwOjQwLjAwMDAwMCIsImlhdCI6MTYxNDYwNjQwMH0.1",
   }

   context.JSON(200, data)

}