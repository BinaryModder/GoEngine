package editor

import (
	"github.com/AllenDang/giu"
)

func Loop(project string) {

	giu.SingleWindow().
		Layout(

			giu.Label(
				"GoEngine Editor",
			),

			giu.Label(
				project,
			),
		)

}
