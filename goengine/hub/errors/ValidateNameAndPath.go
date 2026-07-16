package errors

import (
	"errors"
	"goengine/hub"
	"strings"
)

func ValidateNamePath() error {

	if strings.TrimSpace(hub.State.NewCreateName) == "" {
		return errors.New("Enter project name")
	}
	if strings.TrimSpace(hub.State.NewCreatePath) == "" {
		return errors.New("Enter project path")
	}
	return nil

}
