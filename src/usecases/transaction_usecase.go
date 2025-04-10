package usecase

import (
	"api-finantracker/src/model"
	"api-finantracker/src/repositories"
)

type TransactionUsecase struct {
	repository repository.TransactionRepository
}

func NewTransactionUsecase(repo repository.TransactionRepository) TransactionUsecase {
	return TransactionUsecase{
		repository: repo,
	}
}

func (pu *TransactionUsecase) GetTransactions() ([]model.Transaction, error) {
	return pu.repository.GetTransactions()
}

func (pu *TransactionUsecase) CreateTransaction(transaction model.Transaction) error {
	return pu.repository.CreateTransaction(transaction)
}