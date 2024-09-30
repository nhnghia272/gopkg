package gopkg

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/sirupsen/logrus"
)

func init() {
	if lv, err := logrus.ParseLevel(os.Getenv("LOG_LEVEL")); err == nil {
		logrus.SetLevel(lv)
	}
	SetLogTextFormatter()
}

func SetLogJSONFormatter() {
	logrus.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true})
}

func SetLogTextFormatter() {
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
}

func Debug(msg any) {
	os.Stderr.WriteString(fmt.Sprintf("panic: %v\n%s\n", msg, debug.Stack()))
}

func Fatal(args ...any) {
	logrus.Fatalln(args...)
}

func Error(args ...any) {
	logrus.Errorln(args...)
}

func Warn(args ...any) {
	logrus.Warnln(args...)
}

func Info(args ...any) {
	logrus.Infoln(args...)
}

func Errorf(format string, args ...any) {
	logrus.Errorf(format, args...)
}

func Warnf(format string, args ...any) {
	logrus.Warnf(format, args...)
}

func Infof(format string, args ...any) {
	logrus.Infof(format, args...)
}

type Log struct {
	entry *logrus.Entry
}

func Fields(kv map[string]any) *Log {
	return &Log{logrus.WithFields(logrus.Fields(kv))}
}

func (s Log) Error(args ...any) {
	s.entry.Errorln(args...)
}

func (s Log) Warn(args ...any) {
	s.entry.Warnln(args...)
}

func (s Log) Info(args ...any) {
	s.entry.Infoln(args...)
}

func (s Log) Errorf(format string, args ...any) {
	s.entry.Errorf(format, args...)
}

func (s Log) Warnf(format string, args ...any) {
	s.entry.Warnf(format, args...)
}

func (s Log) Infof(format string, args ...any) {
	s.entry.Infof(format, args...)
}
