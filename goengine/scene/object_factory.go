package scene

func DefaultTransform() Transform {
	return Transform{
		Position: [3]float32{0, 1, 0},
		Rotation: [3]float32{0, 0, 0},
		Scale:    [3]float32{1, 1, 1},
	}
}
func NewCube(name string, color [3]float32) SceneObject {
	return SceneObject{
		Name:      name,
		Type:      "Mesh",
		MeshType:  "Cube",
		Transform: DefaultTransform(),
		Parameters: map[string]any{
			"Color": []interface{}{
				float64(color[0]),
				float64(color[1]),
				float64(color[2]),
			},
		},
	}
}

func NewPyramid(name string, color [3]float32) SceneObject {
	return SceneObject{
		Name:      name,
		Type:      "Mesh",
		MeshType:  "Pyramid",
		Transform: DefaultTransform(),
		Parameters: map[string]any{
			"Color": []interface{}{
				float64(color[0]),
				float64(color[1]),
				float64(color[2]),
			},
		},
	}
}

func NewCamera(name string) SceneObject {
	t := DefaultTransform()
	t.Position = [3]float32{0, 5, 10}
	t.Rotation = [3]float32{-25, 0, 0}

	return SceneObject{
		Name:      name,
		Type:      "Camera",
		Transform: t,
		Parameters: map[string]any{
			"FOV":  float64(60.0),
			"Near": float64(0.1),
			"Far":  float64(1000.0),
		},
	}
}

func NewLight(name string, intensity float32, color [3]float32) SceneObject {
	t := DefaultTransform()
	t.Position = [3]float32{0, 10, 0}

	return SceneObject{
		Name:      name,
		Type:      "Light",
		Transform: t,
		Parameters: map[string]any{
			"Intensity": float64(intensity),
			"Color": []interface{}{
				float64(color[0]),
				float64(color[1]),
				float64(color[2]),
			},
			"CastShadows": true,
		},
	}
}
