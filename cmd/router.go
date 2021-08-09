package main

import (
	"github.com/PabloGilvan/transaction/cmd/handlers"
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

	account := router.Group("/accounts")
	{
		account.POST("/", accountController.CreateAccount)
		account.GET("/:id", accountController.LoadAccount)
	}
}

func getPort() string {
	return "8080"
}
