package app

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

const env = "dev"

// InitLogger is a function for initializing a logger for main.go file.
func InitLogger() *logrus.Logger {
	switch env {
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

// InitLoggerEndpoint is a function for initializing a logger for endpoints.
func InitLoggerEndpoint(r *http.Request) *logrus.Entry {
	switch env {
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
