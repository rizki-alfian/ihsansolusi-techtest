package main

import (
	"github.com/joho/godotenv"
	"os"
	"ihsansolusi-account/apps/routes"
	"ihsansolusi-account/apps/cores"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
        log.Fatal("Warning: .env file not found, using system environment variables")
    }

	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "local"
	}

	server := routes.NewServer()

	cores.RegisterMiddlewares(server.Echo)

	APP_PORT := os.Getenv("APP_PORT")
	if APP_PORT == "" {
        log.Fatal("APP_PORT not set in environment variables")
    }
	server.Echo.Start(":" + APP_PORT)
}