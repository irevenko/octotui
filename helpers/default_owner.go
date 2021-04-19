package helpers

import (
	"fmt"
	"log"
)

const (
	ownerFilename = "default_owner"
)

// OwnerType is the type of owner. Either "user" or "org".
type OwnerType string

const (
	// Org signifies that the owner type is a GitHub organization.
	Org OwnerType = "org"
	// User signifies that the owner type is a GitHub user.
	User OwnerType = "user"
)

// LoadOwner loads the config file for the default owner.
func LoadOwner() string {
	owner, filepath, err := loadConfigFile(ownerFilename)

	if err == errCreatedConfigFile {
		fmt.Printf("Created owner file in: %v\n", filepath)
		fmt.Println("Put your owner in this file in the format \"name:type\"")
		fmt.Printf("Where type is either %q or %q\n", Org, User)
	} else if err != nil {
		log.Fatalf("Unable to load/create owner file: %v", err)
	}
	return owner
}
