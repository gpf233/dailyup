package api

import (
	"log"
	"os"
)

func init() {
	logFile, err := os.OpenFile("dailyup.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	log.SetPrefix("[夏驰c@RS小组]")
	log.SetOutput(logFile)
}
