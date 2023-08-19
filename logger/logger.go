package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

type Logger struct {
    LogDirectory  string
    LogFilePath   *os.File 
    InfoLogger    *log.Logger
    WarnLogger    *log.Logger
    ErrorLogger   *log.Logger
    FatalLogger   *log.Logger
}

var LOG  = &Logger{
    LogDirectory: "application_logs",
}

func init() {
   
    err := os.Mkdir(LOG.LogDirectory, 0777)
    if err != nil {
        fmt.Println("Logger initialization failed." + err.Error())
    }
    year, month, day := time.Now().Date()
    fileName := fmt.Sprintf("%v-%v-%v.log", day, month.String(), year)
    filePath, _ := os.OpenFile(LOG.LogDirectory+"/"+fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    LOG.LogFilePath = filePath
    LOG.InfoLogger = log.New(filePath, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
    LOG.WarnLogger = log.New(filePath, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
    LOG.ErrorLogger = log.New(filePath, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
    LOG.FatalLogger = log.New(filePath, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func (l *Logger) Info(message string)  {
    l.InfoLogger.Println(message)
}

func (l *Logger) Warning(message string)  {
    l.WarnLogger.Println(message)
}

func (l *Logger) Error(message string)  {
    l.ErrorLogger.Println(message)
}

func (l *Logger) Fatal(message string)  {
    l.FatalLogger.Println(message)
}