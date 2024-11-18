package main

import (
	"log"
	"net/http"

	"github.com/sithsithsith/mugi-cognito/internal/app"
	"github.com/sithsithsith/mugi-cognito/pkg/logger"
)

func main() {
	logger.Init()
	app := app.NewApp()
	app.RegisterRoutes()

	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
