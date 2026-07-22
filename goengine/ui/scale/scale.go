package scale

import (
	"github.com/AllenDang/cimgui-go/imgui"
)

const UIScale = 0.75
const GlobalFontScale = 0.75

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
	imgui.CurrentIO().SetFontGlobalScale(0.75)

}
