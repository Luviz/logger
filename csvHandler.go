package logger

import (
	"fmt"

	csvTools "github.com/luviz/logger/csvTools"
)

type CSVHandler struct {
	info     string
	filepath string
}

func NewCSVHandler(info, filepath string) *CSVHandler {
	// Create log file
	if !csvTools.FileExist(filepath) {
		csvTools.CreateCSVLogFile(filepath)
	}

	return &CSVHandler{
		info:     info,
		filepath: filepath,
	}
}

func (s CSVHandler) Save(log Log) {
	csvTools.AppendToFile(s.filepath, log.String(";"))
}

func (s CSVHandler) Info() string {
	return fmt.Sprintf("Stdout: %v", s.info)
}
