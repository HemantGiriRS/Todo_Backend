package main

import (
	"Todo/database"
	"Todo/server"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func main() {
	//load env
	err := godotenv.Load()
	if err != nil {
		logrus.Println("Error loading .env file")
	}

	// connect to db and migrate
	fmt.Println("Connecting to database...")
	if err := database.ConnectAndMigrate(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		database.SSLModeDisable); err != nil {
		logrus.Errorf("Failed to initialize and migrate database with error: %v", err)
	}

	//defer database.Todo.Close()

	// run server
	r := server.SetupRoutes()
	fmt.Println("server running on port http://localhost:8080")
	if err := http.ListenAndServe(":"+string(os.Getenv("PORT")), r); err != nil {
		logrus.Errorf("Failed to start server with error: %v", err)
	}
}
