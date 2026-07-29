package saver

import (
	"encoding/json"
	"fmt"
	"goengine/scene"
	"os"
	"path/filepath"
)

func WriteMaterialFile(path, name, albedo string) error {

	if name == "" {
		name = "default"
	}
	if albedo == "" {
		albedo = "null"
	}
	materialFolder := filepath.Join(
		path,
		"Assets",
		"Materials",
	)

	var material scene.Material
	material = scene.Material{
		Name:   name,
		Albedo: albedo,
		Color:  [3]float32{1.0, 1.0, 1.0},
	}

	fileData, err := json.MarshalIndent(material, "", "    ")

	if err != nil {
		return err
	}
	filePath := filepath.Join(
		materialFolder,
		fmt.Sprintf("%s.material", name),
	)

	return os.WriteFile(
		filePath,
		fileData, 0644)

}
