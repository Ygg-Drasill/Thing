package handlers

import (
	"net/http"

	"github.com/Ygg-Drasill/Thing/src/features/penalties"
	"github.com/gin-gonic/gin"
)

func GetPenaltyInfo(context *gin.Context) {
	selectedPenaltyName := context.Query("penalty")
	if selectedPenaltyName == "" {
		context.String(http.StatusOK, "Select a penalty to see its info")
		return
	}
	selectedPenalty, found := penalties.FindPenalty(selectedPenaltyName)
	if found {
		context.String(http.StatusOK, selectedPenalty.Info)
	} else {
		context.String(http.StatusOK, "No info available for this penalty")
	}
}
