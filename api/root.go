package api

import (
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

func Register(router *gin.RouterGroup) {
	router.GET("/hello", func(ctx *gin.Context) {
		log.Info("in hello", "headers", ctx.Request.Header)
		ctx.String(http.StatusOK, `<p>Hello World</p>`)
	})

	router.GET("/all", allDocs)
}

func allDocs(ctx *gin.Context) {}
