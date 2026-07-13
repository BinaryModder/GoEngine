package editor

import (
	"time"
)

type ProjectConfig struct {
	Name string `json:"name"`

	Version string `json:"version"`

	EngineVersion string `json:"engineVersion"`

	CreatedAt time.Time `json:"createdAt"`
}
