package main

import (
	"fmt"
	"os"

	"github.com/bhoriuchi/go-bunyan/bunyan"
)

var LoggerConfig bunyan.Config
var Logger bunyan.Logger

func registerLogger() {
	name := "spot-interruption-service"
	level := bunyan.LogLevelInfo
	if os.Getenv("SERVICE_NAME") != "" {
		name = os.Getenv("SERVICE_NAME")
	}
	if os.Getenv("LOG_LEVEL") != "" {
		name = os.Getenv("LOG_LEVEL")
	}
	LoggerConfig = bunyan.Config{
		Name:   name,
		Level:  level,
		Stream: os.Stdout,
	}
	log, err := bunyan.CreateLogger(LoggerConfig)
	if err != nil {
		fmt.Println("Error creating logger")
		panic(err)
	}
	log.Info("Logger registered")
	Logger = log
}
