package github

import (
	"context"
	"log"

	"github.com/google/go-github/github"
)

func SearchUser(ctx context.Context, restClient *github.Client, username string) *github.UsersSearchResult {
	opts := &github.SearchOptions{ListOptions: github.ListOptions{PerPage: 50}}

	users, _, err := restClient.Search.Users(ctx, username, opts)
	if err != nil {
		log.Fatal(err)
	}

	return users
}