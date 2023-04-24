package logger

import (
	"log"
	"os"
	"runtime"
	"time"
)

var (
	errorLogger, infoLogger, debugLogger *log.Logger
)

func InitLogger() {
	errorLogger = log.New(os.Stderr, "ERROR : "+time.Now().Format("2006-01-02 3:04:05 PM")+": ", 0)
	infoLogger = log.New(os.Stdout, "INFO : "+time.Now().Format("2006-01-02 3:04:05 PM")+": ", 0)
	debugLogger = log.New(os.Stdout, "DEBUG : "+time.Now().Format("2006-01-02 3:04:05 PM")+": ", 0)
}

func Info(v ...interface{}) {
	infoLogger.Println(v...)
}

func Error(v ...interface{}) {
	errorLogger.Println(v...)
}

func Debug(v ...interface{}) {
	debugLogger.Println(v...)
}

func FuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}
