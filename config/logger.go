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
	logger := log.New(writer, prefix, log.Ldate|log.Ldate|log.Ltime|log.Lshortfile)

	return &Logger{
		debug:   log.New(writer, "DEBUG: ", logger.Flags()),
		info:    log.New(writer, "INFO: ", logger.Flags()),
		warning: log.New(writer, "WARN: ", logger.Flags()),
		err:     log.New(writer, "ERROR: ", logger.Flags()),
		writer:  writer,
	}
}

// helpers
func (l *Logger) Debug(v ...interface{}) {
	l.debug.Println(v...)
}
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.debug.Printf(format, v...)
}
func (l *Logger) Info(v ...interface{}) {
	l.debug.Println(v...)
}
func (l *Logger) Infof(format string, v ...interface{}) {
	l.debug.Printf(format, v...)
}
func (l *Logger) Warning(v ...interface{}) {
	l.debug.Println(v...)
}
func (l *Logger) Warningf(format string, v ...interface{}) {
	l.debug.Printf(format, v...)
}
func (l *Logger) Error(v ...interface{}) {
	l.debug.Println(v...)
}
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.debug.Printf(format, v...)
}
