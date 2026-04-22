package api

import (
	"embed"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed web/*
var webFS embed.FS

func RegisterUI(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		data, _ := webFS.ReadFile("web/index.html")
		c.Data(http.StatusOK, "text/html; charset=utf-8", data)
	})

	r.GET("/static/:file", func(c *gin.Context) {
		name := c.Param("file")
		data, err := webFS.ReadFile("web/" + name)
		if err != nil {
			c.Status(404)
			return
		}

		switch {
		case name[len(name)-4:] == ".css":
			c.Data(http.StatusOK, "text/css", data)
		case name[len(name)-3:] == ".js":
			c.Data(http.StatusOK, "application/javascript", data)
		default:
			c.Data(http.StatusOK, "application/octet-stream", data)
		}
	})
}
