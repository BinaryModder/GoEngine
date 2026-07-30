package scene

import (
	"errors"
	"strings"
)

func ValidateMaterialNamePath(NewCreateName, NewCreatePath string) error {
	if strings.TrimSpace(NewCreateName) == "" {
		return errors.New("Enter material name")
	}
	if strings.TrimSpace(NewCreatePath) == "" {
		return errors.New("Enter material albedo source path")
	}
	return nil
}

func ValidateMaterialNameUnique(name string, existing []Material) error {
	for _, m := range existing {
		if m.Name == strings.TrimSpace(name) {
			return errors.New("material with this name already exists")
		}
	}
	return nil
}
