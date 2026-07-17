package functions

import (
	"os"
	"path/filepath"

	"goengine/project"
)

func LoadFolder(path string) ([]project.ProjectFile, string, error) {

	files, err := os.ReadDir(path)
	if err != nil {
		return nil, "", err
	}

	projectFiles := make([]project.ProjectFile, 0, len(files))

	for _, file := range files {

		projectFiles = append(projectFiles,

			project.ProjectFile{
				Name:  file.Name(),
				Path:  filepath.Join(path, file.Name()),
				IsDir: file.IsDir(),
			},
		)
	}

	return projectFiles, path, nil
}
func AbsolutePath(path string) string {

	abs_path, err := filepath.Abs(path)

	if err != nil {
		return ""
	}

	return abs_path

}
