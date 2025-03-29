package service

import (
	"errors"
	"github.com/allang-4779/financer/internal/models"
	"github.com/allang-4779/financer/internal/repository"
	"github.com/allang-4779/financer/internal/types"
	"log"
	"time"
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
	category, categoryErr := repository.GetCategoryByName(transaction.CategoryID)
	if categoryErr != nil {
		return errors.New("could not retrieve category")
	}
	log.Printf("Adding transaction to user %v account %v", user, account)
	if account.UserID != user.ID {
		return errors.New("account does not belong to user")
	}
	var rTransaction models.Transaction
	rTransaction.AccountID = transaction.AccountID
	rTransaction.Amount = transaction.Amount
	rTransaction.CategoryID = category.ID
	rTransaction.Description = transaction.Description
	// Formatted date
	date, dateErr := formatDate(transaction.Date)
	if dateErr != nil {
		return errors.New("could not format date")
	}
	rTransaction.Date = date
	rTransaction.Type = category.Type

	return repository.AddTransaction(rTransaction)

}

func CreateBudget(budget types.BudgetRequest, username string) error {
	log.Printf("Creating budget %v for user %v", budget, username)
	user, err := repository.GetUser(username, "username")
	if err != nil {
		return errors.New("could not retrieve user from context")
	}
	category, categoryErr := repository.GetCategoryByName(budget.CategoryName)
	if categoryErr != nil {
		return errors.New("could not retrieve category")
	}
	var rBudget models.Budget
	rBudget.Amount = budget.Amount
	rBudget.Balance = budget.Balance
	rBudget.CategoryID = category.ID
	rBudget.UserID = user.ID
	// Formatted date
	startDate, startDateErr := formatDate(budget.StartDate)
	if startDateErr != nil {
		return errors.New("could not format start date")
	}
	rBudget.StartDate = startDate
	// Formatted date
	endDate, endDateErr := formatDate(budget.EndDate)
	if endDateErr != nil {
		return errors.New("could not format end date")
	}
	rBudget.EndDate = endDate
	return repository.CreateBudget(rBudget)
}

func formatDate(date string) (string, error) {
	format := "02-01-2006"
	formatted, err := time.Parse(format, date)
	if err != nil {
		return "", errors.New("could not parse date")
	}
	date = formatted.Format("2006-01-02")
	return date, nil
}
