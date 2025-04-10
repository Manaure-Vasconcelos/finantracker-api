package main

import (
	"api-finantracker/src/controller"
	"api-finantracker/src/db"
	"api-finantracker/src/repositories"
	"api-finantracker/src/usecase"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	dbConnectionPool, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer dbConnectionPool.Close()

	TransactionRepository := repository.NewTransactionRepository(dbConnectionPool)

	TransactionUsecase := usecase.NewTransactionUsecase(TransactionRepository)

	TransactionController := controller.NewTransactionController(TransactionUsecase)

	server.GET("/transaction", TransactionController.GetTransactions)
	server.POST("/transaction", TransactionController.CreateTransaction)

	server.Run(":8080")
}
