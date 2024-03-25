package main

import (
	. "github.com/Ygg-Drasill/Thing/src/pages"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Static("/static", "./assets")

	router.LoadHTMLGlob("./templates/*.html")

	router.GET("/", HomePage)

	router.Run("localhost:8080")
}
