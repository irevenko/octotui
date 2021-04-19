package helpers

import (
	"fmt"
	"log"
)

const (
	tokenFilename = "token"
)

func LoadToken() string {
	token, filepath, err := loadConfigFile(tokenFilename)

	if err == errCreatedConfigFile {
		fmt.Printf("Created token file in: %v\n", filepath)
		fmt.Println("Put your github token in this file")
	} else if err != nil {
		log.Fatalf("Unable to load/create token file: %v", err)
	}
	return token
}
