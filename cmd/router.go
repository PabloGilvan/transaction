package main

import (
	"github.com/PabloGilvan/transaction/cmd/handlers"
	"github.com/PabloGilvan/transaction/commons"
	"github.com/PabloGilvan/transaction/internal/config/global"
	"github.com/PabloGilvan/transaction/internal/container"
	"github.com/gin-gonic/gin"
)

func StartServer(container container.Dependency) {
	router := gin.Default()
	v1 := router.Group("/v1")

	setupAccountPaths(v1, container)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, "Not found")
	})

	err := router.Run(":" + getPort())

	if err != nil {
		panic(err)
	}
}

func setupAccountPaths(router *gin.RouterGroup, container container.Dependency) {
	var accountController = handlers.NewAccountController(container.Services.AccountService)
	var transactionController = handlers.NewTransactionController(container.Services.TransactionService, container.Services.AccountService)

	accountController.Router(router)
	transactionController.Router(router)
}

func getPort() string {
	var port = global.Viper.GetString("app.port")
	if len(port) == 0 {
		return commons.DEFAULT_PORT
	}
	return port
}
