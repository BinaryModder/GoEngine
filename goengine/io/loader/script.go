package loader

import (
	"encoding/json"
	"goengine/scene"
	"os"
	"path/filepath"
)

// Unpacking one .material file

func LoadScriptFile(path string) (*scene.Material, error) {

	data, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	var result_material scene.Material

	err = json.Unmarshal(
		data,
		&result_material,
	)

	if err != nil {
		return nil, err
	}

	return &scene.Material{
		Index:  0,
		Name:   result_material.Name,
		Albedo: result_material.Albedo,
		Color:  result_material.Color,
	}, nil

}

// Loading all materials to the current session Editor State
func LoadProjectScripts(path string) ([]scene.Material, error) {

	materialsPath := filepath.Join(
		path,
		"Assets",
		"Materials",
	)

	var result_materials []scene.Material

	files, err := os.ReadDir(materialsPath)

	if err != nil {
		return nil, err
	}

	var file_path string
	var index int
	for _, file := range files {
		if file.IsDir() || filepath.Ext(file.Name()) != ".material" {
			continue
		}
		file_path = filepath.Join(
			materialsPath,
			file.Name(),
		)
		material, err := LoadMaterialFile(file_path)
		if err != nil {
			return nil, err
		}
		material.Index = index
		index++
		result_materials = append(result_materials, *material)
	}

	return result_materials, nil
}
