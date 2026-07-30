package scene

import (
	"errors"
	"strings"
)

func ValidateMaterialNamePath(NewCreateName, NewCreatePath string) error {
	if strings.TrimSpace(NewCreateName) == "" {
		return errors.New("Enter project name")
	}
	if strings.TrimSpace(NewCreatePath) == "" {
		return errors.New("Enter project path")
	}
	return nil
}
