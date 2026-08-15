package logger

import (
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Entry struct {
	Level string
	Text  string
}

var Log *zap.Logger
var Entries []Entry
var entriesMu sync.RWMutex
var logFile *os.File

func Init() error {
	encoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	file, err := os.OpenFile("goengine.log", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	logFile = file
	Log = zap.New(zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stderr), zapcore.AddSync(file)),
		zap.DebugLevel,
	))
	return nil
}

func Sync() {
	if Log != nil {
		_ = Log.Sync()
	}
	if logFile != nil {
		_ = logFile.Close()
		logFile = nil
	}
}

func Info(msg string) {
	Log.Info(msg)

	entriesMu.Lock()
	defer entriesMu.Unlock()
	Entries = append(Entries,
		Entry{
			Level: "INFO",
			Text:  msg,
		})
}

func Warning(msg string) {
	Log.Warn(msg)

	entriesMu.Lock()
	defer entriesMu.Unlock()
	Entries = append(Entries,
		Entry{
			Level: "WARN",
			Text:  msg,
		})
}
func Error(msg string) {
	Log.Error(msg)

	entriesMu.Lock()
	defer entriesMu.Unlock()
	Entries = append(Entries,
		Entry{
			Level: "ERROR",
			Text:  msg,
		})
}
func Message(msg string) {
	Log.Info(msg)

	entriesMu.Lock()
	defer entriesMu.Unlock()
	Entries = append(Entries,
		Entry{
			Level: "MESSAGE",
			Text:  msg,
		})
}

func GetEntries() []Entry {
	entriesMu.RLock()
	defer entriesMu.RUnlock()

	entries := make([]Entry, len(Entries))
	copy(entries, Entries)
	return entries
}
