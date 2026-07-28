package console

import (
	"fmt"
	"time"
)

type ConsoleState struct {
	Logs       []LogMessage
	AutoScroll bool
}

var State ConsoleState

func (s *ConsoleState) Info(msg string) {
	s.Logs = append(s.Logs,
		LogMessage{
			Time:    time.Now(),
			Type:    Info,
			Message: fmt.Sprintf("[INFO]: %s", msg),
		})
}

func (s *ConsoleState) Warning(msg string) {
	s.Logs = append(s.Logs,
		LogMessage{
			Time:    time.Now(),
			Type:    Warning,
			Message: fmt.Sprintf("[WARNING]: %s", msg),
		})
}
func (s *ConsoleState) Error(msg string) {
	s.Logs = append(s.Logs,
		LogMessage{
			Time:    time.Now(),
			Type:    Error,
			Message: fmt.Sprintf("[ERROR]: %s", msg),
		})
}
