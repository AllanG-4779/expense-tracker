package service

import (
	"errors"
	"log"

	"github.com/allang-4779/financer/internal/models"

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

	_, errUsername := repository.GetUser(user.Username, constants.USERNAME)
	userName = errUsername == nil

	if emailExists {
		return errors.New(constants.EmailAlreadyExists)
	}
	if userName {
		return errors.New(constants.UsernameAlreadyExists)
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
		return loginResponse, errors.New(constants.InvalidCredentialsMessage)
	}
	if !security.VerifyPassword(password, user.Password) {
		log.Print("Password is incorrect")
		return loginResponse, errors.New(constants.InvalidCredentialsMessage)
	}
	var payload = make(map[string]interface{})
	payload["username"] = user.Username
	response, err := security.GenerateAccessToken(payload, 3600)
	if err != nil {
		loginResponse.Status = 500
		loginResponse.Message = constants.InternalServerErrorMessage
		return loginResponse, err
	}

	loginResponse.Token = types.TokenResponse{Token: response, ValidFor: 3600}
	loginResponse.Message = constants.LoginSuccessfulMessage
	loginResponse.Status = 200
	usr := types.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Username:  user.Username,
	}
	loginResponse.User = usr
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
			return models.SystemUser{}, errors.New(constants.EmailAlreadyExists)
		}
		systemUser.Email = user.Email
	}
	if user.LastName != "" {
		systemUser.LastName = user.LastName
	}
	if user.NewPassword != "" {
		if ok := security.VerifyPassword(user.Password, systemUser.Password); ok {
			log.Println("password verification successful")
			pass, verificationError := security.EncryptPassword(user.NewPassword)
			if verificationError != nil {
				return models.SystemUser{}, verificationError
			}
			systemUser.Password = pass
		} else {
			return models.SystemUser{}, errors.New(constants.PasswordVerificationErrorMessage)
		}
	}
	return repository.UpdateProfile(systemUser)
}
