package script

import (
	"os"
	"path/filepath"
	"strings"
)

func GetAvailableScripts(basePath string) []string {
	scriptsDir := filepath.Join(basePath, "Assets", "Scripts")
	entries, err := os.ReadDir(scriptsDir)
	if err != nil {
		return nil
	}

	names := make([]string, 0, len(entries))
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		if filepath.Ext(entry.Name()) != ".go" {
			continue
		}
		name := strings.TrimSuffix(entry.Name(), ".go")
		names = append(names, name)
	}
	return names
}
