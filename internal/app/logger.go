package app

import (
	"fmt"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()
var environment = "dev"

// InitLogger is a function for initializing the logger for main.go file.
func InitLogger() *logrus.Logger {
	log.Out = os.Stdout
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

// InitLoggerEndpoint is a function for initializing the logger for endpoints.
func InitLoggerEndpoint(r *http.Request) *logrus.Entry {
	log.Out = os.Stdout
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
	logger := log.WithFields(logrus.Fields{
		"request_id":  r.Header.Get("request_id"),
		"http_method": r.Method,
		"request_uri": fmt.Sprintf("%s", r.RequestURI),
		"user_agent":  r.UserAgent(),
	})
	return logger
}
