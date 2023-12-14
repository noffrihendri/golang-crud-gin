package main

import (
	"context"
	"crud-cleancode/databases"
	"crud-cleancode/internal/handler"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	db := databases.NewDBPostgres()

	productHandler := handler.NewProductHandler(db)

	//userHandler := handler.NewUserHandler(db)

	// router.GET("/users", func(c *gin.Context) {
	// 	userHandler.GetUser(c)
	// })
	router.GET("/products", func(c *gin.Context) {
		productHandler.GetProduct(c)
	})

	router.GET("/products/:id", func(c *gin.Context) {
		productHandler.GetProductByID(c)
	})

	router.POST("/products", func(c *gin.Context) {
		productHandler.CreateProduct(c)
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "already running",
		})
	})

	// api := router.Group("/api")
	// {
	// 	api.GET("/ping", func(c *gin.Context) {
	// 		c.JSON(200, gin.H{
	// 			"message": "pong",
	// 		})
	// 	})

	// }

	server := &http.Server{
		Addr:    ":8083", // Change this to your desired port
		Handler: router,
	}

	gin.SetMode(gin.ReleaseMode)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	gracefulShutdown(server)
}

func gracefulShutdown(server *http.Server) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	fmt.Println("\nShutting down gracefully...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("Error shutting down: %v\n", err)
	}
	os.Exit(0)
}
