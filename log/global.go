package log

import (
	"log"
	"os"
)

var globalLogger = &logger{
	// Same as std global logger
	stdLog: log.New(os.Stderr, "", log.LstdFlags),
	debug:  false,
}

func Default() {
	globalLogger.Default()
}

func SetOptions(options ...Option) {
	globalLogger.SetOptions(options...)
}

func GetStdLogger() *log.Logger {
	return globalLogger.GetStdLogger()
}

//

func Print(v ...interface{}) {
	globalLogger.Print(v...)
}

func Printf(format string, v ...interface{}) {
	globalLogger.Printf(format, v...)
}

func Println(v ...interface{}) {
	globalLogger.Println(v...)
}

//

func Fatal(v ...interface{}) {
	globalLogger.Fatal(v...)
}

func Fatalf(format string, v ...interface{}) {
	globalLogger.Fatalf(format, v...)
}

func Fatalln(v ...interface{}) {
	globalLogger.Fatalln(v...)
}
