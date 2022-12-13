package gin

import (
	"github.com/gin-gonic/gin"
	"go-module/internal/post"
	"go-module/internal/post/model"
	"go-module/internal/post/viewmodel"
	"net/http"
)

type postHandler struct {
	postUC post.UseCase
}

func NewPostHandler(postUC post.UseCase) *postHandler {
	return &postHandler{postUC: postUC}
}

func (handler *postHandler) GetPosts() gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter model.Filter
		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		posts, _ := handler.postUC.ListPosts(c.Request.Context(), &model.Filter{Published: true})

		/* swap to post response */
		var postResponse []viewmodel.PostResponse
		for _, m := range *posts {
			postResponse = append(postResponse, viewmodel.PostResponse{
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
