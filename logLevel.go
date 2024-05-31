package logger

type LogLevel int

const (
	Trace = iota
	Debug
	Info
	Waring
	Error
)

var logLevel = map[LogLevel]string{
	Trace:  "TRACE",
	Debug:  "Debug",
	Info:   "Info",
	Waring: "Waring",
	Error:  "Error",
}

func (ll LogLevel) String() string {
	return logLevel[ll]
}
