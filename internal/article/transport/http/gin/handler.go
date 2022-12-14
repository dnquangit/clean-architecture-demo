package gin

import (
	"github.com/gin-gonic/gin"
	"go-module/internal/article"
	"go-module/internal/article/model"
	"go-module/internal/article/viewmodel"
	"net/http"
)

type articleHandler struct {
	postUC article.UseCase
}

func NewPostHandler(articleUC article.UseCase) *articleHandler {
	return &articleHandler{postUC: articleUC}
}

func (handler *articleHandler) GetArticles() gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter model.Filter
		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		posts, _ := handler.postUC.ListArticles(c.Request.Context(), &model.Filter{Published: true})

		/* swap to article response */
		var postResponse []viewmodel.ArticleResponse
		for _, m := range *posts {
			postResponse = append(postResponse, viewmodel.ArticleResponse{
				Id:           m.Id,
				Title:        m.Title,
				ShortContent: m.ShortContent,
				UpdateAt:     m.UpdateAt,
			})
		}

		/* return */
		c.JSON(http.StatusOK, &postResponse)
	}
}
