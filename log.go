package gopkg

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true})
	if lv, err := logrus.ParseLevel(os.Getenv("LOG_LEVEL")); err == nil {
		logrus.SetLevel(lv)
	}
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
