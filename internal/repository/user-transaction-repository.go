package repository

import (
	"errors"
	"strconv"
	"strings"

	"github.com/allang-4779/financer/internal/constants"
	"github.com/allang-4779/financer/internal/database"
	"github.com/allang-4779/financer/internal/models"
	"github.com/allang-4779/financer/internal/types"
	"gorm.io/gorm"
)

func CreateTransactionAccount(request models.Account) error {
	err := database.DB.Create(&request)
	if err != nil {
		return err.Error
	}
	return nil
}

func UpdateTransactionAccount(request models.Account) error {

	err := database.DB.Updates(request)
	if err != nil {
		return err.Error
	}
	return nil
}

func GetUserAccount(id uint) (*models.Account, error) {
	var account models.Account
	err := database.DB.Find(&account, id).Error
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func GetTransactionAccounts(request types.AccountRequest) ([]models.Transaction, error) {
	var accounts []models.Transaction
	size := request.Size
	offset := request.Size * (request.Page)

	err := database.DB.Limit(size).Offset(offset).Find(&accounts).Error
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func AddTransaction(request models.Transaction) error {

	return database.DB.Transaction(func(tx *gorm.DB) error {
		var account models.Account
		err := tx.Find(&account, request.AccountID)
		if err.Error != nil {
			return err.Error
		}
		if account.ID == 0 {
			return errors.New("account does not exist")
		}
		if account.UserID != request.UserID {
			return errors.New("account does not belong to user")
		}
		if strings.ToLower(request.Type) == constants.EXPENSE {
			if account.Balance < request.Amount {
				return errors.New(constants.InsufficientFunds)
			}
			account.Balance -= request.Amount
		} else if strings.ToLower(request.Type) == constants.INCOME {
			account.Balance += request.Amount

		} else {
			return errors.New("expense type undefined")
		}
		result := tx.Save(&account)
		if result.Error != nil {
			return result.Error
		}
		result = tx.Create(&request)
		if result.Error != nil {
			return result.Error
		}
		return nil
	})
}

func CreateBudget(budget models.Budget) error {
	err := database.DB.Create(&budget)
	if err != nil {
		return err.Error
	}
	return nil
}
func GetTransactions(request types.FetchRequest) ([]models.Transaction, error) {
	var transactions []models.Transaction
	size := request.Size
	offset := request.Size * (request.Page)
	userId, errEnt := strconv.Atoi(request.Username)
	if errEnt != nil {
		return nil, errors.New("could not convert user id")
	}
	err := database.DB.Limit(size).Offset(offset).
		Where("user_id = ? AND deleted_at IS NULL", userId).
		Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	return transactions, nil
}
