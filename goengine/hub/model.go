package hub

import (
	"time"
)

type Project struct {
	Name       string
	Path       string
	LastOpened string
}

type ProjectConfig struct {
	Name string `json:"name"`

	Version string `json:"version"`

	EngineVersion string `json:"engineVersion"`

	CreatedAt time.Time `json:"createdAt"`
}
