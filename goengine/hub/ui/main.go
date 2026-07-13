package ui

import (
	"github.com/AllenDang/giu"
)

func Loop() {

	giu.SingleWindow().
		Layout(

			giu.Row(

				Sidebar(),
				giu.Label("Projects"),
				MainPanel(),
			),
		)

}
