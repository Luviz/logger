package logger

import (
	"fmt"
	"strings"
	"time"
)

type Log struct {
	timestamp time.Time
	isMetric  bool
	logLevel  LogLevel
	source    string
	message   string
	value     float64
	tag       map[string]string
}

func (l Log) String(sep string) string {
	values := []string{
		l.timestamp.Local().UTC().Format(time.RFC3339),
		l.logLevel.String(),
		fmt.Sprintf("%v", l.value),
		l.message,
		l.source,
	}
	return strings.Join(values, sep)
}
