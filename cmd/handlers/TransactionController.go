package handlers

import (
	"context"
	"github.com/PabloGilvan/transaction/cmd/helpers"
	"github.com/PabloGilvan/transaction/internal/services/account"
	"github.com/PabloGilvan/transaction/internal/services/transaction"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TransactionController struct {
	TransactionService transaction.TransactionService
	AccountService     account.AccountService
}

func NewTransactionController(service transaction.TransactionService, accountService account.AccountService) TransactionController {
	return TransactionController{
		TransactionService: service,
		AccountService:     accountService,
	}
}

func (crtl TransactionController) Router(router *gin.RouterGroup) {
	transactions := router.Group("/transactions")
	{
		transactions.POST("/", crtl.SaveTransaction)
	}
}

// SaveTransaction @Title saveTransaction
// @Tags Transaction
// @Summary Persist a transaction
// @Description Persist a transaction validating the operation type and account
// @Param content body transaction.TransactionPersist true "Object for persisting the transaction"
// @Success 201 {object} transaction.TransactionResponse
// @Failure 400 "operation not found"
// @Failure 400 "account inactive"
// @Failure 404 "account not found"
// @Accept json
// @Router /transactions/{id} [get]
func (crtl TransactionController) SaveTransaction(c *gin.Context) {

	var transactionPersist transaction.TransactionPersist
	if err := c.BindJSON(&transactionPersist); err != nil {
		c.IndentedJSON(http.StatusBadRequest, helpers.ErrorMessage{ErrorMessage: helpers.ErrInvalidRequest.Error()})
		return
	}

	_, err := crtl.AccountService.LoadAccount(context.Background(), transactionPersist.AccountId)
	if errMessage, statusCode := helpers.ProcessIfBusinessError(err); errMessage != nil {
		c.IndentedJSON(statusCode, errMessage)
		return
	}

	transactionIdentifier, err := crtl.TransactionService.SaveTransaction(context.Background(), transactionPersist)
	if errMessage, statusCode := helpers.ProcessIfBusinessError(err); errMessage != nil {
		c.IndentedJSON(statusCode, errMessage)
		return
	}

	c.IndentedJSON(http.StatusAccepted, transactionIdentifier)
}
