package screenlog

import (
	"log"
	"os"
	"time"

	"github.com/fareez-ahamed/opt-build/logger"
)

// ScreenLog is the filelogging struct
type ScreenLog struct {
	File *os.File
}

// Info is for printing the message
func (f *ScreenLog) Info(message string) {
	log.Printf(logger.Format, time.Now(), message)
}

// NewFileLog creates a new log
func NewFileLog(filename string) *ScreenLog {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		panic(err)
	}
	return &ScreenLog{file}
}

func init() {
	logger.RegisterLog(NewFileLog("log.txt"))
}
