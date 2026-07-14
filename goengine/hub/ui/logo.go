package ui

import (
	"github.com/AllenDang/giu"
)

func Logo() giu.Widget {
	return giu.Image(
		Icon,
	).Size(170, 170)
}
