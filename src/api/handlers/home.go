package handlers

import (
	"net/http"

	"github.com/Ygg-Drasill/Thing/src/features/penalties"
	"github.com/Ygg-Drasill/Thing/src/features/people"
	"github.com/Ygg-Drasill/Thing/src/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func HomePage(context *gin.Context) {
	session := sessions.Default(context)
	selectedPerson := session.Get("selectedPerson")

	var selectedPersonStr string
	if selectedPerson != nil {
		selectedPersonStr = selectedPerson.(string)
	} else {
		selectedPersonStr = ""
	}

	person := people.PersonList
	penalty := penalties.PenaltyMap

	data := struct {
		Title          string
		Header         string
		Items          []string
		Person         []string
		SelectedPerson string
		Penalty        map[string]int
		Penalties      map[string]int
	}{
		Title:          "Thing",
		Header:         "Welcome to Thing!",
		Items:          []string{"You", "Will", "Pay"},
		Person:         person,
		SelectedPerson: selectedPersonStr,
		Penalty:        penalty,
		Penalties:      make(map[string]int),
	}

	for _, person := range data.Person {
		data.Penalties[person] = utils.GetPenalty(session, person)
	}

	context.HTML(http.StatusOK, "home.html", data)
}
