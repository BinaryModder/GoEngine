package layout

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

	//Inputs
	ParameterInputNameObjectSize = 150
	ParameterInputSize           = 60
	ParameterSliderSize          = 80

	//Buttons

	//save and edit buttons (menu bar)
	SaveeditSizeWeight = 63.5
	SaveeditSizeHeight = 0

	//run project button (menu bar)
	RunSizeWeight = 63.5
	RunSizeHeight = 0
)

var (
	ToMiddleDummyWeight = float32(0)
	ToMiddleDummyHeight = float32(0)
)

func ConfigureSize() {
	switch {
	case s.CurrentScaling == s.ScalingOther:

		ToMiddleDummyWeight = 330.5
		ToMiddleDummyHeight = 0

	default:
		ToMiddleDummyWeight = 230
		ToMiddleDummyHeight = 0
	}
}
