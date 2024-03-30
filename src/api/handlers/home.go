package handlers

import (
	"net/http"

	"github.com/Ygg-Drasill/Thing/src/features/penalties"
	"github.com/Ygg-Drasill/Thing/src/features/people"
	"github.com/Ygg-Drasill/Thing/src/types/data"
	"github.com/Ygg-Drasill/Thing/src/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func HomePage(context *gin.Context) {
	session := sessions.Default(context)
	selectedPerson := session.Get("selectedPerson")
	submitted := session.Get("submitted")

	var selectedPersonStr string
	if selectedPerson != nil {
		selectedPersonStr = selectedPerson.(string)
	} else {
		selectedPersonStr = ""
	}

	person := people.PersonList
	penalty := penalties.PenaltyMap

	templateData := data.NewTemplateData(person, selectedPersonStr, penalty)

	for _, person := range templateData.Person {
		templateData.Penalties[person] = utils.GetPenalty(session, person)
	}

	if submitted != nil && submitted.(bool) {
		session.Set("submitted", false)
		session.Save()
		context.HTML(http.StatusOK, "penalties.html", templateData)
	} else {
		context.HTML(http.StatusOK, "home.html", templateData)
	}
}
