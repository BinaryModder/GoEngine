package scale

import (
	"github.com/AllenDang/cimgui-go/imgui"
)

const (
	ScalingMac   = 0.75
	ScalingOther = 1.3
	UIScale      = 0.75
)

var CurrentScaling float32

func X(el float32) float32 {
	return el * UIScale
}

func Y(el float32) float32 {
	return el * UIScale
}

func I(el int) int {
	return int(float32(el) * UIScale)
}

func SetFontScale() {
	fbScale := imgui.CurrentIO().DisplayFramebufferScale()

	switch {
	case fbScale.X > 1:
		CurrentScaling = ScalingMac

		imgui.CurrentIO().SetFontGlobalScale(CurrentScaling)
	default:
		CurrentScaling = ScalingOther
		imgui.CurrentIO().SetFontGlobalScale(ScalingOther)
	}

}
