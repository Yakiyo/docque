package api

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/Yakiyo/docque/db"
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

func allDocs(ctx *gin.Context) {
	docs, err := db.Client.Doctor.FindMany().With(db.Doctor.Appointments.Fetch()).Exec(context.Background())
	if err != nil {
		errResponse(ctx, err)
		return
	}
	output := []string{}
	for _, doc := range docs {
		as := iter_map(doc.Appointments(), atoqa)
		d := QueueDoc{
			Id:           doc.ID,
			Name:         doc.Name,
			Appointments: as,
		}
		buff := &bytes.Buffer{}
		queueTmpl.Execute(buff, d)
		output = append(output, buff.String())
	}
	wrap := fmt.Sprintf(
		`<div class="flex w-full flex-col items-center justify-between">%s</div>`,
		strings.Join(output, "\n\n"),
	)
	ctx.String(http.StatusOK, wrap)
}

func errResponse(ctx *gin.Context, err error) {
	log.Error("Unexpected error", "err", err)
	ctx.String(http.StatusInternalServerError, fmt.Sprintf(`<p>ERROR: %s</p>`, err))
}
