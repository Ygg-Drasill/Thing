package handlers

import (
	"log"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/Ygg-Drasill/Thing/src/penalties"
)

func SubmitHandler(context *gin.Context) {
	session := sessions.Default(context)
	person := context.PostForm("person")
	penaltyStr := context.PostForm("penalty")

	// Get the penalty from the map using the penalty string
	penalty, ok := penalties.PenaltyMap[penaltyStr]
	if !ok {
		log.Printf("Invalid penalty: %v", penaltyStr)
		return
	}

	// Get the current penalty for the person from the session
	currentPenaltyStr := session.Get("penalty_" + person)
	if currentPenaltyStr != nil {
		// Convert the current penalty to an integer and add the new penalty to it
		currentPenalty, err := strconv.Atoi(currentPenaltyStr.(string))
		if err != nil {
			log.Printf("Error converting current penalty to integer: %v", err)
			return
		}
		penalty += currentPenalty
	}

	// Store the updated penalty in the session
	session.Set("penalty_"+person, strconv.Itoa(penalty))
	log.Printf("Person: %v", person)
	log.Printf("Penalty string: %v", penaltyStr)
	log.Printf("Penalty from map: %v", penalty)
	log.Printf("Current penalty string: %v", currentPenaltyStr)
	session.Save()

	PenaltiesHandler(context)
}
