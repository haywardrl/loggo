package loggo

import (
	"fmt"
	"io"
	"os"
)

// Logger is used to log information.
type Logger struct {
	threshold Level
	output    io.Writer
}

// New returns you a logger, ready to log at the desired threshold
// Give it a list of configuration functions to tune it at your will.
// The default output is Stdout.
func New(threshold Level) *Logger {
	return &Logger{
		threshold: threshold,
		output:    os.Stdout,
	}
}

func (l *Logger) logf(level Level, format string, args ...any) {
	message := fmt.Sprintf(format, args...)
	_, _ = fmt.Fprintf(l.output, "%s %s", level, message)
}

// Debugf formats and prints a message if the log level is debug or higher.
func (l *Logger) Debugf(format string, args ...any) {
	if l.threshold > LevelDebug {
		return
	}
	l.logf(LevelDebug, format, args...)
}

// Infof formats and prints a message if the log level is info or higher.
func (l *Logger) Infof(format string, args ...any) {
	if l.threshold > LevelInfo {
		return
	}
	l.logf(LevelInfo, format, args...)
}

// Warnf formats and prints a message if the log level is warn or higher.
func (l *Logger) Warnf(format string, args ...any) {
	if l.threshold > LevelWarn {
		return
	}
	l.logf(LevelWarn, format, args...)
}

// Errorf formats and prints a message if the log level is error or higher.
func (l *Logger) Errorf(format string, args ...any) {
	if l.threshold > LevelError {
		return
	}
	l.logf(LevelError, format, args...)
}

// Fatalf formats and prints a message if the log level is fatal.
func (l *Logger) Fatalf(format string, args ...any) {
	if l.threshold > LevelFatal {
		return
	}
	l.logf(LevelFatal, format, args...)
}
