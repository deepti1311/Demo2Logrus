package main

import (
	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	"log"
	"os"
)

type Logger struct {
	adapter Adapter
}

func (l *Logger) SetLogger(a Adapter) {
	l.adapter = a
}

func (l *Logger) Debugf(fmt string, args ...interface{}) {
	l.adapter.Debugf(fmt, args...)
}

func (l *Logger) Debug(fmt string, args ...interface{}) {

	//creates a log file if it does not exists

	log := logrus.New()

	file := "logrus.txt"
	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		//fmt.Println("failed create log file ", file)
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

	l.adapter.Debug(fmt, args...)

}

type Adapter interface {
	Debugf(string, ...interface{})
	Debug(string, ...interface{})
}

type StdLoggerAdapter struct{}

func (l StdLoggerAdapter) Debugf(fmt string, args ...interface{}) {
	log.Printf(fmt, args...)
}
func (l StdLoggerAdapter) Debug(fmt string, args ...interface{}) {
	log.Printf(fmt, args...)
}

type LogrusAdapter struct {
}

func (l LogrusAdapter) Debugf(fmt string, args ...interface{}) {
	logrus.Debugf(fmt, args...)

}

func (l LogrusAdapter) Debug(fmt string, args ...interface{}) {
	logrus.Debug(fmt, args)

}

func NewLogger(a Adapter) Logger {

	return Logger{adapter: a}

}
