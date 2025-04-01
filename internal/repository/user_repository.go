package repository

import (
	"github.com/allang-4779/financer/internal/constants"
	"github.com/allang-4779/financer/internal/database"
	"github.com/allang-4779/financer/internal/models"
	"github.com/allang-4779/financer/internal/types"
	"gorm.io/gorm"
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
	err := database.DB.Create(&userEntity)
	if err != nil {
		return err.Error
	}
	log.Print("User created successfully")
	return nil
}

func GetUser(email string, param string) (*models.SystemUser, error) {
	var user models.SystemUser
	var err *gorm.DB

	switch param {
	case constants.EMAIL:
		err = database.DB.First(&user, models.SystemUser{Email: email}) // FIXED
	case constants.USERNAME:
		err = database.DB.First(&user, models.SystemUser{Username: email}) // FIXED
	}

	if err != nil && err.Error != nil {
		log.Printf("Error fetching user: %v", err)
		return nil, err.Error
	}
	log.Printf("User: %v", user)
	return &user, nil
}

func UpdateProfile(user *models.SystemUser) (models.SystemUser, error) {
	err := database.DB.Create(&user)
	if err != nil {
		log.Printf("Error updating user: %v", err)
		return models.SystemUser{}, err.Error
	}
	data, errorReturned := GetUser(user.Email, constants.EMAIL)
	if data != nil {
		log.Printf("User updated successfully")
		return *data, errorReturned
	}
	return models.SystemUser{}, nil
}

func GetLoginAccount(email string) (error, models.SystemUser) {
	var loginAccount models.SystemUser

	err := database.DB.First(&loginAccount, models.SystemUser{Username: email})
	if err != nil {
		log.Printf("Error fetching user: %v", err)
		return err.Error, loginAccount
	}
	return nil, loginAccount
}
