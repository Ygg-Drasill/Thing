package handlers

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func LogsHandler(context *gin.Context) {
	file, err := os.Open("src/logs/penalties.log")
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to open logs"})
		return
	}
	defer file.Close()

	logs, err := io.ReadAll(file)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to read logs"})
		return
	}

	context.String(http.StatusOK, string(logs))
}
