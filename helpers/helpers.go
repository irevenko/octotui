package helpers

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

var configPath = path.Join(".config", "octotui")

var errCreatedConfigFile error = errors.New("File created")

func loadConfigFile(filename string) (contents string, createdPath string, err error) {
	home, err := os.UserHomeDir()

	if err != nil {
		err = fmt.Errorf("Unable to get home directory: %w", err)
		return
	}

	fullConfigPath := path.Join(home, configPath)
	fullFilename := path.Join(fullConfigPath, filename)

	if _, innerErr := os.Stat(fullFilename); innerErr != nil {
		if os.IsNotExist(innerErr) {
			innerErr := os.Mkdir(fullConfigPath, 0755)
			if innerErr != nil && !os.IsExist(innerErr) {
				err = fmt.Errorf("Unable to create octotui folder in %v", fullConfigPath)
				return
			}

			blackListFile, innerErr := os.OpenFile(fullFilename, os.O_RDONLY|os.O_CREATE, 0644)
			if innerErr != nil {
				err = fmt.Errorf("Unable to create file %v: %w", fullFilename, innerErr)
				return
			}
			blackListFile.Close()

			createdPath = fullFilename
			err = errCreatedConfigFile
			return
		}
		err = innerErr
		return
	}

	fileContents, err := ioutil.ReadFile(fullFilename)
	contents = strings.TrimSpace(string(fileContents))
	return
}
