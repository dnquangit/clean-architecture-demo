package gin

import (
	"github.com/gin-gonic/gin"
	postRepo "go-module/internal/post/repo"
	postTranport "go-module/internal/post/transport/http/gin"
	postUseCase "go-module/internal/post/usecase"
)

type server struct {
	ginEngine *gin.Engine
}

func NewServer() *server {
	return &server{}
}

func (s *server) Run() {
	s.ginEngine = gin.Default()

	// Init repositories

	// Init repositories
	postRepo := postRepo.NewSQLRepo()

	// Init useCases
	postUC := postUseCase.NewPostUC(postRepo)

	// Init handlers
	postHandler := postTranport.NewPostHandler(postUC)
	postTranport.MapRoutes(s.ginEngine, postHandler)

	s.ginEngine.Run(":8000")
}
