package main

func main() {

	logger := NewLogger(StdLoggerAdapter{})
	logger.Debugf("stdlib logger debug msg")
	logger.Debug("stdlib logger debug msg")

	Logger := NewLogger(LogrusAdapter{})
	Logger.Debugf("loggers debug msg")
	Logger.Debug("Debug msg")

}
