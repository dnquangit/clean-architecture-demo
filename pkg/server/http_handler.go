package server

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-module/pkg/log"
	"net/http"
	"os"
	"time"
)

type GinHTTPHandler struct {
	Router  *gin.Engine
	Address string
}

func NewGinHTTPHandler(address string) (GinHTTPHandler, error) {

	router := gin.Default()

	// PING API
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "Ready")
	})

	// CORS
	router.Use(cors.New(cors.Config{
		ExposeHeaders:   []string{"Data-Length"},
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "OPTIONS"},
		AllowAllOrigins: true,
		AllowHeaders:    []string{"Content-Type", "Authorization"},
		MaxAge:          12 * time.Hour,
	}))

	return GinHTTPHandler{
		Router:  router,
		Address: address,
	}, nil

}

func (r *GinHTTPHandler) Run() {

	if err := r.Router.Run(r.Address); err != nil && err != http.ErrServerClosed {
		log.Error(context.Background(), "listen: %s", err)
		os.Exit(1)
	}

	log.Info(context.Background(), "server is running ...")
}
