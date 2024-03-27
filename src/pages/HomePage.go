package pages

import (
	"net/http"

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

	fine := session.Get("fine_" + selectedPersonStr)

	var fineStr string
	if fine != nil {
		fineStr = fine.(string)
	} else {
		fineStr = ""
	}

	penalty := map[string]int{
		"§1 Gasbøde":                 2,
		"§2 Dummebøde":               2,
		"§3 Spildebøde":              2,
		"§4 Push to main":            25,
		"§5 Glemme bøde":             5,
		"§6 Fremmøde med tømmermænd": 5,
		"§7 Dårligt menneske bøde":   2,
		"§8 Mangel af kilder":        2,
		"§9 AI":                      20,
		"§10 Bailer bøde":            20,
	}

	data := struct {
		Title          string
		Header         string
		Items          []string
		Person         []string
		SelectedPerson string
		Penalty        map[string]int
		Fine           string
		Fines          map[string]string
	}{
		Title:          "Thing",
		Header:         "Welcome to Thing!",
		Items:          []string{"You", "Will", "Pay"},
		Person:         []string{"Androkles", "Alexander", "Mathias", "Gustav", "Tobias", "Mikael"},
		SelectedPerson: selectedPersonStr,
		Penalty:        penalty,
		Fine:           fineStr,
		Fines:          make(map[string]string),
	}

	for _, person := range data.Person {
		fine := session.Get("fine_" + person)
		if fine != nil {
			data.Fines[person] = fine.(string)
		} else {
			data.Fines[person] = "0"
		}
	}

	context.HTML(http.StatusOK, "index.html", data)
}
