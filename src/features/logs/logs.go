package logs

import (
	"fmt"
	"log"
)

var logs []string

func LogPerson(person string) {
	logStr := fmt.Sprintf("Person: %v", person)
	logs = append(logs, logStr)
	log.Println(logStr)
}

func LogPenaltyString(penaltyStr string) {
	logStr := fmt.Sprintf("Penalty string: %v", penaltyStr)
	logs = append(logs, logStr)
	log.Println(logStr)
}

func LogPenaltyFromMap(penalty int) {
	logStr := fmt.Sprintf("Penalty from map: %v", penalty)
	logs = append(logs, logStr)
	log.Println(logStr)
}

func LogCurrentPenaltyString(currentPenaltyStr interface{}) {
	logStr := fmt.Sprintf("Current penalty string: %v", currentPenaltyStr)
	logs = append(logs, logStr)
	log.Println(logStr)
}

func GetLogs() []string {
	return logs
}
