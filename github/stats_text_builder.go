package github

import (
	"context"
	"strconv"
	"strings"

	"github.com/google/go-github/github"
	g "github.com/irevenko/octostats/graphql"
	r "github.com/irevenko/octostats/rest"
	"github.com/shurcooL/githubv4"
)

func BuildProfileInfo(user g.User) string {
	var baseProfile string
	joinedAt := strings.Split(user.CreatedAt.String(), " ")

	if user.Name == "" {
		baseProfile = user.Login + "\n" +
			"[Followers:](fg:yellow) " + strconv.Itoa(user.Followers.TotalCount) + "\n" +
			"[Following:](fg:yellow) " + strconv.Itoa(user.Following.TotalCount) + "\n" +
			"[Starred Repos:](fg:yellow) " + strconv.Itoa(user.StarredRepositories.TotalCount) + "\n" +
			"[Joined:](fg:yellow) " + joinedAt[0] + "\n"
	} else {
		baseProfile = user.Name + "\n" +
			"[Followers:](fg:yellow) " + strconv.Itoa(user.Followers.TotalCount) + "\n" +
			"[Following:](fg:yellow) " + strconv.Itoa(user.Following.TotalCount) + "\n" +
			"[Starred Repos:](fg:yellow) " + strconv.Itoa(user.StarredRepositories.TotalCount) + "\n" +
			"[Joined:](fg:yellow) " + joinedAt[0] + "\n"
	}

	if user.Bio != "" {
		baseProfile += "[Bio:](fg:yellow) " + user.Bio + "\n"
	}

	if user.Status.Message != "" {
		baseProfile += "[Status:](fg:yellow) " + user.Status.Message + "\n"
	}

	if user.Location != "" {
		baseProfile += "[Location:](fg:yellow) " + user.Location + "\n"
	}

	if user.Email != "" {
		baseProfile += "[Email:](fg:yellow) " + user.Email + "\n"
	}

	if user.Company != "" {
		baseProfile += "[Company:](fg:yellow) " + user.Company + "\n"
	}

	if user.TwitterUsername != "" {
		baseProfile += "[Twitter:](fg:yellow) @" + user.TwitterUsername + "\n"
	}

	if user.WebsiteURL != "" {
		baseProfile += "[Website:](fg:yellow) " + user.WebsiteURL + "\n"
	}

	return baseProfile
}

func BuildProfileStats(ctx context.Context, restClient *github.Client, qlClient *githubv4.Client, user g.User, allRepos []*github.Repository) string {
	var baseStats string

	s, f, c, i, p := FetchStats(ctx, restClient, qlClient, user.Login, allRepos)
	usedLicenses, _ := r.MostUsedLicenses(restClient, allRepos)

	baseStats = "Profile Statistics" + "\n" +
		"[Total stars:](fg:magenta) " + strconv.Itoa(s) + "\n" +
		"[Total forks:](fg:magenta) " + strconv.Itoa(f) + "\n" +
		"[Total commits (last year):](fg:magenta) " + strconv.Itoa(c) + "\n" +
		"[Total issues (last year):](fg:magenta) " + strconv.Itoa(i) + "\n" +
		"[Total PRs (last year):](fg:magenta) " + strconv.Itoa(p) + "\n" +
		"[Total repos:](fg:magenta) " + strconv.Itoa(user.Repositories.TotalCount) + "\n" +
		"[Total gists:](fg:magenta) " + strconv.Itoa(user.Gists.TotalCount) + "\n" +
		"[Total packages:](fg:magenta) " + strconv.Itoa(user.Packages.TotalCount) + "\n" +
		"[Total projects:](fg:magenta) " + strconv.Itoa(user.Projects.TotalCount) + "\n" +
		"[Organizations:](fg:magenta) " + strconv.Itoa(len(user.Organizations.Nodes)) + "\n" +
		"[Sponsors:](fg:magenta) " + strconv.Itoa(len(user.SponsorshipsAsMaintainer.Nodes)) + "\n" +
		"[Sponsoring:](fg:magenta) " + strconv.Itoa(len(user.SponsorshipsAsSponsor.Nodes)) + "\n" +
		"[Watching:](fg:magenta) " + strconv.Itoa(user.Watching.TotalCount) + " repos\n"

	if len(usedLicenses) > 0 {
		baseStats += "[Favorite licenses:](fg:magenta) " + usedLicenses[0]
	}

	return baseStats
}

func BuildReposStats(restClient *github.Client, user g.User, allRepos []*github.Repository) string {
	var baseRepos string

	starredRepos, starredNums := r.MostStarredRepos(restClient, allRepos)
	forkedRepos, forkedNums := r.MostForkedRepos(restClient, allRepos)

	if len(allRepos) > 0 {
		baseRepos = "Most starred repos" + "\n" +
			"[" + starredRepos[0] + ":](fg:red) " + strconv.FormatFloat(starredNums[0], 'f', -1, 64) + "\n" +
			"[" + starredRepos[1] + ":](fg:red) " + strconv.FormatFloat(starredNums[1], 'f', -1, 64) + "\n" +
			"[" + starredRepos[2] + ":](fg:red) " + strconv.FormatFloat(starredNums[2], 'f', -1, 64) + "\n" +
			"Most forked repos" + "\n" +
			"[" + forkedRepos[0] + ":](fg:red) " + strconv.FormatFloat(forkedNums[0], 'f', -1, 64) + "\n" +
			"[" + forkedRepos[1] + ":](fg:red) " + strconv.FormatFloat(forkedNums[1], 'f', -1, 64) + "\n" +
			"[" + forkedRepos[2] + ":](fg:red) " + strconv.FormatFloat(forkedNums[2], 'f', -1, 64)
	} else {
		baseRepos = "No repositories"
	}

	return baseRepos
}
