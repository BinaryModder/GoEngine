package console

import (
	"time"
)

type LogType int

const (
	Info LogType = iota
	Warning
	Error
)

type LogMessage struct {
	Time    time.Time
	Type    LogType
	Message string
}
