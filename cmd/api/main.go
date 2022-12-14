package main

import (
	"fmt"
	"go-module/cmd/api/config"
	"go-module/internal/article/repo"
	"go-module/internal/article/transport/http/gin"
	"go-module/internal/article/usecase"
	"go-module/pkg/database"
	"go-module/pkg/server"
)

func main() {
	cfg, _ := config.NewConfig()

	db, _ := database.PostgresDB(cfg.PG.URL)

	/* article */
	articleRepo := repo.NewSQLRepo(db)
	articleUC := usecase.NewArticleUC(articleRepo)
	articleHandler := gin.NewPostHandler(articleUC)

	/* handler */
	httpHandler, _ := server.NewGinHTTPHandler(fmt.Sprintf(":%v", cfg.Port))

	/* map article to handler */
	gin.MapRoutes(httpHandler.Router, articleHandler)

	httpHandler.Run()
}
