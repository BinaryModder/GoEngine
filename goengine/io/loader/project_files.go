package loader

import (
	"goengine/core/filesystem"
	"goengine/project"
	"path/filepath"
)

// Loading all assets files
func LoadProjectFiles(path string) ([]project.ProjectFile, string, error) {

	assetsPath := filepath.Join(
		path,
		"Assets",
	)

	files, _, err := filesystem.LoadFolder(assetsPath)
	if err != nil {
		return []project.ProjectFile{}, "", err
	}

	return files, assetsPath, nil
}
