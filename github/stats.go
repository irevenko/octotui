package github

import (
	"context"
	"time"

	"github.com/google/go-github/github"
	g "github.com/irevenko/octostats/graphql"
	r "github.com/irevenko/octostats/rest"
	"github.com/shurcooL/githubv4"
)

func FetchStats(ctx context.Context, restClient *github.Client, qlClient *githubv4.Client, username string, allRepos []*github.Repository) (int, int, int, int, int) {
	totalStars := r.TotalStars(restClient, allRepos)
	totalForks := r.TotalForks(restClient, allRepos)

	year, _, _ := time.Now().Date()
	contribs := g.AllContributions(qlClient, username, year-1, year)

	var totalCommits int
	var totalIssues int
	var totalPrs int

	for _, v := range contribs.CommitContributionsByRepository {
		totalCommits += v.Contributions.TotalCount
	}

	for _, v := range contribs.IssueContributionsByRepository {
		totalIssues += v.Contributions.TotalCount
	}

	for _, v := range contribs.PullRequestContributionsByRepository {
		totalPrs += v.Contributions.TotalCount
	}

	return totalStars, totalForks, totalCommits, totalIssues, totalPrs
}
