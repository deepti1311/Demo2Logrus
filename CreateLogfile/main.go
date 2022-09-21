package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	"os"
)

func main() {

	//creates a log file if it does not exists

	log := logrus.New()

	file := "log.txt"
	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		fmt.Println("failed create log file ", file)
		panic(err)
	}

	//changed out put format to: <time> <level> <msg> format
	log = &logrus.Logger{
		Out:   os.Stdout,
		Level: logrus.DebugLevel,
		Formatter: &easy.Formatter{
			TimestampFormat: "2006-01-02 15:04:05",
			LogFormat:       " %time%      [  %lvl%    ]         : - %msg%\n",
		},
	}

	defer f.Close()

	// Output to stdout instead of the default stderr
	log.SetOutput(f)

	// Only log the debug severity or above
	log.SetLevel(logrus.DebugLevel)

	log.SetLevel(logrus.TraceLevel)

	log.Trace("Trace level")
	log.Debug("Debug level")
	log.Info("Information  level")

	log.Warn("Warning level")
	log.Error("Error level")
	log.Fatal("Trace level")

}
