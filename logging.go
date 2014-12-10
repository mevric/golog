package golog

import ()

type loggerI interface {
	SetLevel()
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
}
