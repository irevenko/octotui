package main

import (
	"fmt"
	"strconv"

	g "github.com/irevenko/octostats/graphql"
	r "github.com/irevenko/octostats/rest"
	"github.com/shurcooL/githubv4"
)

func data() {
	qlClient := g.AuthGraphQL("f204672235aa1d2b321da46238a8498ad9a91d60")
	ctx, client := r.AuthREST("f204672235aa1d2b321da46238a8498ad9a91d60")

	allRepos := r.AllRepos(ctx, client, "hajimehoshi")

	usedLangs, langsNum := r.LanguagesByRepo(client, allRepos)
	fmt.Println("Languages By Repo")
	for i, v := range usedLangs {
		fmt.Println(v + ": " + strconv.Itoa(langsNum[i]))
	}

	totalStars := r.TotalStars(client, allRepos)
	fmt.Println("Total stars")
	fmt.Println(totalStars)

	totalForks := r.TotalForks(client, allRepos)
	fmt.Println("Total forks")
	fmt.Println(totalForks)

	contribs := g.AllContributions(qlClient, "irevenko", 2020, 2021)
	fmt.Println(contribs)
	var allCommits githubv4.Int
	var allIssues githubv4.Int
	var allPrs githubv4.Int

	fmt.Println("\nAll commits 2020-2021:")
	for _, v := range contribs.CommitContributionsByRepository {
		allCommits += v.Contributions.TotalCount
	}
	fmt.Println(allCommits)

	fmt.Println("\nAll issues 2020-2021:")
	for _, v := range contribs.IssueContributionsByRepository {
		allIssues += v.Contributions.TotalCount
	}
	fmt.Println(allIssues)

	fmt.Println("\nAll prs 2020-2021:")
	for _, v := range contribs.PullRequestContributionsByRepository {
		allPrs += v.Contributions.TotalCount
	}
	fmt.Println(allPrs)
}
