package handlers

import (
	"net/http"

	"github.com/Ygg-Drasill/Thing/src/features/penalties"
	"github.com/Ygg-Drasill/Thing/src/features/people"
	"github.com/Ygg-Drasill/Thing/src/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func PenaltiesHandler(context *gin.Context) {
	session := sessions.Default(context)

	person := people.PersonList
	penalty := penalties.PenaltyMap

	data := struct {
		Person    []string
		Penalty   map[string]int
		Penalties map[string]int
	}{
		Person:    person,
		Penalty:   penalty,
		Penalties: make(map[string]int),
	}

	for _, person := range data.Person {
		data.Penalties[person] = utils.GetPenalty(session, person)
	}

	context.HTML(http.StatusOK, "penalties.html", data)
}
