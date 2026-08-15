package build

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	goruntime "runtime"
	"strings"
	"sync"
	"time"

	"goengine/engine/logger"
)

type Options struct {
	ProjectPath string
	OutputPath  string
	Name        string
}

type Result struct {
	ExecutablePath string
}

func Project(options Options) (*Result, error) {
	projectPath, err := filepath.Abs(options.ProjectPath)
	if err != nil {
		return nil, fmt.Errorf("resolve project path: %w", err)
	}
	outputPath, err := filepath.Abs(options.OutputPath)
	if err != nil {
		return nil, fmt.Errorf("resolve output path: %w", err)
	}

	name := strings.TrimSpace(options.Name)
	if name == "" || name == "." || name != filepath.Base(name) {
		return nil, fmt.Errorf("invalid build name %q", options.Name)
	}
	if err := validateProject(projectPath); err != nil {
		return nil, err
	}

	moduleRoot, err := findModuleRoot()
	if err != nil {
		return nil, err
	}
	if err := os.MkdirAll(outputPath, 0755); err != nil {
		return nil, fmt.Errorf("create output directory: %w", err)
	}

	buildPath := filepath.Join(outputPath, name)
	if err := validateBuildPaths(projectPath, outputPath, buildPath); err != nil {
		return nil, err
	}
	if _, err := os.Stat(buildPath); err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("check build destination: %w", err)
	}

	tempPath, err := os.MkdirTemp(outputPath, "goengine-build-*")
	if err != nil {
		return nil, fmt.Errorf("create temporary build directory: %w", err)
	}
	defer os.RemoveAll(tempPath)
	logger.Info("Build output directory: " + buildPath)
	logger.Info("Temporary build directory: " + tempPath)

	executableName := name
	if goruntime.GOOS == "windows" {
		executableName += ".exe"
	}
	executablePath := filepath.Join(tempPath, executableName)

	logger.Info(fmt.Sprintf("Build started: %s (%s/%s)", name, goruntime.GOOS, goruntime.GOARCH))
	logger.Info("Building player executable")
	if err := buildPlayer(moduleRoot, executablePath); err != nil {
		return nil, err
	}

	for _, directory := range []string{"Assets", "ProjectSettings"} {
		logger.Info("Copying " + directory)
		if err := copyDirectory(filepath.Join(projectPath, directory), filepath.Join(tempPath, directory)); err != nil {
			return nil, fmt.Errorf("copy %s: %w", directory, err)
		}
	}

	if err := replaceBuild(tempPath, buildPath); err != nil {
		return nil, fmt.Errorf("finalize build: %w", err)
	}

	result := &Result{ExecutablePath: filepath.Join(buildPath, executableName)}
	logger.Info("Build completed: " + result.ExecutablePath)
	return result, nil
}

func validateBuildPaths(projectPath, outputPath, buildPath string) error {
	if pathContains(buildPath, projectPath) {
		return errors.New("build destination cannot replace the source project")
	}
	for _, sourceDirectory := range []string{
		filepath.Join(projectPath, "Assets"),
		filepath.Join(projectPath, "ProjectSettings"),
	} {
		if pathContains(sourceDirectory, outputPath) {
			return fmt.Errorf("build output cannot be inside %s", sourceDirectory)
		}
	}
	return nil
}

func pathContains(parent, child string) bool {
	relativePath, err := filepath.Rel(parent, child)
	return err == nil && relativePath != ".." && !strings.HasPrefix(relativePath, ".."+string(filepath.Separator))
}

func replaceBuild(tempPath, buildPath string) error {
	if _, err := os.Stat(buildPath); errors.Is(err, os.ErrNotExist) {
		return os.Rename(tempPath, buildPath)
	} else if err != nil {
		return err
	}

	backupPath, err := os.MkdirTemp(filepath.Dir(buildPath), ".goengine-previous-*")
	if err != nil {
		return err
	}
	if err := os.Remove(backupPath); err != nil {
		return err
	}
	if err := os.Rename(buildPath, backupPath); err != nil {
		return err
	}

	logger.Info("Replacing previous build: " + buildPath)
	if err := os.Rename(tempPath, buildPath); err != nil {
		if rollbackErr := os.Rename(backupPath, buildPath); rollbackErr != nil {
			return fmt.Errorf("install build: %w; restore previous build: %v", err, rollbackErr)
		}
		return err
	}
	if err := os.RemoveAll(backupPath); err != nil {
		logger.Warning("Failed to remove previous build: " + err.Error())
	}
	return nil
}

func validateProject(projectPath string) error {
	for _, path := range []string{
		filepath.Join(projectPath, "Assets"),
		filepath.Join(projectPath, "ProjectSettings", "project.json"),
	} {
		if _, err := os.Stat(path); err != nil {
			return fmt.Errorf("invalid project, cannot access %s: %w", path, err)
		}
	}
	return nil
}

func findModuleRoot() (string, error) {
	candidates := make([]string, 0, 3)
	if workingDirectory, err := os.Getwd(); err == nil {
		candidates = append(candidates, workingDirectory)
	}
	if _, sourcePath, _, ok := goruntime.Caller(0); ok {
		if absolutePath, err := filepath.Abs(sourcePath); err == nil {
			candidates = append(candidates, filepath.Dir(absolutePath))
		}
	}
	if executablePath, err := os.Executable(); err == nil {
		candidates = append(candidates, filepath.Dir(executablePath))
	}

	for _, candidate := range candidates {
		for directory := candidate; ; directory = filepath.Dir(directory) {
			if _, err := os.Stat(filepath.Join(directory, "go.mod")); err == nil {
				return directory, nil
			}
			parent := filepath.Dir(directory)
			if parent == directory {
				break
			}
		}
	}

	return "", errors.New("cannot find GoEngine module root; building requires the engine sources and Go SDK")
}

func buildPlayer(moduleRoot, executablePath string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	cmd := exec.CommandContext(
		ctx,
		"go",
		"build",
		"-trimpath",
		"-ldflags",
		"-s -w",
		"-o",
		executablePath,
		"./cmd/player",
	)
	cmd.Dir = moduleRoot
	logger.Info("Build command directory: " + moduleRoot)
	logger.Info("Build command: go build -trimpath -ldflags \"-s -w\" -o " + executablePath + " ./cmd/player")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("capture build stdout: %w", err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return fmt.Errorf("capture build stderr: %w", err)
	}
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("start go build: %w", err)
	}
	logger.Info(fmt.Sprintf("go build started with PID %d", cmd.Process.Pid))

	done := make(chan struct{})
	go logBuildProgress(done)

	var readers sync.WaitGroup
	readers.Add(2)
	go logOutput(stdout, logger.Info, &readers)
	go logOutput(stderr, logger.Error, &readers)

	readers.Wait()
	err = cmd.Wait()
	close(done)
	if ctx.Err() == context.DeadlineExceeded {
		return errors.New("go build timed out after 10 minutes")
	}
	if err != nil {
		return fmt.Errorf("build player executable: %w", err)
	}
	return nil
}

func logBuildProgress(done <-chan struct{}) {
	timer := time.NewTimer(10 * time.Second)
	defer timer.Stop()

	select {
	case <-timer.C:
		logger.Info("go build is still running; the first build may take longer while dependencies are compiled")
	case <-done:
		return
	}
}

func logOutput(reader io.Reader, writeLog func(string), readers *sync.WaitGroup) {
	defer readers.Done()
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		writeLog("go build: " + scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		logger.Error("Failed to read build output: " + err.Error())
	}
}

func copyDirectory(source, destination string) error {
	return filepath.WalkDir(source, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}

		relativePath, err := filepath.Rel(source, path)
		if err != nil {
			return err
		}
		destinationPath := filepath.Join(destination, relativePath)
		info, err := entry.Info()
		if err != nil {
			return err
		}

		if entry.IsDir() {
			return os.MkdirAll(destinationPath, info.Mode().Perm())
		}
		if !entry.Type().IsRegular() {
			return fmt.Errorf("unsupported project file type: %s", path)
		}

		return copyFile(path, destinationPath, info.Mode().Perm())
	})
}

func copyFile(source, destination string, mode fs.FileMode) error {
	input, err := os.Open(source)
	if err != nil {
		return err
	}

	output, err := os.OpenFile(destination, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, mode)
	if err != nil {
		_ = input.Close()
		return err
	}
	_, copyErr := io.Copy(output, input)
	inputCloseErr := input.Close()
	outputCloseErr := output.Close()
	if copyErr != nil {
		return copyErr
	}
	if inputCloseErr != nil {
		return inputCloseErr
	}
	return outputCloseErr
}
