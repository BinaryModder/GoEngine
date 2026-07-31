package saver

import (
	"encoding/json"
	"fmt"
	"goengine/scene"
	"io"
	"os"
	"path/filepath"
)

func copyTextureToProject(sourcePath, projectPath string) (string, error) {
	texturesDir := filepath.Join(projectPath, "Assets", "Textures")
	if err := os.MkdirAll(texturesDir, 0755); err != nil {
		return "", err
	}

	baseName := filepath.Base(sourcePath)
	destPath := filepath.Join(texturesDir, baseName)

	srcFile, err := os.Open(sourcePath)
	if err != nil {
		return "", err
	}
	defer srcFile.Close()

	destFile, err := os.Create(destPath)
	if err != nil {
		return "", err
	}
	defer destFile.Close()

	if _, err := io.Copy(destFile, srcFile); err != nil {
		return "", err
	}

	return destPath, nil
}

func WriteMaterialFile(name, albedo, projectPath string) (*scene.Material, error) {

	if name == "" {
		name = "default"
	}

	resolvedAlbedo := "null"
	if albedo != "" && albedo != "null" {
		copiedPath, err := copyTextureToProject(albedo, projectPath)
		if err != nil {
			return nil, fmt.Errorf("failed to copy texture: %v", err)
		}
		resolvedAlbedo = copiedPath
	}

	materialFolder := filepath.Join(
		projectPath,
		"Assets",
		"Materials",
	)

	var material scene.Material
	material = scene.Material{
		Name:   name,
		Albedo: resolvedAlbedo,
		Color:  [3]float32{1.0, 1.0, 1.0},
	}

	fileData, err := json.MarshalIndent(material, "", "    ")

	if err != nil {
		return nil, err
	}
	filePath := filepath.Join(
		materialFolder,
		fmt.Sprintf("%s.material", name),
	)

	return &material, os.WriteFile(
		filePath,
		fileData, 0644)

}
