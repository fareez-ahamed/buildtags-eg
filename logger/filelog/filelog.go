package filelog

import (
	"fmt"
	"os"
	"time"

	"github.com/fareez-ahamed/opt-build/logger"
)

// FileLog is the filelogging struct
type FileLog struct {
	File *os.File
}

// Info is for printing the message
func (f *FileLog) Info(message string) {
	fmt.Fprintf(f.File, logger.Format, time.Now(), message)
}

// NewFileLog creates a new log
func NewFileLog(filename string) *FileLog {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		panic(err)
	}
	return &FileLog{file}
}

func init() {
	logger.RegisterLog(NewFileLog("log.txt"))
}
