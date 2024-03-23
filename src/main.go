package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Static("/static", "./src/assets")

	router.LoadHTMLGlob("./src/templates/*.html")

	router.GET("/", HomePage)

	router.Run("localhost:8080")
}

func HomePage(c *gin.Context) {
	data := struct {
		Title  string
		Header string
		Items  []string
	}{
		Title:  "Thing",
		Header: "Welcome to Thing!",
		Items:  []string{"You", "Will", "Pay"},
	}

	c.HTML(http.StatusOK, "index.html", data)
}
