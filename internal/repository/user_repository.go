package repository

import (
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

func GetUserByEmail(email string) (*models.SystemUser, error) {
	var user models.SystemUser
	err := database.DB.Get(&user, database.GetUserByEmail, email) // FIXED
	if err != nil {
		log.Printf("Error fetching user: %v", err)
		return nil, err
	}
	log.Printf("User: %v", user)
	return &user, nil
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
