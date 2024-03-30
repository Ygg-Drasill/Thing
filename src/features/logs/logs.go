package logs

import (
	"fmt"
	"log"
	"os"
)

var (
	logs   []string
	file   *os.File
	logger *log.Logger
)

func init() {
	var err error
	file, err = os.OpenFile("src/logs/penalties.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	logger = log.New(file, "", log.LstdFlags)
}

func LogIP(ip string) {
	logStr := fmt.Sprintf("IP: %v", ip)
	logs = append(logs, logStr)
	logger.Println(logStr)
}

func LogPerson(person string) {
	logStr := fmt.Sprintf("Person: %v", person)
	logs = append(logs, logStr)
	logger.Println(logStr)
}

func LogPenaltyString(penaltyStr string) {
	logStr := fmt.Sprintf("Penalty: %v", penaltyStr)
	logs = append(logs, logStr)
	logger.Println(logStr)
}

func LogPenaltyAmount(penaltyAmount int) {
	logStr := fmt.Sprintf("Penalty amount: %v kr.", penaltyAmount)
	logs = append(logs, logStr)
	logger.Println(logStr)
}

func LogOldPenaltyTotal(oldPenaltyTotal int) {
	logStr := fmt.Sprintf("Old penalty total: %v kr.", oldPenaltyTotal)
	logs = append(logs, logStr)
	logger.Println(logStr)
}

func LogNewPenaltyTotal(newPenaltyTotal int) {
	logStr := fmt.Sprintf("New penalty total: %v kr.\n", newPenaltyTotal)
	logs = append(logs, logStr)
	logger.Println(logStr)
}

func GetLogs() []string {
	return logs
}

func Close() {
	file.Close()
}
