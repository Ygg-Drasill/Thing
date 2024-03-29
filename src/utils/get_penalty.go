// utils/penalty.go
package utils

import (
	"log"
	"strconv"

	"github.com/gin-contrib/sessions"
)

func GetPenalty(session sessions.Session, person string) int {
	penalty := session.Get("penalty_" + person)
	if penalty != nil {
		if penaltyStr, ok := penalty.(string); ok {
			if penaltyStr == "No Penalty" {
				return 0
			}
			if value, err := strconv.Atoi(penaltyStr); err == nil {
				return value
			} else {
				log.Printf("Error converting string to int: %v", err)
			}
		}
	}
	return 0
}
