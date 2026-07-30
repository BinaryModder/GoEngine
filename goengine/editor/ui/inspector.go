package ui

import (
	"fmt"
	"github.com/AllenDang/giu"
	"goengine/editor"
	"goengine/editor/functions"
	"goengine/scene"
)

var (
	//Object Script
	selectedScriptIndex int32
	selectedScriptName  string
	scriptsNameList     []string

	//Object Name
	selectedObjectName string

	//Object Material
	selectedMaterialIndex int32
	selectedMaterial      *scene.Material
	materialNamesList     []string
)

func Inspector() giu.Widget {
	widgets := []giu.Widget{
		giu.Label("Inspector"),
		giu.Separator(),
	}

	if editor.State.CurrentScene == nil || editor.State.SelectedObject == "" {
		widgets = append(widgets, giu.Label("No object selected"))
		return giu.Child().Size(
			InspectorWidth, -ProjectHeight,
		).Layout(widgets...)
	}
	//Object selections
	var obj *scene.SceneObject
	for i := range editor.State.CurrentScene.Objects {
		if editor.State.CurrentScene.Objects[i].Name == editor.State.SelectedObject {
			obj = &editor.State.CurrentScene.Objects[i]
			break
		}
	}

	if obj == nil {
		widgets = append(widgets, giu.Label("Object not found"))
		return giu.Child().Size(InspectorWidth, -ProjectHeight).Layout(widgets...)
	}

	//Name object
	selectedObjectName = obj.Name
	//End of name object

	// Script inspector initializing
	if obj.Script != "" {
		selectedScriptName = obj.Script
	} else {
		selectedScriptName = "No script"
	}

	scriptsNameList = functions.LoadScriptsNames(editor.State.ProjectPath) // Loading script list for choosing

	if len(scriptsNameList) <= 1 {
		scriptsNameList = []string{"No scripts found"}
	}
	// End of script inspector initializing

	//Material inspector initializing

	if obj.Material != nil {
		selectedMaterial = obj.Material
	} else {
		selectedMaterial = &scene.Material{
			Name: "No material",
		}
	}

	materialNamesList = functions.LoadMaterialsNames(&editor.State.Materials)

	// End of material inspector initializing

	widgets = append(widgets,
		giu.Row(
			giu.Label("Name: "),
			giu.InputText(&selectedObjectName).Size(parameterInputNameObjectSize).Flags(
				giu.InputTextFlagsEnterReturnsTrue,
			).
				OnChange(func() {

					contains := editor.State.CurrentScene.HasObject(selectedObjectName)

					if contains {
						return
					}
					obj.Name = selectedObjectName
					editor.State.SelectedObject = selectedObjectName
				}),
		),
		giu.Label(fmt.Sprintf("Type: %s", obj.Type)),
		giu.Separator(),
	)

	widgets = append(widgets,
		giu.Label("Transform"),

		giu.Label("Position"),
		giu.Row(
			giu.InputFloat(&obj.Transform.Position[0]).Label("X##pos").Size(parameterInputSize),
			giu.InputFloat(&obj.Transform.Position[1]).Label("Y##pos").Size(parameterInputSize),
			giu.InputFloat(&obj.Transform.Position[2]).Label("Z##pos").Size(parameterInputSize),
		),

		giu.Label("Rotation"),
		giu.Row(
			giu.InputFloat(&obj.Transform.Rotation[0]).Label("X##rot").Size(parameterInputSize),
			giu.InputFloat(&obj.Transform.Rotation[1]).Label("Y##rot").Size(parameterInputSize),
			giu.InputFloat(&obj.Transform.Rotation[2]).Label("Z##rot").Size(parameterInputSize),
		),

		giu.Label("Scale"),
		giu.Row(
			giu.InputFloat(&obj.Transform.Scale[0]).Label("X##scl").Size(parameterInputSize),
			giu.InputFloat(&obj.Transform.Scale[1]).Label("Y##scl").Size(parameterInputSize),
			giu.InputFloat(&obj.Transform.Scale[2]).Label("Z##scl").Size(parameterInputSize),
		),

		giu.Separator(),

		giu.Label("Script"),
		giu.Row(
			giu.Combo("", selectedScriptName, scriptsNameList, &selectedScriptIndex).Size(200).
				OnChange(func() {
					// here could be func where we should give index of selected object
					obj.Script = scriptsNameList[selectedScriptIndex]
					selectedScriptName = scriptsNameList[selectedScriptIndex]
				}),
		),

		giu.Separator(),

		giu.Label("Material"),
		giu.Row(
			giu.Combo("", selectedMaterial.Name, materialNamesList, &selectedMaterialIndex).Size(200).
				OnChange(func() {
					// here could be func where we should give index of selected object
					obj.Material = &editor.State.Materials[selectedMaterialIndex]
				}),
		),
	)

	if len(obj.Parameters) > 0 {
		widgets = append(widgets, giu.Label("Parameters"))

		for key, val := range obj.Parameters {
			switch v := val.(type) {
			case string:
				widgets = append(widgets, giu.Label(fmt.Sprintf("%s: %s", key, v)))

			case float64:
				val32 := float32(v)
				sliderID := fmt.Sprintf("%s##param_%s", key, key)
				widgets = append(widgets,
					giu.SliderFloat(&val32, 0.1, 200.0).Label(sliderID).OnChange(func() {
						obj.Parameters[key] = float64(val32)
					}),
				)

			case bool:
				bVal := v
				checkboxID := fmt.Sprintf("%s##param_%s", key, key)
				widgets = append(widgets,
					giu.Checkbox(checkboxID, &bVal).OnChange(func() {
						obj.Parameters[key] = bVal
					}),
				)

			case []interface{}:
				if len(v) == 3 {
					var vec [3]float32
					valid := true

					for j := 0; j < 3; j++ {
						if num, ok := v[j].(float64); ok {
							vec[j] = float32(num)
						} else {
							valid = false
						}
					}

					if valid {
						widgets = append(widgets, giu.Label(key+" (RGB/XYZ)"))
						widgets = append(widgets, giu.Row(
							giu.SliderFloat(&vec[0], 0.0, 1.0).Label("X/R##v0_"+key).Size(parameterSliderSize).OnChange(func() {
								obj.Parameters[key] = []interface{}{float64(vec[0]), float64(vec[1]), float64(vec[2])}
							}),
							giu.SliderFloat(&vec[1], 0.0, 1.0).Label("Y/G##v1_"+key).Size(parameterSliderSize).OnChange(func() {
								obj.Parameters[key] = []interface{}{float64(vec[0]), float64(vec[1]), float64(vec[2])}
							}),
						))
						widgets = append(widgets,
							giu.Row(
								giu.SliderFloat(&vec[2], 0.0, 1.0).Label("Z/B##v2_"+key).Size(parameterSliderSize).OnChange(func() {
									obj.Parameters[key] = []interface{}{float64(vec[0]), float64(vec[1]), float64(vec[2])}
								}),
							))
					}
				}
			}
		}
	}

	return giu.Child().
		Size(InspectorWidth, -ProjectHeight).
		Layout(widgets...)
}
