package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Ygg-Drasill/Thing/src/penalties"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func PenaltiesHandler(context *gin.Context) {
	session := sessions.Default(context)

	penalty := penalties.PenaltyMap

	data := struct {
		Person    []string
		Penalty   map[string]int
		Penalties map[string]int
	}{
		Person:    []string{"Androkles", "Alexander", "Mathias", "Gustav", "Tobias", "Mikael"},
		Penalty:   penalty,
		Penalties: make(map[string]int),
	}

	for _, person := range data.Person {
		penalty := session.Get("penalty_" + person)
		if penalty != nil {
			if str, ok := penalty.(string); ok {
				if value, err := strconv.Atoi(str); err == nil {
					data.Penalties[person] = value
				} else {
					fmt.Println("Error converting string to int:", err)
				}
			}
		} else {
			data.Penalties[person] = 0
		}
	}

	context.HTML(http.StatusOK, "penalties.html", data)
}
