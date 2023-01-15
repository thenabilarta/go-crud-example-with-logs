package config

import (
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

func LoadConfig(logFile *os.File) {
	t := time.Now()

	sDate := fmt.Sprintf("%d%02d%02d", t.Year(), t.Month(), t.Day())

	logFileName := "log/" + sDate + ".log"

	logFile, err := os.OpenFile(logFileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Failed to create logfile" + logFileName)
		panic(err)
	}
	// Output to stdout instead of the default stderr

	log.SetOutput(logFile)

	// defer f.Close()
}
