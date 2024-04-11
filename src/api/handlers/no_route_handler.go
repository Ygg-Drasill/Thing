package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NoRouteHandler(context *gin.Context) {
	context.HTML(http.StatusNotFound, "404.html", nil)
}
