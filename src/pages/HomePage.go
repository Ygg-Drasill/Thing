package pages

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	data := struct {
		Title  string
		Header string
		Items  []string
	}{
		Title:  "Thing",
		Header: "Welcome to Thing!",
		Items:  []string{"You", "Will", "Pays"},
	}

	c.HTML(http.StatusOK, "index.html", data)
}
