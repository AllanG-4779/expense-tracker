package service

import (
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/allang-4779/financer/internal/models"
	"github.com/allang-4779/financer/internal/repository"
	"github.com/allang-4779/financer/internal/types"
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
	if account.UserID != user.ID || account.ID == 0 {
		return errors.New("account does not belong to user or does not exist")
	}
	var rTransaction models.Transaction
	rTransaction.AccountID = transaction.AccountID
	rTransaction.Amount = transaction.Amount
	rTransaction.CategoryID = category.ID
	rTransaction.Title = transaction.Title
	rTransaction.Description = transaction.Description
	rTransaction.UserID = user.ID
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
	rBudget.Balance = budget.Amount
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

func GetTransactions(request types.FetchRequest, username string) ([]models.Transaction, error) {
	user, err := repository.GetUser(username, "username")
	if err != nil {
		return nil, errors.New("could not retrieve user from context")
	}
	request.Username = strconv.Itoa(int(user.ID))
	transactions, err := repository.GetTransactions(request)
	
	if err != nil {
		return nil, errors.New("could not retrieve transactions")
	}

	return transactions, nil
}

func formatDate(date string) (string, error) {
	expectedFormat := "2006-01-02"
	_, err := time.Parse(expectedFormat, date)
	if err == nil {
		// Already in correct format
		return date, nil
	}

	// Try parsing fallback format: DD-MM-YYYY
	fallbackFormat := "02-01-2006"
	parsedDate, err := time.Parse(fallbackFormat, date)
	if err != nil {
		return "", errors.New("could not parse date")
	}

	return parsedDate.Format(expectedFormat), nil
}

func GetUserAccounts(username string) ([]models.Account, error) {
	user, err := repository.GetUser(username, "username")
	if err != nil {
		return nil, errors.New("could not retrieve user from context")
	}
	accounts, err := repository.GetTransactionAccounts(types.AccountRequest{
		UserId: user.ID,
		Page:   0,
		Size:   10,
	})
	if err != nil {
		return nil, errors.New("could not retrieve accounts")
	}
	return accounts, nil
}
func UpdateTransaction(request types.TransactionRequest, username string) error {
	user, err := repository.GetUser(username, "username")
	if err != nil {
		return errors.New("could not retrieve user from context")
	}

     trxAccount, err := repository.GetAccount(request.AccountID)
	if err != nil {
		return errors.New("could not retrieve account")
	}
	if trxAccount.UserID != user.ID {
		return errors.New("account does not belong to user")
	}
	 category , err := repository.GetCategoryByName(request.CategoryID)
	
	if err != nil {
		return errors.New("could not retrieve category")
	}
	transaction, err := repository.GetTransactionById(request.TransactionID)
	if err != nil {
		return errors.New("could not retrieve transaction")
	}
	
	if request.Amount> 0{		
		// update the account balance if the previous amount was less than the new amount
		if category.Type == "expense"  {
			if (transaction.Amount != request.Amount){
				if (transaction.Amount < request.Amount){
					if trxAccount.Balance < (request.Amount - transaction.Amount) {
						return errors.New("insufficient funds")
					}
					trxAccount.Balance -= (request.Amount - transaction.Amount)
				} else {
					trxAccount.Balance += (transaction.Amount - request.Amount)
				}
			}
		}else if category.Type == "income" {
			if (transaction.Amount != request.Amount){
				if (transaction.Amount < request.Amount){
					trxAccount.Balance += (request.Amount - transaction.Amount)
				} else {
					trxAccount.Balance -= (transaction.Amount - request.Amount)
				}
			}
		} else {
			return errors.New("expense type undefined")
		}
	   
		
	}
	if request.Date != "" {
		date, dateErr := formatDate(request.Date)
		if dateErr != nil {
			return errors.New("could not format date")
		}
		transaction.Date = date
	}
	if request.CategoryId != 0 {
		category, categoryErr := repository.GetCategoryByName(request.CategoryID)
		if categoryErr != nil {
			return errors.New("could not retrieve category")
		}
		transaction.CategoryID = category.ID
	}
	
	if request.Description != "" {
		transaction.Description = request.Description
	}
	if request.Title != "" {
		transaction.Title = request.Title
	}
	repository.UpdateTransactionAccount(*trxAccount)
	return repository.UpdateTransaction(*transaction)
}

func FilterTransactions(request types.FilterRequest, username string) ([]models.Transaction, error) {
	user, err := repository.GetUser(username, "username")
	if err != nil {
		return nil, errors.New("could not retrieve user from context")
	}
	
	if request.AccountID > 0 {
		account, err := repository.GetTransactionsByAccountId(request.AccountID)
		if err != nil {
			return nil, errors.New("could not retrieve account")
		}
		return account, nil
		
	}
	if request.CategoryID > 0 {
		category, err := repository.GetTransactionByCategoryId(request.CategoryID, user.ID)
		if err != nil {
			return nil, errors.New("could not retrieve category")
		}
		return category, nil
	}
	if request.StartDate != "" && request.EndDate != "" {
		startDate, err := formatDate(request.StartDate)
		if err != nil {
			return nil, errors.New("could not format start date")
		}
		endDate, err := formatDate(request.EndDate)
		if err != nil {
			return nil, errors.New("could not format end date")
		}
		category, err := repository.GetTransactionByDate(startDate, endDate, user.ID)
		if err != nil {
			return nil, errors.New("could not retrieve category")
		}

		return category, nil
	}
	return nil, errors.New("could not retrieve transactions")	
	
}

func DeleteTransaction(id uint, username string) error {
	user, err := repository.GetUser(username, "username")
	if err != nil {
		return errors.New("could not retrieve user from context")
	}
	transaction, err := repository.GetTransactionById(id)
	if err != nil {
		return errors.New("could not retrieve transaction")
	}
	if transaction.UserID != user.ID {
		return errors.New("transaction does not belong to user")
	}
	return repository.DeleteTransaction(id)
}
