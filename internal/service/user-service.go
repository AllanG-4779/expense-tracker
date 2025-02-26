package service

import (
	"errors"
	

	"github.com/allang-4779/financer/internal/constants"
	"github.com/allang-4779/financer/internal/models"
	"github.com/allang-4779/financer/internal/repository"
	"github.com/allang-4779/financer/internal/types"
)

func RegisterUser(user *types.UserRegistration) error {
	var newUser models.SystemUser
	
	newUser, err := repository.GetUserByEmail(user.Email)
	if newUser.Email == user.Email {
		return errors.New(constants.USER_ALREADY_EXISTS)
	}
	if err != nil {
		return  repository.CreateUser(user)
	}
	return errors.New(constants.INTERNAL_SERVER_ERROR)
	
	}






	 


func LoginUser(email string, password string) error {
	return repository.CreateLoginAccount(email, password)
}
