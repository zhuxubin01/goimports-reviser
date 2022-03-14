package helper

import (
	"os"

	"github.com/zhuxubin01/bbimports/pkg/module"
	"github.com/zhuxubin01/bbimports/reviser"
)

func DetermineProjectName(projectName, filePath string) (string, error) {
	if filePath == reviser.StandardInput {
		var err error
		filePath, err = os.Getwd()
		if err != nil {
			return "", err
		}
	}

	return module.DetermineProjectName(projectName, filePath)
}
