package logging

import (
	"fmt"
	"irisStudy/conf"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Level int

var (
	logPrefix  = ""
	levelFlags = []string{"DEBUG", "INFO", "WARNING", "ERROR", "FATAL"}
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

type Log struct {
	Config conf.Config
	File   *os.File
	Logger *log.Logger
}

func NewLog(config conf.Config) *Log {
	filePath := getLogFileFullPath(config.LogSavePath, config.LogSaveName, config.TimeFormat, config.LogFileExt)
	F := openLogFile(filePath, config.LogSavePath)
	l := log.New(F, config.DefaultPrefix, log.LstdFlags)
	return &Log{
		File: F,
		Logger: l,
	}
}

func (log *Log) Debug(v ...interface{}) {
	log.setPrefix(DEBUG)
	log.Logger.Println(v)
}

func (log *Log) Info(v ...interface{}) {
	log.setPrefix(INFO)
	log.Logger.Println(v)
}

func (log *Log) WARNING(v ...interface{}) {
	log.setPrefix(WARNING)
	log.Logger.Println(v)
}

func (log *Log) ERROR(v ...interface{}) {
	log.setPrefix(ERROR)
	log.Logger.Println(v)
}

func (log *Log) FATAL(v ...interface{}) {
	log.setPrefix(FATAL)
	log.Logger.Println(v)
}

func (log *Log) setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(log.Config.DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}
	log.Logger.SetPrefix(logPrefix)
}

