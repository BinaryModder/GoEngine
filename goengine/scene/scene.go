package scene

type Scene struct {
	Name string

	Objects []GameObject
}

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
