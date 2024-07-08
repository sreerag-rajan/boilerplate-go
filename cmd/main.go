package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	// "github.com/sreerag-rajan/boilerplate-go/config"
	"github.com/sreerag-rajan/boilerplate-go/internal"
	databaseimplementation "github.com/sreerag-rajan/boilerplate-go/pkg/database/database_implementation"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file", err)
	}

	r := gin.Default()

	ctx := context.Background()
	db, err := databaseimplementation.GetDatabase(ctx)

	fmt.Println(db) // this is temperory

	if err != nil {
		fmt.Println("Error initializing database", err)
		os.Exit(1)
	}

	routes := internal.NewRoutes(r)
	routes.RegisterRoutes()

	PORT := os.Getenv("PORT")
	r.Run(":" + PORT)
}
