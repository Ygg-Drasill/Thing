package pages

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func FinesPage(context *gin.Context) {
	session := sessions.Default(context)
	data := struct {
		Person []string
		Fines  map[string]string
	}{
		Person: []string{"Androkles", "Alexander", "Mathias", "Gustav", "Tobias", "Mikael"},
		Fines:  make(map[string]string),
	}

	for _, person := range data.Person {
		fine := session.Get("fine_" + person)
		if fine != nil {
			data.Fines[person] = fine.(string)
		} else {
			data.Fines[person] = "0"
		}
	}

	context.HTML(http.StatusOK, "fines.html", data)
}
