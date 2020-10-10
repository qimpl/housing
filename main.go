package main

import (
	_ "github.com/qimpl/housing/docs"
	"github.com/qimpl/housing/router"

	_ "github.com/joho/godotenv/autoload"
)

// @title Housing API
// @version 0.1.0
// @BasePath /api/v1
func main() {
	router.CreateRouter()
}
