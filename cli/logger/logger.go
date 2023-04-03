package logger

import (
	"errors"
	"io"
	"io/fs"
	"log"
	"os"
)

var file *os.File
var logger Logger
var openFile func(string, int, fs.FileMode) (*os.File, error)

type Logger struct {
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
}

func init() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalln("Failed to open current working directory.")
	}

	openFile = os.OpenFile

	file, err := findLogsFile(cwd)
	if err != nil {
		log.Fatalln(err.Error())
	}

	l := Logger{
		Info:    log.New(io.MultiWriter(file, os.Stdout), "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		Warning: log.New(io.MultiWriter(file, os.Stdout), "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile),
		Error:   log.New(io.MultiWriter(file, os.Stderr), "ERROR: ", log.Ldate|log.Ltime|log.Llongfile),
	}

	logger = l
}

func GetLogger() Logger {
	return logger
}

func findLogsFile(currentWorkingDir string) (*os.File, error) {
	var err error

	file, err = openFile(currentWorkingDir+"\\logs.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, errors.New("Failed to open/create the '.log' file: ")
	}

	return file, nil
}
