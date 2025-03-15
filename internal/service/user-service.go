package service

import (
	"errors"
	"log"

	"github.com/allang-4779/financer/internal/constants"
	"github.com/allang-4779/financer/internal/models"
	"github.com/allang-4779/financer/internal/repository"
	"github.com/allang-4779/financer/internal/security"
	"github.com/allang-4779/financer/internal/types"
)

func RegisterUser(user *types.UserRegistration) error {
	var newUser models.SystemUser

	newUser, err := repository.GetUserByEmail(user.Email)
	if newUser.Email == user.Email {
		return errors.New(constants.USER_ALREADY_EXISTS)
	}
	if err != nil {
		pass, err := security.EncryptPassword(user.Password)
		if err != nil {
			return err
		}
		user.Password = pass

		return repository.CreateUser(user)
	}
	return errors.New(constants.INTERNAL_SERVER_ERROR)

}

func LoginUser(email string, password string) (types.LoginResponse, error) {
	var loginResponse types.LoginResponse
	err, user := repository.GetLoginAccountByEmail(email)
	log.Println(user)
	if err != nil {
		return loginResponse, errors.New(constants.INVALID_CREDENTIALS)
	}
	if !security.VerifyPassword(password, user.Password) {
		log.Print("Password is incorrect")
		return loginResponse, errors.New(constants.INVALID_CREDENTIALS)
	}
	var payload = make(map[string]interface{})
	payload["email"] = user.Username
	response, err := security.GenerateAccessToken(payload, 3600)
	if err != nil {
		loginResponse.Status = 500
		loginResponse.Message = constants.INTERNAL_SERVER_ERROR
		return loginResponse, err
	}

	loginResponse.Token = types.TokenResponse{Token: response, ValidFor: 3600}
	loginResponse.Message = constants.LOGIN_SUCCESS
	loginResponse.Status = 200
	loginResponse.Successful = true
	return loginResponse, nil

}

func GetUser(email string) (types.UserRegistration, error) {
	user, err := repository.GetUserByEmail(email)
	if err != nil {
		return types.UserRegistration{}, errors.New("could not retrieve user")
	}
	return types.UserRegistration{
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		Dob:         user.Dob,
		Residential: user.Residential,
		Phone:       user.Phone,
	}, nil
}
