package service

import (
	"errors"
	"github.com/allang-4779/financer/internal/models"
	"github.com/allang-4779/financer/internal/repository"
	"github.com/allang-4779/financer/internal/types"
	"log"
)

func ActivateAccount(account types.AccountRequest, username string) error {
	user, err := repository.GetUser(username, "username")
	if err != nil {
		return errors.New("could not retrieve user from context")
	}
	userAccount := models.Account{
		Name:    account.Name,
		Balance: 0,
		UserID:  user.ID,
	}
	log.Printf("User account: %v", userAccount)
	return repository.CreateTransactionAccount(userAccount)
}

func AddTransaction(transaction types.TransactionRequest, username string) error {
	user, err := repository.GetUser(username, "username")
	if err != nil {
		return errors.New("could not retrieve user from context")
	}
	account, err := repository.GetUserAccount(transaction.AccountID)
	if err != nil {
		return errors.New("could not retrieve account")
	}
	if account.UserID != user.ID {
		return errors.New("account does not belong to user")
	}
	var rTransaction models.Transaction
	rTransaction.AccountID = transaction.AccountID
	rTransaction.Amount = transaction.Amount
	rTransaction.CategoryID = transaction.CategoryID
	rTransaction.Description = transaction.Description
	rTransaction.Date = transaction.Date
	rTransaction.Type = transaction.Type

	return repository.AddTransaction(rTransaction)

}
