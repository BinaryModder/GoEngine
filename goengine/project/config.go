package project

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

type ProjectFile struct {
	Name        string
	Path        string
	IsDir       bool
	AmountFiles int
}
