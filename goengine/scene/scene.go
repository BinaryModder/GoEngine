package scene

type GameObject struct {
	Name string

	Type string

	Transform Transform
}

type Transform struct {
	Position [3]float32

	Rotation [3]float32

	Scale [3]float32
}

type Scene struct {
	Name    string        `json:"name"`
	Objects []SceneObject `json:"objects"`
}

type SceneObject struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
