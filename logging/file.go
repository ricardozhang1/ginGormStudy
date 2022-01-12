package logging

import (
	"fmt"
	"log"
	"os"
	"time"
)

func getLogFilePath(logSavePath string) string {
	return fmt.Sprintf("%s", logSavePath)
}

func getLogFileFullPath(logSavePath string, logSaveName, timeFormat, logFileExt string) string {
	prefixPath := getLogFilePath(logSavePath)
	suffixPath := fmt.Sprintf("%s%s.%s", logSaveName, time.Now().Format(timeFormat), logFileExt)
	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

func openLogFile(filePath string, logSavePath string) *os.File {
	_, err := os.Stat(filePath)
	switch {
	case os.IsNotExist(err):
		mkDir(logSavePath)
	case os.IsPermission(err):
		log.Fatalf("Permission :%v", err)
	}
	handle, err := os.OpenFile(filePath, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Fail to OpenFile :%v", err)
	}
	return handle
}

func mkDir(logSavePath string) {
	dir, _ := os.Getwd()
	err := os.MkdirAll(dir + "/" + getLogFilePath(logSavePath), os.ModePerm)
	if err != nil {
		panic(err)
	}
}


