package scene

type Material struct {
	Index  int        `json:"index"`  // index of material in current editor session
	Name   string     `json:"name"`   // name of material
	Albedo string     `json:"albedo"` // current path of material
	Color  [3]float32 `json:"color"`  // color of material
}
