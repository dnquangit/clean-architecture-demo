package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type GinHTTPHandler struct {
	GracefullyShutdown
	Router *gin.Engine
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
		GracefullyShutdown: NewGracefullyShutdown(router, address),
		Router:             router,
	}, nil

}
