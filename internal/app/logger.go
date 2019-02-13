package app

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func InitLogger() *logrus.Logger {
	log.Out = os.Stdout
	environment := "dev"
	switch environment {
	case "dev":
		log.Formatter = &logrus.TextFormatter{}
		log.Level = logrus.InfoLevel
	default:
		log.Formatter = &logrus.JSONFormatter{}
		log.Level = logrus.WarnLevel
	}
	return log
}
