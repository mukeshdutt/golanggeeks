package main

import (
	"golanggeeks/common"
	"golanggeeks/services"

	"github.com/gin-gonic/gin"
)

func init() {
	common.Initialize()
}

func main() {

	router := gin.Default()

	// Menus routers
	router.POST("/menu", services.AddMenu)
	router.PUT("/menu", services.EditMenu)
	router.DELETE("/menu/:name", services.DeleteMenu)
	router.GET("/menu/:id", services.GetMenu)
	router.GET("/menus", services.GetMenus)

	router.Run(":8080")
}
