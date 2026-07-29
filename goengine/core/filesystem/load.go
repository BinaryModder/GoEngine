package filesystem

import (
	"goengine/project"
	"os"
	"path/filepath"
)

func LoadFolder(path string) ([]project.ProjectFile, string, error) {

	files, err := os.ReadDir(path)
	if err != nil {
		return nil, "", err
	}

	projectFiles := make([]project.ProjectFile, 0, len(files))

	for _, file := range files {

		if file.IsDir() {
			subPath := filepath.Join(
				path,
				file.Name(),
			)
			contains, err := os.ReadDir(subPath)

			if err != nil {
				return nil, "", err
			}

			projectFiles = append(projectFiles,
				project.ProjectFile{
					Name:        file.Name(),
					Path:        filepath.Join(path, file.Name()),
					IsDir:       true,
					AmountFiles: len(contains),
				})
		} else {

			if file.Name() == ".DS_Store" {
				continue
			}
			projectFiles = append(projectFiles,

				project.ProjectFile{
					Name:        file.Name(),
					Path:        filepath.Join(path, file.Name()),
					IsDir:       false,
					AmountFiles: 0,
				},
			)
		}
	}

	return projectFiles, path, nil
}
