package logger

import (
	"go.uber.org/zap"
)

type Entry struct {
	Level string
	Text  string
}

var Log *zap.Logger
var Entries []Entry

func Init() error {

	var err error

	Log, err = zap.NewDevelopment()

	if err != nil {
		return err
	}

	return nil

}

func Sync() {
	if Log != nil {
		_ = Log.Sync()
	}
}

func Info(msg string) {
	Log.Info(msg)

	Entries = append(Entries,
		Entry{
			Level: "INFO",
			Text:  msg,
		})
}

func Warning(msg string) {
	Log.Warn(msg)

	Entries = append(Entries,
		Entry{
			Level: "WARN",
			Text:  msg,
		})
}
func Error(msg string) {
	Log.Error(msg)

	Entries = append(Entries,
		Entry{
			Level: "ERROR",
			Text:  msg,
		})
}
