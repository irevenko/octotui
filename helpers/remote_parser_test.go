package helpers

import "testing"

// TestOwnerByHttp tests that the proper owner is calculated from an HTTP URL.
func TestOwnerByHttp(t *testing.T) {
	url := "http://github.com/irevenko/octostats.git"
	owner, err := OwnerFromRemote(url)

	if owner != "irevenko" || err != nil {
		t.Fatalf(`OwnerFromRemote(%q) = %q %v, want "irevenko", nil`, url, owner, err)
	}
}

// TestOwnerByHttps tests that the proper owner is calculated from an HTTPS URL.
func TestOwnerByHttps(t *testing.T) {
	url := "https://github.com/irevenko/octostats.git"
	owner, err := OwnerFromRemote(url)

	if owner != "irevenko" || err != nil {
		t.Fatalf(`OwnerFromRemote(%q) = %q %v, want "irevenko", nil`, url, owner, err)
	}
}

// TestOwnerBySsh tests that the proper owner is calculated from an SSH URL.
func TestOwnerBySsh(t *testing.T) {
	url := "git@github.com:irevenko/octostats.git"
	owner, err := OwnerFromRemote(url)

	if owner != "irevenko" || err != nil {
		t.Fatalf(`OwnerFromRemote(%q) = %q %v, want "irevenko", nil`, url, owner, err)
	}
}
