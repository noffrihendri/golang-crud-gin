package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/noffrihendri/golang-crud-gin.git/databases"
	"github.com/noffrihendri/golang-crud-gin.git/internal/handler"
)

func Router() *gin.Engine {
	router := gin.Default()
	db := databases.NewDBPostgres()

	productHandler := handler.NewProductHandler(db)
	router.GET("/products", func(c *gin.Context) {
		productHandler.GetProduct(c)
	})

	router.GET("/products/:id", func(c *gin.Context) {
		productHandler.GetProductByID(c)
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "already running",
		})
	})
	return router
}
