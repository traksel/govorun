package modules

import (
	"log"
	"os"
)

func RaiseErrorLog(message ...interface{}) {
	errLog := log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime)
	errLog.Printf("%v%v", message[0], message[1])
	return
}
