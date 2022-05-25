package bumblogger

import (
	"fmt"
	"os"
	"strings"
	"time"
)

var logLevel = "info"

func SetLogLevel(lvl string) error {
	switch strings.ToLower(lvl) {
	case "info":
		logLevel = "info"
		return nil
	case "debug":
		logLevel = "debug"
		return nil
	case "error", "err":
		logLevel = "error"
		return nil
	case "warning", "warn":
		logLevel = "warning"
		return nil
	}

	Fatal("Not valid log level [", lvl, "] please, set the correct level. "+
		"\nAvalible levels: \n 'info' \n 'debug' \n 'warning \n 'error'")

	return fmt.Errorf("Not valid log level %d", lvl)
}

func getTime() string {
	fullTime := time.Now().Format("01-02-2006 15:04:05")
	return fullTime
}

func DebugMsg(message ...interface{}) {
	if logLevel == "debug" {
		currentTime := getTime()
		fmt.Printf("\n [%s] \033[32m[DEBUG]\u001B[0m lvl= %s  msg= %s", currentTime, logLevel, message)
	}
}

func Info(message ...interface{}) {
	if logLevel == "info" || logLevel == "debug" {
		currentTime := getTime()
		fmt.Printf("\n [%s] \u001B[36m[INFO]\u001B[0m lvl= %s  msg= %s", currentTime, logLevel, message)
	}
}

func Error(message ...interface{}) {
	currentTime := getTime()
	fmt.Printf("\n [%s] \033[31m[ERROR]\u001B[0m lvl= %s  msg= %s", currentTime, logLevel, message)
}

func Warning(message ...interface{}) {
	if logLevel != "error" {
		currentTime := getTime()
		fmt.Printf("\n [%s] \033[33m[WARNING]\u001B[0m lvl= %s  msg= %s", currentTime, logLevel, message)
	}
}

func Fatal(message ...interface{}) {
	currentTime := getTime()
	fmt.Printf("\n [%s] \033[31m[FATAL]\u001B[0m lvl= %s  msg= %s", currentTime, logLevel, message)
	os.Exit(1)
}
