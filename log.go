package logger

import "time"

type Log struct {
	timestamp time.Time
	isMetric  bool
	logLevel  LogLevel
	source    string
	message   string
	value     float64
	tag       map[string]string
}
