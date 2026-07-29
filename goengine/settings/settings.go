package settings

type GoEngineSettings struct {
	Login string `json:"login"`
	Theme string `json:"theme"`
}

var State GoEngineSettings
