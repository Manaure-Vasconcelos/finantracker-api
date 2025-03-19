package controller

import (
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
