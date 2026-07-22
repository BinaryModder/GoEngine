package editor_ui

import (
	s "goengine/ui/scale"
)

const (
	//interface components
	HierarchyWidth = 250
	InspectorWidth = 270
	ViewportWidth  = 959
	ViewportHeight = 650
	ProjectHeight  = 233

	//widgets
	parameterInputSize  = 60
	parameterSliderSize = 80

	//Buttons

	//save and edit buttons (menu bar)
	saveeditSizeWeight = 63.5
	saveeditSizeHeight = 0

	//run project button (menu bar)
	runSizeWeight = 63.5
	runSizeHeight = 0

	//folder and file button

	folbutWeight = 125
	folbutHeight = 120
)

var (
	toMiddleDummyWeight = float32(0)
	toMiddleDummyHeight = float32(0)
)

func ConfigureSize() {
	switch {
	case s.CurrentScaling == s.ScalingOther:

		toMiddleDummyWeight = 450.5
		toMiddleDummyHeight = 0

	default:
		toMiddleDummyWeight = 215
		toMiddleDummyHeight = 0
	}
}
