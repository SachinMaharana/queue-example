package logger

import "github.com/sirupsen/logrus"

// Logger ...
type Logger struct {
	*logrus.Logger
}

// NewLogger ...
func NewLogger() *Logger {
	return &Logger{logrus.New()}
}

// Critical ...
func (log *Logger) Critical(args ...interface{}) {
	log.Error(args...)
}

// Criticalf ...
func (log *Logger) Criticalf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

// Notice ...
func (log *Logger) Notice(args ...interface{}) {
	log.Info(args...)
}

// Noticef ...
func (log *Logger) Noticef(format string, args ...interface{}) {
	log.Infof(format, args...)
}
