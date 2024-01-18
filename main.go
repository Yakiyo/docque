package main

import (
	"os"

	"github.com/Yakiyo/docque/api"
	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

func main() {
	port := "8080"
	if val := os.Getenv("PORT"); val != "" {
		port = val
	}
	// log stuff to stderr, not stdout
	gin.DefaultWriter = os.Stderr
	
	app := gin.Default()
	
	// serve all files in /public
	app.Static("/", "./public")

	router := app.Group("/api")

	api.Register(router)

	log.Info("Starting server", "port", port)
	err := app.Run(":" + port)
	if err != nil {
		log.Fatal("Error when starting server", "err", err)
	}
}
