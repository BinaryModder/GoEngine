package functions

import (
	"goengine/scene"
	"os"
	"path/filepath"
	"strings"
)

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
	scriptNames := make([]string, 0, len(files))
	scriptNames = append(scriptNames, "No script")

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if !(filepath.Ext(file.Name()) == ".go") {
			continue
		}

		scriptNames = append(scriptNames,
			strings.TrimSuffix(file.Name(), ".go"),
		)

	}

	return scriptNames

}

func LoadMaterialsNames(materials *[]scene.Material) []string {
	if len(*materials) == 0 {
		return []string{"No materials"}
	}

	result_names := make([]string, 0, len(*materials))

	for _, cur_material := range *materials {
		result_names = append(result_names,
			cur_material.Name,
		)
	}
	result_names = append(result_names, "No material")

	return result_names

}
