package hub

import (
	"fmt"
	"goengine/engine/logger"
	"goengine/engine/platform"
	"goengine/io/loader"
	"goengine/io/saver"
	"goengine/project"
	"goengine/settings"
)

type Page int

const (
	PageProjects Page = iota //0
	PageSettings             //1
)

type HubState struct {
	ShowCreateProject bool

	CurrentPage Page // Current Page on Interface (for switching)

	Projects []project.Project // List of loaded project (for project cards)

	NewCreateName string // Creating project part. Just storing name of new project (from InputText)

	NewCreatePath string // Creating project part. Just storing  of new project(from InputText)

	SaveSettingsShowButton bool // State for "Save" button

	//Settings
	ShowConsole bool
}

func (s *HubState) Init() error {

	// Loading settings
	err := loader.LoadSettings()

	if err != nil {
		if err.Error() == "Settings file does not exists" {
			if err = saver.CreateSettings(); err != nil {
				logger.Warning("Failed to create settings.json")
			}
		}

	}

	State.ShowConsole = settings.State.Console

	logger.Info("Settings are loaded")

	//Loading Platform Information

	platform.Init()

	logger.Info(fmt.Sprintf("Platform is initialized: %s", platform.State.OS))

	return nil
}
