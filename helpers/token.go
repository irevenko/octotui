package helpers

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const (
	tokenPath = "/.config/octotui/token"
)

func LoadToken() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	if _, err := os.Stat(home + tokenPath); err != nil {
		if os.IsNotExist(err) {
			err := os.Mkdir(home+"/.config/octotui", 0755)
			if err != nil {
				log.Fatal("Unable to create octotui folder in " + home + "/.config")
			}

			blackListFile, err := os.OpenFile(home+tokenPath, os.O_RDONLY|os.O_CREATE, 0644)
			if err != nil {
				log.Fatal("Unable to create token file in " + home + tokenPath)
			}
			blackListFile.Close()

			fmt.Println("Created token file in: " + home + tokenPath)
			fmt.Println("Put your github token in this file")
		}
	}

	token, err := ioutil.ReadFile(home + tokenPath)
	if err != nil {
		log.Fatal("Can't read token file in: " + home + tokenPath)
	}

	return strings.TrimSpace(string(token))
}
