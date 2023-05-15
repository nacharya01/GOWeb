package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

const (
    DirectoryForLogFiles = "application_logs"
)

type LogDir struct {
    LogDirectory string
}

//Create a directory and return the struct containing the directory path.
func New() *LogDir {
    err := os.Mkdir(DirectoryForLogFiles, 0777)
    if err != nil {
        return nil
    }
    return &LogDir{
        LogDirectory: DirectoryForLogFiles,
    }
}

// Setting up the pattern of the log file.
func SetLogFile() *os.File {
    year, month, day := time.Now().Date()
    fileName := fmt.Sprintf("%v-%v-%v.log", day, month.String(), year)
    filePath, _ := os.OpenFile(DirectoryForLogFiles+"/"+fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

    return filePath
}

// Info level log
func (l *LogDir) Info() *log.Logger {
    getFilePath := SetLogFile()
    return log.New(getFilePath, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// Waringin level log
func (l *LogDir) Warning() *log.Logger {
    getFilePath := SetLogFile()
    return log.New(getFilePath, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// Error level log
func (l *LogDir) Error() *log.Logger {
    getFilePath := SetLogFile()
    return log.New(getFilePath, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// Fatal level log
func (l *LogDir) Fatal() *log.Logger {
    getFilePath := SetLogFile()
    return log.New(getFilePath, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile)
}