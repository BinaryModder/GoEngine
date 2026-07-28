package functions

import (
	"goengine/project"
	"goengine/scene"
	"os"
	"path/filepath"
	"strings"
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

func LoadScriptsNames(path string) []string {
	scripts_path := filepath.Join(
		path,
		"Assets",
		"Scripts",
	)
	files, err := os.ReadDir(scripts_path)

	if err != nil {
		return []string{}
	}
	scriptFiles := make([]string, 0, len(files))
	scriptFiles = append(scriptFiles, "No script")

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if !(filepath.Ext(file.Name()) == ".go") {
			continue
		}

		scriptFiles = append(scriptFiles,
			strings.TrimSuffix(file.Name(), ".go"),
		)

	}

	return scriptFiles

}
func AbsolutePath(path string) string {

	abs_path, err := filepath.Abs(path)

	if err != nil {
		return ""
	}

	return abs_path

}

func ContainsInArrayOfObjects(array *scene.Scene, word string) bool {
	for _, obj := range array.Objects {
		if obj.Name == word {
			return true
		}
	}

	return false

}
