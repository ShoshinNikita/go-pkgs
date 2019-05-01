// Package log is a wrapper for std log package
//
// Levels
//
// There are 3 level:
//
//  * Debug – [DBG] prefix
//  * Info – [INF] prefix
//  * Error – [ERR] prefix
//
// Messages with Debug level are not printed by default, but user can
// change by passing Debug option to SetOptions() function
//
package log

import (
	"log"
	"strings"
)

// Prefixes
const (
	DebugPrefix = "[DBG]"
	InfoPrefix  = "[INF]"
	ErrorPrefix = "[ERR]"
)

type logger struct {
	// Can be exported for specific management
	stdLog *log.Logger

	debug bool
}

func (l *logger) Default() {
	l.debug = false
}

func (l *logger) SetOptions(optins ...Option) {
	for _, opt := range optins {
		opt(l)
	}
}

func (l *logger) GetStdLogger() *log.Logger {
	return l.stdLog
}

// Print[f|ln] functions

func (l *logger) Print(v ...interface{}) {
	if l.debug && len(v) > 0 {
		if str, ok := v[0].(string); ok {
			if strings.HasPrefix(str, DebugPrefix) {
				return
			}
		}
	}

	l.stdLog.Print(v...)
}

func (l *logger) Printf(format string, v ...interface{}) {
	if !l.debug && strings.HasPrefix(format, DebugPrefix) {
		return
	}

	l.stdLog.Printf(format, v...)
}

func (l *logger) Println(v ...interface{}) {
	if !l.debug && len(v) > 0 {
		if str, ok := v[0].(string); ok {
			if strings.HasPrefix(str, DebugPrefix) {
				return
			}
		}
	}

	l.stdLog.Println(v...)
}

// Fatal[f|ln] functions

func (l *logger) Fatal(v ...interface{}) {
	l.stdLog.Fatal(v...)
}

func (l *logger) Fatalf(format string, v ...interface{}) {
	l.stdLog.Fatalf(format, v...)
}

func (l *logger) Fatalln(v ...interface{}) {
	l.stdLog.Fatalln(v...)
}
