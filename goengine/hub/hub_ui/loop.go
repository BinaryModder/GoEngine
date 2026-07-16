package hub_ui

import (
	"github.com/AllenDang/giu"
)

func Loop() {

	if !isAssetsLoaded {
		LoadAssets()
	}
	giu.SingleWindow().
		Layout(

			giu.Row(
				Sidebar(),
				MainPanel(),
			),
		)

}
