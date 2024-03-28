package handlers

import (
	"net/http"
	"strings"

	"github.com/Ygg-Drasill/Thing/src/features/logs"
	"github.com/gin-gonic/gin"
)

func LogsHandler(context *gin.Context) {
	logsData := strings.Join(logs.GetLogs(), "\n")
	context.String(http.StatusOK, logsData)
}
