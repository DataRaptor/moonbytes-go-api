package logger

import (
	log "github.com/sirupsen/logrus"
)

func Init() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
		ForceColors:   true,
	})
	log.SetReportCaller(false)
	log.SetLevel(log.DebugLevel)
}

func Debug(message string) {
	log.Debug(message)
}

func Panic(message string) {
	log.Panic(message)
}

func Info(message string) {
	log.Info(message)
}

func Warning(message string) {
	log.Warning(message)
}

func Error(message string) {
	log.Error(message)
}

func Critical(message string) {
	log.Info(message)
}
