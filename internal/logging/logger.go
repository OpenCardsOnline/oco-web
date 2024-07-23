package logger

import (
	"fmt"
	"time"
)

var Log *Logger

type Logger struct {
}

func InitializeLogger() {
	Log = &Logger{}
}

func (logger *Logger) Info(message string) {
	fmt.Printf("%s <INFO> - %s \n", time.Now().Format(time.DateTime), message)
}

func (logger *Logger) Error(message string, error error) {
	fmt.Printf("%s <ERROR> - %s \n", time.Now().Format(time.DateTime), message)
	fmt.Println(error.Error())
}
