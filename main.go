package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.GET("/recognition", recognition)
	r.Run()
}

func recognition(c *gin.Context) {
	c.HTML(http.StatusOK, "recognition.html", nil)
}
