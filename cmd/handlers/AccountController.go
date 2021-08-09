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

// CreateAccount @Title createAccount
// @Tags Accounts
// @Summary Creates an Account
// @Description Creates a new Account generating am account number
// @Param content body account.AccountPersist true "Object for persisting the account"
// @Success 201 {object} account.AccountResponse
// @Failure 400 "Bad request"
// @Accept json
// @Router /accounts [post]
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

// LoadAccount @Title loadAccount
// @Tags Accounts
// @Summary Load an account
// @Description Load an account by the UUID identifier
// @Description Accounts resources can be used to create ("/accounts" POST)
// @Param id path string true "Person's identification code"
// @Success 201 {object} account.AccountResponse
// @Failure 400 "account inactive"
// @Failure 404 "account not found"
// @Accept json
// @Router /accounts/{id} [get]
func (crtl AccountController) LoadAccount(c *gin.Context) {
	accountID := c.Param("id")

	response, err := crtl.AccountService.LoadAccount(context.Background(), accountID)

	if errMessage, statusCode := helpers.ProcessIfBusinessError(err); errMessage != nil {
		c.IndentedJSON(statusCode, errMessage)
		return
	}

	c.IndentedJSON(http.StatusOK, response)
}
