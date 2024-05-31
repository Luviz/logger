package csvtools_test

import (
	"testing"

	csvTools "github.com/luviz/logger/csvTools"
)

var logFilePath string = "./tmp/me.log"

func TestCreateFiles(t *testing.T) {
	err := csvTools.CreateCSVLogFile(logFilePath)
	if err != nil {
		t.Error(err)
	}
}

func TestAppend(t *testing.T) {
	err := csvTools.AppendToFile(logFilePath, "hello world\n")
	if err != nil {
		t.Error(err)
	}
}

func TestFileExist(t *testing.T) {
	isFile := csvTools.FileExist(logFilePath)
	if !isFile {
		t.Fatal("This file should exist")
	}
	isNotFile := csvTools.FileExist("not-a-file")
	if isNotFile {
		t.Fatal("This file should not exist")
	}
}
