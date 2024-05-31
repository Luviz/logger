package logger

import (
	"fmt"
	"time"
)

type StdOutLogHandler struct {
	info string
}

func NewStdOutLogHandler(info string) *StdOutLogHandler {
	return &StdOutLogHandler{
		info: info,
	}
}

func (s StdOutLogHandler) LogLine(log Log) string {
	return fmt.Sprintf(
		"[%v] - %v: %v",
		log.timestamp.Format(time.RFC3339),
		log.logLevel,
		log.message,
	)
}

func (s StdOutLogHandler) Save(log Log) {
	fmt.Println(s.LogLine(log))
}

func (s StdOutLogHandler) Info() string {
	return fmt.Sprintf("Stdout: %v", s.info)
}
