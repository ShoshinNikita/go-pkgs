package log

type Option func(*logger)

// Debug turns debug mode on
var Debug = func() Option {
	return func(l *logger) {
		l.debug = true
	}
}()
