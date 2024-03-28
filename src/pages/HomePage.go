package pages

import (
	"net/http"

	"github.com/Ygg-Drasill/Thing/src/penalties"
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
		Person:         []string{"Androkles", "Alexander", "Mathias", "Gustav", "Tobias", "Mikael"},
		SelectedPerson: selectedPersonStr,
		Penalty:        penalty,
		Penalties:      make(map[string]int),
	}

	for _, person := range data.Person {
		penalty := session.Get("penalty_" + person)
		if penalty != nil {
			// Assuming fine is a string that represents a key in the Penalty map
			data.Penalties[person] = data.Penalty[penalty.(string)]
		} else {
			data.Penalties[person] = 0
		}
	}

	context.HTML(http.StatusOK, "index.html", data)
}
