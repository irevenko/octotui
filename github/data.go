package github

import (
	"context"

	"github.com/google/go-github/github"
	g "github.com/irevenko/octostats/graphql"
	r "github.com/irevenko/octostats/rest"
	"github.com/shurcooL/githubv4"
)

func FetchStats(ctx context.Context, restClient *github.Client, qlClient *githubv4.Client, username string) (langs []string, stars int, forks int) {
	allRepos := r.AllRepos(ctx, restClient, username)

	usedLangs, _ := r.LanguagesByRepo(restClient, allRepos)

	totalStars := r.TotalStars(restClient, allRepos)
	totalForks := r.TotalForks(restClient, allRepos)

	contribs := g.AllContributions(qlClient, "irevenko", 2020, 2021)

	var allCommits githubv4.Int
	var allIssues githubv4.Int
	var allPrs githubv4.Int

	for _, v := range contribs.CommitContributionsByRepository {
		allCommits += v.Contributions.TotalCount
	}

	for _, v := range contribs.IssueContributionsByRepository {
		allIssues += v.Contributions.TotalCount
	}

	for _, v := range contribs.PullRequestContributionsByRepository {
		allPrs += v.Contributions.TotalCount
	}

	return usedLangs, totalStars, totalForks
}
