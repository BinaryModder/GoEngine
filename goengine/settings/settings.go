package settings

type GoEngineSettings struct {
	Login   string `json:"login"`
	Theme   string `json:"theme"`
	Console bool   `json:"console"`
}

var State GoEngineSettings
