package hub_ui

import (
	"fmt"
	"github.com/AllenDang/giu"
	"goengine/engine/platform"
	"goengine/settings"
	"goengine/ui/scale"
	"log"
)

// Some flags for Initializing
var (
	isSettingsReady          bool
	isSettingsFailed         bool
	isFontScalingInitialized bool
	isPlatformInitialized    bool
)

// The centre of HUB Interface
func Loop() {

	//Loading Settings
	if !isSettingsReady && !isSettingsFailed {

		err := settings.LoadSettings()

		if err != nil {
			if err.Error() == "Settings file does not exists" {
				if err = settings.CreateSettings(); err != nil {
					isSettingsFailed = true
					fmt.Println("Failed to create settings.json")
				}
			}

		}

		isSettingsReady = true

	}

	//Loading Assets
	if !isAssetsLoaded {
		if err := LoadTextures(); err != nil {
			log.Fatal("Failed to load hub textures")
		}

		isAssetsLoaded = true

	}

	//Loading Fonts
	if !isFontScalingInitialized {
		scale.SetFontScale()

		isFontScalingInitialized = true
	}

	//Loading Platform Information
	if !isPlatformInitialized {
		platform.Init()
		isPlatformInitialized = true
	}

	//Connecting all widgets
	giu.SingleWindow().
		Layout(

			giu.Row(
				Sidebar(),
				MainPanel(),
			),
		)

}
