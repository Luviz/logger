package csvtools

import (
	"fmt"
	"os"
	"path/filepath"
)

type CSVError struct {
	path, message string
}

func (err *CSVError) Error() string {
	return fmt.Sprintf("%v @ %v", err.message, err.path)
}

func FileExist(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func CreateCSVLogFile(path string) error {
	os.MkdirAll(filepath.Dir(path), 0755)
	file, err := os.Create(path)
	if err != nil {
		return &CSVError{
			path:    path,
			message: err.Error(),
		}
	}
	addCSVHeaderToLogFile(file)
	file.Close()
	return err
}

func AppendToFile(path, message string) error {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	file.Write([]byte(message + "\n"))
	file.Close()
	return err
}

func addCSVHeaderToLogFile(f *os.File) {
	f.Write([]byte(getCSVLogHeader()))
}

func getCSVLogHeader() string {
	return "timestamp;level;value;message;source\n"
}
