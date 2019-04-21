package log

type Option func(*logger)

// Debug turns debug mode on
func Debug(l *logger) {
	l.debug = true
}
