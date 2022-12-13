package main

import (
	server "go-module/internal/server/http/gin"
)

func main() {
	server.NewServer().Run()
}
