package logger

import (
	"fmt"
	"os"
	"time"
)

type LoggerLevel uint

const (
	LevelDebug LoggerLevel = iota
	LevelInfo
	LevelWarning
	LevelError
)

type LoggerError int

const (
	ErrorNone LoggerError = iota
	ErrorInvalidArgument
	ErrorLogFileOpenFailed
)

type Logger struct {
	level LoggerLevel
	path  string
	fp    *os.File
}

func (logger *Logger) Init(level LoggerLevel, path string) LoggerError {
	if level > LevelError {
		fmt.Printf("Invalid argument. level[%d]\n", level)
		return ErrorInvalidArgument
	}

	// default path
	if path == "" {
		path = "./info.log"
	}

	var err error
	logger.fp, err = os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("open error. reason[%v]\n", err)
		return ErrorLogFileOpenFailed
	}

	logger.level = level
	logger.path = path
	return ErrorNone
}

func (logger *Logger) Deinit() {
	if err := logger.fp.Close(); err != nil {
		fmt.Printf("close error. reason[%v]\n", err)
	}
}

func GetDateStr() string {
	now := time.Now()
	str := fmt.Sprintf(
		"%04d/%02d/%02d(%s) %02d:%02d:%02d",
		now.Year(),
		now.Month(),
		now.Day(),
		now.Weekday(),
		now.Hour(),
		now.Minute(),
		now.Second())

	return str
}

func GetLevelStr(level LoggerLevel) string {
	lvstr := "???"
	switch level {
	case LevelDebug:
		lvstr = "Debug"
	case LevelInfo:
		lvstr = "Info"
	case LevelWarning:
		lvstr = "Warning"
	case LevelError:
		lvstr = "Error"
	}

	return lvstr
}

func (logger *Logger) Print(level LoggerLevel, str string) {
	if level < logger.level {
		return
	}

	lvstr := GetLevelStr(level)
	now := GetDateStr()
	fmt.Fprintf(logger.fp, "[%s] [%s] %s\n", now, lvstr, str)
}

func (logger *Logger) Debug(str string) {
	logger.Print(LevelDebug, str)
}

func (logger *Logger) Info(str string) {
	logger.Print(LevelInfo, str)
}

func (logger *Logger) Warning(str string) {
	logger.Print(LevelWarning, str)
}

func (logger *Logger) Error(str string) {
	logger.Print(LevelError, str)
}
