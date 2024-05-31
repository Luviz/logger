package logger_test

import (
	"testing"

	"github.com/luviz/logger"
)

func TestLoggerSTDOUT(t *testing.T) {
	log := logger.NewLogger(map[string]string{
		"env": "Test",
	})

	log.AddLogHandler(logger.NewStdOutLogHandler("Just Testing"))
	lls := log.ListLogHandlers()
	if len(lls) != 1 {
		t.Error("too many Handlers")
	}
	log.LogEvent(logger.Trace, "Hello world", "test")
	log.LogEvent(logger.Debug, "Hello world", "test")
	log.LogEvent(logger.Info, "Hello world", "test")
	log.LogEvent(logger.Waring, "Hello world", "test")
	log.LogEvent(logger.Error, "Hello world", "test")
}
