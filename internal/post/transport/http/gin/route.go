package gin

import "github.com/gin-gonic/gin"

func MapRoutes(router *gin.Engine, h *postHandler) {
	router.GET("/api/v1/post/find", h.GetPosts())
}
