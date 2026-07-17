package hub_ui

import (
	"github.com/AllenDang/giu"
	"log"
)

func Loop() {

	if !isAssetsLoaded {
		if err := LoadTextures(); err != nil {
			log.Fatal("Failed to load hub textures")
		}
	}
	giu.SingleWindow().
		Layout(

			giu.Row(
				Sidebar(),
				MainPanel(),
			),
		)

}
