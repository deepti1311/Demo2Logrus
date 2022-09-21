package main

import (
	logger "github.com/sirupsen/logrus"
)

func main() {

	//logger := logrus.New()

	logger.SetFormatter(&logger.TextFormatter{
		ForceColors:   true,
		DisableColors: false,
	})

	logger.SetLevel(logger.TraceLevel)

	logger.Trace("Trace level")
	logger.Debug("Debug level")
	logger.Info("Information  level")

	logger.Warn("Warning level")
	logger.Error("Error level")
	logger.Fatal("Trace level")

}
