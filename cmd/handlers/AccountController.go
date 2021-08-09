package handlers

import (
	"context"
	"github.com/PabloGilvan/transaction/cmd/helpers"
	"github.com/PabloGilvan/transaction/internal/services/account"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AccountController struct {
	AccountService account.AccountService
}

func NewAccountController(service account.AccountService) AccountController {
	return AccountController{
		AccountService: service,
	}
}

func (crtl AccountController) Router(router *gin.RouterGroup) {
	routerGroup := router.Group("/accounts")
	{
		routerGroup.POST("/", crtl.CreateAccount)
		routerGroup.GET("/:id", crtl.LoadAccount)
	}
}

func (crtl AccountController) CreateAccount(c *gin.Context) {

	var accountPersist account.AccountPersist
	if err := c.BindJSON(&accountPersist); err != nil {
		c.IndentedJSON(http.StatusBadRequest, helpers.ErrorMessage{ErrorMessage: helpers.ErrInvalidRequest.Error()})
		return
	}

	response, err := crtl.AccountService.SaveAccount(context.Background(), accountPersist)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	c.IndentedJSON(http.StatusCreated, response)
}

func (crtl AccountController) LoadAccount(c *gin.Context) {
	accountID := c.Param("id")

	response, err := crtl.AccountService.LoadAccount(context.Background(), accountID)

	if errMessage, statusCode := helpers.ProcessIfBusinessError(err); errMessage != nil {
		c.IndentedJSON(statusCode, errMessage)
		return
	}

	c.IndentedJSON(http.StatusOK, response)
}
