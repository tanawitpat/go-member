package app

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

// InitLogger is a function for initializing the logger.
func InitLogger() *logrus.Logger {
	log.Out = os.Stdout
	environment := "dev"
	switch environment {
	case "dev":
		log.Formatter = &logrus.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
		}
		log.Level = logrus.InfoLevel

	default:
		log.Formatter = &logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		}
		log.Level = logrus.InfoLevel
	}
	return log
}
