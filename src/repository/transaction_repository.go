package repository

import (
	"api-finantracker/src/model"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TransactionRepository struct {
	connection *pgxpool.Pool
}

func NewTransactionRepository(connection *pgxpool.Pool) TransactionRepository {
	return TransactionRepository{
		connection: connection,
	}
}

func (pr *TransactionRepository) GetTransactions() ([]model.Transaction, error) {

	rows, err := pr.connection.Query(context.Background(), "SELECT * FROM transactions")
	if err != nil {
		fmt.Println(err)
		return []model.Transaction{}, err
	}
	defer rows.Close()

	var transactionList []model.Transaction
	var transactionObj model.Transaction

	for rows.Next() {
		err = rows.Scan(&transactionObj.ID, &transactionObj.Name, &transactionObj.Description, &transactionObj.Amount, &transactionObj.Type, &transactionObj.Category, &transactionObj.Date)
		if err != nil {
			fmt.Println(err)
			return []model.Transaction{}, err
		}

		transactionList = append(transactionList, transactionObj)
	}

	return transactionList, nil
}

func (pr *TransactionRepository) CreateTransaction(transaction model.Transaction) error {

	_, err := pr.connection.Exec(context.Background(), "INSERT INTO transactions (name, description, amount, type, category, date) VALUES ($1, $2, $3, $4, $5, $6)", transaction.Name, transaction.Description, transaction.Amount, transaction.Type, transaction.Category, transaction.Date)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
