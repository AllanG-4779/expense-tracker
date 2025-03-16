package repository

import (
	"github.com/allang-4779/financer/internal/constants"
	"github.com/allang-4779/financer/internal/database"
	"github.com/allang-4779/financer/internal/models"
	"github.com/allang-4779/financer/internal/types"
	"log"
)

/**
* This package contains the repository layer for the user entity
 */

func CreateUser(user *types.UserRegistration) error {

	userEntity := models.SystemUser{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Username:  user.Username,
		Password:  user.Password,
	}

	log.Print(userEntity)
	row, err := database.DB.NamedExec(database.CreateUser, userEntity)
	if err != nil {
		log.Panic(err)
	}
	log.Print(row)
	return nil
}

func GetUser(email string, param string) (*models.SystemUser, error) {
	var user models.SystemUser
	var err error

	switch param {
	case constants.EMAIL:
		err = database.DB.Get(&user, database.GetUserByEmail, email) // FIXED
	case constants.USERNAME:
		err = database.DB.Get(&user, database.LoginUsernameQuery, email) // FIXED
	}

	if err != nil {
		log.Printf("Error fetching user: %v", err)
		return nil, err
	}
	log.Printf("User: %v", user)
	return &user, nil
}

func UpdateProfile(user *models.SystemUser) (models.SystemUser, error) {
	_, err := database.DB.NamedExec(database.UpdateUserProfile, user)
	if err != nil {
		log.Printf("Error updating user: %v", err)
		return models.SystemUser{}, err
	}
	data, errorReturned := GetUser(user.Email, constants.EMAIL)

	return *data, errorReturned
}

func GetLoginAccount(email string) (error, models.SystemUser) {
	var loginAccount models.SystemUser

	err := database.DB.Get(&loginAccount, database.LoginUsernameQuery, email)
	if err != nil {
		log.Printf("Error fetching user: %v", err)
		return err, loginAccount
	}

	return nil, loginAccount

}
