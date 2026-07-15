package hub

import (
	"time"
)

type Project struct {
	Name       string
	Path       string
	CreatedAt  time.Time
	LastOpened time.Time
}

type ProjectConfig struct {
	Name string `json:"name"`

	Version string `json:"version"`

	EngineVersion string `json:"engineVersion"`

	CreatedAt time.Time `json:"createdAt"`
}

type Scene struct {
	Name    string        `json:"name"`
	Objects []SceneObject `json:"objects"`
}

type SceneObject struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
