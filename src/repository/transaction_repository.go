package repository

import (
	"api-finantracker/src/model"
	"database/sql"
	"fmt"
)

type TransactionRepository struct {
	connection *sql.DB
}

func NewTransactionRepository(connection *sql.DB) TransactionRepository {
	return TransactionRepository{
		connection: connection,
	}
}

func (pr *TransactionRepository) GetTransactions() ([]model.Transaction, error) {

	query := "SELECT * FROM transactions"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Transaction{}, err
	}

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

	rows.Close()

	return transactionList, nil
}
