package service

import (
	"errors"
	"github.com/allang-4779/financer/internal/models"
	"log"

	"github.com/allang-4779/financer/internal/constants"
	"github.com/allang-4779/financer/internal/repository"
	"github.com/allang-4779/financer/internal/security"
	"github.com/allang-4779/financer/internal/types"
)

func RegisterUser(user *types.UserRegistration) error {
	var emailExists bool
	var userName bool

	_, err := repository.GetUser(user.Email, constants.EMAIL)
	emailExists = err == nil

	_, err = repository.GetUser(user.Username, constants.USERNAME)
	userName = err == nil

	if emailExists {
		return errors.New(constants.EMAIL_ALREADY_EXISTS)
	}
	if userName {
		return errors.New(constants.USERNAME_ALREADY_EXISTS)
	}

	pass, err := security.EncryptPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = pass
	return repository.CreateUser(user)

}

func LoginUser(username string, password string) (types.LoginResponse, error) {
	var loginResponse types.LoginResponse
	err, user := repository.GetLoginAccount(username)
	log.Println(user)
	if err != nil {
		return loginResponse, errors.New(constants.INVALID_CREDENTIALS)
	}
	if !security.VerifyPassword(password, user.Password) {
		log.Print("Password is incorrect")
		return loginResponse, errors.New(constants.INVALID_CREDENTIALS)
	}
	var payload = make(map[string]interface{})
	payload["username"] = user.Username
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

func GetUser(email string, param string) (types.UserRegistration, error) {
	user, err := repository.GetUser(email, param)
	if err != nil {
		return types.UserRegistration{}, errors.New("could not retrieve user")
	}
	return types.UserRegistration{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Username:  user.Username,
	}, nil
}

func UpdateProfile(user *types.UserRegistration, username string) (models.SystemUser, error) {
	var systemUser *models.SystemUser
	systemUser, err := repository.GetUser(username, constants.USERNAME)
	if err != nil {
		return models.SystemUser{}, err
	}
	if user.FirstName != "" {
		systemUser.FirstName = user.FirstName
	}
	if user.Email != "" {
		// Check if the Email exists
		existingUser, err := repository.GetUser(user.Email, constants.EMAIL)

		if err == nil && existingUser.Username != username {
			return models.SystemUser{}, errors.New(constants.EMAIL_ALREADY_EXISTS)
		}
		systemUser.Email = user.Email
	}
	if user.LastName != "" {
		systemUser.LastName = user.LastName
	}
	if user.NewPassword != "" {
		if ok := security.VerifyPassword(systemUser.Password, user.Password); ok {
			systemUser.Password, err = security.EncryptPassword(user.NewPassword)
			if err != nil {
				return models.SystemUser{}, err
			}
		} else {
			return models.SystemUser{}, errors.New(constants.INVALID_CREDENTIALS)
		}
	}
	return repository.UpdateProfile(systemUser)
}
