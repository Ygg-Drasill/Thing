package handlers

import (
	"log"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/Ygg-Drasill/Thing/src/features/logs"
	"github.com/Ygg-Drasill/Thing/src/features/penalties"
	"github.com/Ygg-Drasill/Thing/src/utils"
)

func SubmitHandler(context *gin.Context) {
	session := sessions.Default(context)
	person := context.PostForm("person")
	penaltyStr := context.PostForm("penalty")

	penalty, ok := penalties.FindPenalty(penaltyStr)
	if !ok {
		log.Printf("Invalid penalty: %v", penaltyStr)
		return
	}
	penaltyAmount := penalty.Value

	currentPenalty := utils.GetPenalty(session, person)
	penaltyAmount += currentPenalty

	session.Set("penalty_"+person, strconv.Itoa(penaltyAmount))

	logs.LogIP(context.ClientIP())
	logs.LogPerson(person)
	logs.LogPenaltyString(penaltyStr)
	logs.LogPenaltyAmount(penalty.Value)
	logs.LogOldPenaltyTotal(currentPenalty)
	logs.LogNewPenaltyTotal(penaltyAmount)

	session.Set("submitted", true)
	session.Save()

	HomePage(context)
}
