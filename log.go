package log

import (
	"fmt"
	"io"
	"log"
	"os"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

type Logger struct {
	*log.Logger
}

func (l *Logger) Info(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	l.Println(White + s + Reset)
}

func (l *Logger) Warn(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	l.Println(Yellow + s + Reset)
}

func (l *Logger) Err(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	l.Println(Red + s + Reset)
}

func New(filename string) *Logger {
	var fd io.Writer
	var err error
	if filename == "" {
		fd = os.Stdout
	} else {
		fd, err = os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0o666)
		if err != nil {
			panic(err)
		}
	}
	logger := log.New(fd, "", log.Ldate|log.Ltime|log.Lshortfile)
	return &Logger{logger}
}
