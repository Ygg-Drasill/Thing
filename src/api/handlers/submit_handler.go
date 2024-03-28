package handlers

import (
	"log"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/Ygg-Drasill/Thing/src/features/logs"
	"github.com/Ygg-Drasill/Thing/src/features/penalties"
)

func SubmitHandler(context *gin.Context) {
	session := sessions.Default(context)
	person := context.PostForm("person")
	penaltyStr := context.PostForm("penalty")

	penalty, ok := penalties.PenaltyMap[penaltyStr]
	if !ok {
		log.Printf("Invalid penalty: %v", penaltyStr)
		return
	}
	penaltyAmount := penalty

	currentPenaltyStr := session.Get("penalty_" + person)
	if currentPenaltyStr != nil {
		currentPenalty, err := strconv.Atoi(currentPenaltyStr.(string))
		if err != nil {
			log.Printf("Error converting current penalty to integer: %v", err)
			return
		}
		penalty += currentPenalty
	}

	session.Set("penalty_"+person, strconv.Itoa(penalty))

	logs.LogIP(context.ClientIP())
	logs.LogPerson(person)
	logs.LogPenaltyString(penaltyStr)
	logs.LogPenaltyAmount(penaltyAmount)
	logs.LogCurrentPenaltyString(currentPenaltyStr)
	logs.LogPenaltyFromMap(penalty)
	session.Save()

	PenaltiesHandler(context)
}
