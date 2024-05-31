package logger

import (
	"time"
)

type LogHandler interface {
	Save(log Log)
	Info() string
}

type Logger struct {
	baseTags map[string]string
	handler  []LogHandler
}

func NewLogger(tags map[string]string) Logger {
	return Logger{
		baseTags: tags,
		handler:  []LogHandler{},
	}
}

func (logger Logger) LogEvent(logLevel LogLevel, msg string, source string) Log {
	currentLog := Log{
		timestamp: time.Now(),
		isMetric:  false,
		logLevel:  logLevel,
		message:   msg,
		value:     1,
		source:    source,
		tag:       logger.baseTags,
	}

	for _, s := range logger.handler {
		s.Save(currentLog)
	}

	return currentLog
}

func (logger *Logger) AddLogHandler(store LogHandler) {
	logger.handler = append(logger.handler, store)
}

func (logger Logger) ListLogHandlers() []string {
	res := []string{}
	for _, s := range logger.handler {
		res = append(res, s.Info())
	}
	return res
}
