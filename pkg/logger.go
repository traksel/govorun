package modules

import (
	"fmt"
	"log"
	"os"
)

func InitLogFile() {
	var logFilePath string = os.Getenv("PWD") + "/info.log"
	var logFile, _ = os.Create(logFilePath)
	fmt.Println(logFilePath)
}

func RaiseErrorLog(message ...interface{}) {
	errLog := log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime)
	errLog.Printf("%v%v", message[0], message[1])
	return
}

func RaiseWarnLog(message ...interface{}) {
	errLog := log.New(os.Stderr, "[WARN] ", log.Ldate|log.Ltime)
	errLog.Printf("%v%v", message[0], message[1])
	return
}

func RaiseInfoLog(message ...interface{}) {
	errLog := log.New(os.Stderr, "[INFO] ", log.Ldate|log.Ltime)
	errLog.Printf("%v%v", message[0], message[1])
	return
}
