package repository

import (
	"github.com/allang-4779/financer/internal/models"
	"github.com/allang-4779/financer/internal/types"
	"log"
)

/**
* This package contains the repository layer for the user entity
 */

func CreateUser(user *types.UserRegistration) error {

	userEntity := models.SystemUser{
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		Dob:         user.Dob,
		Residential: user.Residential,
		Phone:       user.Phone,
	}
	log.Print(userEntity)

	//err:= database.DB.Create(&userEntity).Error
	//if (err != nil){
	//	return err
	//}
	//
	//err = CreateLoginAccount(user.Email, user.Password)
	//if (err != nil){
	//	log.Print(err)
	//	return err
	//}
	return nil

}

func GetUserByEmail(email string) (models.SystemUser, error) {
	var user models.SystemUser
	//err := database.DB.Where("email = ?", email).First(&user).Error
	return user, nil
}

func CreateLoginAccount(email string, password string) error {
	//	user, err := GetUserByEmail(email)
	//	loginAccount := models.LoginAccount{
	//		Username:   email,
	//		Password:   password,
	//		FirstLogin: true,
	//		LastLogin:  time.Now(),
	//	}
	//	if err != nil {
	//		return err
	//	}
	//	loginAccount.UserID =
	//	log.Println("User ID: ")
	//	return database.DB.Create(&loginAccount).Error
	return nil
}

func GetLoginAccountByEmail(email string) (error, models.LoginAccount) {
	var loginAccount models.LoginAccount
	//err := database.DB.Where("username = ?", email).First(&loginAccount)
	//if (err!=nil){
	//	return err.Error, loginAccount
	//}
	return nil, loginAccount

}
