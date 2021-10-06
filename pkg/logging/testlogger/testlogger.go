package testlogger

import (
	"testing"

	"github.com/go-logr/logr"
)

// TestLogger is a logr.Logger that prints through a testing.T object.
// Only error logs will have any effect.
type TestLogger struct {
	T *testing.T
}

var _ logr.Logger = TestLogger{}

func (tl TestLogger) Info(msg string, args ...interface{}) {
	tl.T.Logf("%s: %v", msg, args)
}

func (TestLogger) Enabled() bool {
	return true
}

func (tl TestLogger) Error(err error, msg string, args ...interface{}) {
	tl.T.Logf("%s: %v -- %v", msg, err, args)
}

func (tl TestLogger) V(v int) logr.Logger {
	return tl
}

func (tl TestLogger) WithName(_ string) logr.Logger {
	return tl
}

func (tl TestLogger) WithValues(_ ...interface{}) logr.Logger {
	return tl
}

// This will be useful for k8s 1.23+
// func New(t *testing.T) logr.Logger {
// 	return logr.New(TestSink{T: t})
// }

// // TestSink is a logr.LogSink that prints through a testing.T object.
// // Only error logs will have any effect.
// type TestSink struct {
// 	T *testing.T
// }

// var _ logr.LogSink = TestSink{}

// func (tl TestSink) Init(info logr.RuntimeInfo) {}

// func (TestSink) Enabled(level int) bool {
// 	return true
// }

// func (tl TestSink) Error(err error, msg string, args ...interface{}) {
// 	tl.T.Logf("%s: %v -- %v", msg, err, args)
// }

// func (tl TestSink) Info(level int, msg string, args ...interface{}) {
// 	tl.T.Logf("%s: %v", msg, args)
// }

// func (tl TestSink) WithName(_ string) logr.LogSink {
// 	return tl
// }

// func (tl TestSink) WithValues(_ ...interface{}) logr.LogSink {
// 	return tl
// }
