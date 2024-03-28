package main

import (
	. "github.com/Ygg-Drasill/Thing/src/api/handlers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	router.Static("/static", "./src/assets")

	router.LoadHTMLGlob("./src/templates/*.html")

	router.GET("/", HomePage)
	router.POST("/submit", SubmitHandler)

	router.Run("localhost:8080")
}
