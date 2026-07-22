package hub_ui

import (
	"fmt"
	"github.com/AllenDang/giu"
	"goengine/settings"
	"goengine/ui/scale"
	"log"
)

var (
	isSettingsReady          bool
	isSettingsFailed         bool
	isFontScalingInitialized bool
)

func Loop() {
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
	if !isAssetsLoaded {
		if err := LoadTextures(); err != nil {
			log.Fatal("Failed to load hub textures")
		}

		isAssetsLoaded = true

	}
	if !isFontScalingInitialized {
		scale.SetFontScale()

		isFontScalingInitialized = true
	}

	giu.SingleWindow().
		Layout(

			giu.Row(
				Sidebar(),
				MainPanel(),
			),
		)

}
