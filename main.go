package main

import (
	"log"

	_ "github.com/qimpl/housing/docs"
	"github.com/qimpl/housing/router"

	"github.com/joho/godotenv"
)

// @title housing API
// @version 0.1.0
// @BasePath /api/v1
func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file", err)
	}

	router.CreateRouter()
}
