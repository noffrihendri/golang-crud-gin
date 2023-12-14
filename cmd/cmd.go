package cmd

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/noffrihendri/golang-crud-gin.git/config"
	"github.com/noffrihendri/golang-crud-gin.git/internal/handler"
	"github.com/noffrihendri/golang-crud-gin.git/internal/handler/api/controller"
)

func Execute() {
	router := controller.Router()
	cfg := config.MainConfig{}

	server := &http.Server{
		Addr:    cfg.Server.Port,
		Handler: router,
	}
	gin.SetMode(gin.ReleaseMode)

	handler.GracefulShutdown(server)
}
