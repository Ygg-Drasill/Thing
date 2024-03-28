package handlers

import (
	"net/http"
	"strings"

	"github.com/Ygg-Drasill/Thing/src/features/logs"
	"github.com/gin-gonic/gin"
)

func LogsHandler(c *gin.Context) {
	logsData := strings.Join(logs.GetLogs(), "\n")
	c.String(http.StatusOK, logsData)
}
