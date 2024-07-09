package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/sreerag-rajan/boilerplate-go/config"
	"github.com/sreerag-rajan/boilerplate-go/internal"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file", err)
	}

	err = config.DBInit()
	if err != nil {
		fmt.Println("Error initializing database", err)
		os.Exit(1)
	}

	r := gin.Default()

	if err != nil {
		fmt.Println("Error initializing database", err)
		os.Exit(1)
	}

	routes := internal.NewRoutes(r)
	routes.RegisterRoutes()

	PORT := os.Getenv("PORT")
	r.Run(":" + PORT)
}
