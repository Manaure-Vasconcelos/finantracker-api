package main

import (
	"api-finantracker/src/controller"
	"api-finantracker/src/db"
	"api-finantracker/src/repository"
	"api-finantracker/src/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	TransactionRepository := repository.NewTransactionRepository(dbConnection)

	TransactionUsecase := usecase.NewTransactionUsecase(TransactionRepository)

	TransactionController := controller.NewTransactionController(TransactionUsecase)

	server.GET("/transaction", TransactionController.GetTransactions)

	server.Run(":8080")
}
