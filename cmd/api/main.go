package main

import (
	"go-module/internal/post/repo"
	"go-module/internal/post/transport/http/gin"
	"go-module/internal/post/usecase"
	"go-module/pkg/database"
	"go-module/pkg/server"
)

func main() {
	db, _ := database.PostgresDB("host=postgresql.14.2 port=5432 dbname=Blog user=postgres password=12345 sslmode=disable")

	/* post */
	postRepo := repo.NewSQLRepo(db)
	postUC := usecase.NewPostUC(postRepo)
	postHandler := gin.NewPostHandler(postUC)

	/* handler */
	httpHandler, _ := server.NewGinHTTPHandler(":8080")

	/* map post to handler */
	gin.MapRoutes(httpHandler.Router, postHandler)

	httpHandler.RunWithGracefullyShutdown()
}
