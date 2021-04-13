package github

import (
	"context"
	"log"

	"github.com/google/go-github/v33/github"
)

func SearchUser(ctx context.Context, restClient *github.Client, username string) *github.UsersSearchResult {
	opts := &github.SearchOptions{ListOptions: github.ListOptions{PerPage: 50}}

	users, _, err := restClient.Search.Users(ctx, username, opts)
	if err != nil {
		log.Fatalf("Error while searching for: %v: %v", username, err)
	}

	return users
}
