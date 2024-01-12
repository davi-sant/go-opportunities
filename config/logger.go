package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

func NewLogger(prefix string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, prefix, log.Ldate|log.Ltime)
	return &Logger{
		debug:   log.New(writer, " _DEBUG_ ", logger.Flags()),
		info:    log.New(writer, " _INFO_ ", logger.Flags()),
		warning: log.New(writer, " _WARNING_ ", logger.Flags()),
		err:     log.New(writer, " _ERROR_ ", logger.Flags()),
		writer:  writer,
	}

}

// Non-formatted Logs
func (logger *Logger) Debug(v ...interface{}) {
	logger.debug.Println(v...)
}

func (logger *Logger) Info(v ...interface{}) {
	logger.info.Println(v...)
}

func (logger *Logger) Warning(v ...interface{}) {
	logger.warning.Println(v...)
}

func (logger *Logger) Error(v ...interface{}) {
	logger.err.Println(v...)
}

// Formatted Logs
func (logger *Logger) Debugf(format string, v ...interface{}) {
	logger.debug.Printf(format, v...)
}

func (logger *Logger) Infof(format string, v ...interface{}) {
	logger.info.Printf(format, v...)
}

func (logger *Logger) Warningf(format string, v ...interface{}) {
	logger.warning.Printf(format, v...)
}

func (logger *Logger) Errorf(format string, v ...interface{}) {
	logger.err.Printf(format, v...)
}
