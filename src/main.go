package main

import (
	"github.com/Ygg-Drasill/Thing/src/api/handlers"
	"github.com/Ygg-Drasill/Thing/src/features/logs"
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

	router.GET("/", handlers.HomePage)
	router.POST("/submit", handlers.SubmitHandler)
	router.GET("/logs", handlers.LogsHandler)
	router.GET("/get-penalty-info", handlers.GetPenaltyInfo)

	defer logs.Close()

	router.Run(":8080")
}
