package gin

import "github.com/gin-gonic/gin"

func MapRoutes(router *gin.Engine, h *articleHandler) {
	router.GET("/api/v1/article/find", h.GetArticles())
}
