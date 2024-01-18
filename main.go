package main

import (
	"os"

	"github.com/Yakiyo/docque/api"
	"github.com/Yakiyo/docque/db"
	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

func run() error {
	port := "8080"
	if val := os.Getenv("PORT"); val != "" {
		port = val
	}

	if err := db.Init(); err != nil {
		return err
	}

	defer db.Close()

	// log stuff to stderr, not stdout
	gin.DefaultWriter = os.Stderr

	app := gin.Default()

	// serve all files in /public
	app.StaticFile("/", "./public/index.html")

	app.GET("/doc/:id", func(c *gin.Context) {
		c.File("./public/doc.html")
	})

	router := app.Group("/api")

	api.Register(router)

	log.Info("Starting server", "port", port)
	return app.Run(":" + port)
}

func main() {
	if err := run(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
