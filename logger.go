package logger

import (
	"encoding/json"
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

type LoggerError uint

const (
	ErrorNone LoggerError = iota
	ErrorInvalidArgument
	ErrorLogFileOpenFailed
	ErrorLogFileCloseFailed
	ErrorLevelNotEnough
	ErrorJsonConvertFailed
)

type LoggerOutputType uint

const (
	OutputTypeNormal LoggerOutputType = iota
	OutputTypeJson
)

type LoggerJson struct {
	Level   string `json:"level"`
	Message string `json:"message"`
	Date    string `json:"date"`
}

type Logger struct {
	level   LoggerLevel
	logtype LoggerOutputType
	path    string
	fp      *os.File
}

func (logger *Logger) Init(level LoggerLevel, path string, logtype LoggerOutputType) LoggerError {
	if level > LevelError {
		return ErrorInvalidArgument
	}

	if logtype > OutputTypeJson {
		return ErrorInvalidArgument
	}

	// default path
	if path == "" {
		path = "./info.log"
	}

	var err error
	logger.fp, err = os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return ErrorLogFileOpenFailed
	}

	logger.level = level
	logger.path = path
	logger.logtype = logtype
	return ErrorNone
}

func (logger *Logger) Deinit() LoggerError {
	if err := logger.fp.Close(); err != nil {
		//fmt.Printf("close error. reason[%v]\n", err)
		return ErrorLogFileCloseFailed
	}

	return ErrorNone
}

func getDateStr() string {
	now := time.Now()
	str := fmt.Sprintf(
		"%04d/%02d/%02d %02d:%02d:%02d",
		now.Year(),
		now.Month(),
		now.Day(),
		now.Hour(),
		now.Minute(),
		now.Second())

	return str
}

func getLevelStr(level LoggerLevel) string {
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

func (logger *Logger) print(level LoggerLevel, str string) LoggerError {
	if level < logger.level {
		return ErrorLevelNotEnough
	}

	err := ErrorNone
	switch logger.logtype {
	case OutputTypeNormal:
		err = logger.printNormal(level, str)
	case OutputTypeJson:
		err = logger.printJson(level, str)
	}

	return err
}

func (logger *Logger) printNormal(level LoggerLevel, str string) LoggerError {
	lvstr := getLevelStr(level)
	now := getDateStr()

	fmt.Fprintf(logger.fp, "[%s] [%s] %s\n", now, lvstr, str)
	return ErrorNone
}

func (logger *Logger) printJson(level LoggerLevel, str string) LoggerError {
	lvstr := getLevelStr(level)
	now := getDateStr()

	var tmpstruct LoggerJson
	tmpstruct.Level = lvstr
	tmpstruct.Date = now
	tmpstruct.Message = str
	json, err := json.Marshal(tmpstruct)
	if err != nil {
		return ErrorJsonConvertFailed
	}

	fmt.Fprintf(logger.fp, "%+v\n", string(json))
	return ErrorNone
}

func (logger *Logger) Debug(str string) LoggerError {
	return logger.print(LevelDebug, str)
}

func (logger *Logger) Info(str string) LoggerError {
	return logger.print(LevelInfo, str)
}

func (logger *Logger) Warning(str string) LoggerError {
	return logger.print(LevelWarning, str)
}

func (logger *Logger) Error(str string) LoggerError {
	return logger.print(LevelError, str)
}
