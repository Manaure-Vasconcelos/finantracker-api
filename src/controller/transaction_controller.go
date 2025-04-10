package controller

import (
	"api-finantracker/src/model"
	"api-finantracker/src/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	transactionUsecase usecase.TransactionUsecase
}

func NewTransactionController(usecase usecase.TransactionUsecase) TransactionController {
	return TransactionController{
		transactionUsecase: usecase,
	}
}

func (p *TransactionController) GetTransactions(ctx *gin.Context) {
	transactions, err := p.transactionUsecase.GetTransactions()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, transactions)
}

func (p * TransactionController) CreateTransaction(ctx *gin.Context) {
	var transaction model.Transaction
	if err := ctx.ShouldBindJSON(&transaction); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := p.transactionUsecase.CreateTransaction(transaction)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, transaction)
}
