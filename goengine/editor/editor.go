package editor

import ()

func LoadWholeProject() error {
	if err := LoadScene(); err != nil {
		return err
	}

	if err := LoadProjectFile(); err != nil {
		return err
	}

	return nil

}
