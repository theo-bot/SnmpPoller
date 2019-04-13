package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"time"
)

const defaultTimestampFormat = time.RFC3339

func SetupLogger() *log.Logger {
	//
	// Setup loggin
	//
	logFile := fmt.Sprintf("%s", viper.Get("Logging.Logfile"))
	logWriter, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic("Failed to open logfile.")
	}
	//
	// Create new logger object
	//
	var logger = log.New()

	//
	// Connect logger to writer
	//
	logger.SetOutput(logWriter)

	//
	// Format output
	//
	Formatter := new(log.TextFormatter)
	Formatter.TimestampFormat = defaultTimestampFormat
	Formatter.FullTimestamp = true
	Formatter.ForceColors = false
	logger.SetFormatter(Formatter)

	//
	// Set the loglevel
	//
	switch viper.Get("Logging.LogLevel") {
	case "Debug", "debug":
		logger.SetLevel(log.DebugLevel)
	case "Info", "info":
		logger.SetLevel(log.InfoLevel)
	case "Warn", "warn":
		logger.SetLevel(log.WarnLevel)
	case "Error", "error":
		logger.SetLevel(log.ErrorLevel)
	case "Fatal", "fatal":
		logger.SetLevel(log.FatalLevel)
	default:
		logger.SetLevel(log.WarnLevel)
	}

	return logger
}
