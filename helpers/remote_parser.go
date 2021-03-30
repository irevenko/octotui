package helpers

import (
	"errors"
	"regexp"
)

var githubRegex = regexp.MustCompile(
	`^(?:(?:git@github\.com:)|(?:https?://github\.com/))([\w\.\-]+)/(?:[\w\.\-]+)\.git$`,
)

// ErrOwnerNotFound signifies no owner could be found from the remote URL
var ErrOwnerNotFound error = errors.New("No match found")

// OwnerFromRemote gets the GitHub repo owner from a remote URL
func OwnerFromRemote(remote string) (owner string, err error) {
	matches := githubRegex.FindStringSubmatch(remote)
	if len(matches) > 0 {
		owner = matches[1]
	} else {
		err = ErrOwnerNotFound
	}
	return
}
