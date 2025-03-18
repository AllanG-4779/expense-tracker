package repository

import (
	"errors"
	"github.com/allang-4779/financer/internal/constants"
	"github.com/allang-4779/financer/internal/database"
	"github.com/allang-4779/financer/internal/models"
	"github.com/allang-4779/financer/internal/types"
)

func CreateTransactionAccount(request models.Account) error {
	_, err := database.DB.NamedExec(database.CreateAccount, request)
	if err != nil {
		return err
	}
	return nil
}

func UpdateTransactionAccount(request types.AccountRequest) error {
	_, err := database.DB.NamedExec(database.UpdateAccount, request)
	if err != nil {
		return err
	}
	return nil
}

func GetUserAccount(id uint) (*models.Account, error) {
	var account models.Account
	err := database.DB.Get(&account, database.GetAccount, id)
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func GetTransactionAccounts(request types.AccountRequest) ([]types.AccountRequest, error) {
	var accounts []types.AccountRequest
	size := request.Size
	offset := request.Size * (request.Page)
	err := database.DB.Select(&accounts, database.GetAccounts, request.UserId, size, offset)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func AddTransaction(request models.Transaction) error {
	tr, err := database.DB.Beginx()
	if err != nil {
		return errors.New(constants.DatabaseTransactionError)
	}
	var account models.Account
	err = tr.Get(&account, database.GetAccount, request.AccountID)
	if err != nil {
		rError := tr.Rollback()
		if rError != nil {
			return errors.New(constants.DatabaseTransactionError)
		}
	}
	if account.Balance < request.Amount {
		rError := tr.Rollback()
		if rError != nil {
			return errors.New(constants.DatabaseTransactionError)
		}
		return errors.New(constants.InsufficientFunds)
	}
	_, err = tr.NamedExec(database.InsertTransaction, request)
	if err != nil {
		rError := tr.Rollback()
		if rError != nil {
			return errors.New(constants.DatabaseTransactionError)
		}
	}
	account.Balance -= request.Amount
	_, err = tr.NamedExec(database.UpdateAccount, account)
	return err
}
