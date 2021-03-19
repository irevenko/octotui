package tui

import (
	"strconv"
	"strings"

	gh "../github"
	g "github.com/irevenko/octostats/graphql"
	r "github.com/irevenko/octostats/rest"
)

func buildProfileInfo(user g.User) string {
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

func buildProfileStats(user g.User) string {
	var baseStats string

	s, f, c, i, p := gh.FetchStats(ctx, client, qlClient, user.Login)

	allRepos := r.AllRepos(ctx, client, user.Login)
	usedLicenses, _ := r.MostUsedLicenses(client, allRepos)

	baseStats = "Profile Statistics" + "\n" +
		"[Total stars:](fg:magenta) " + strconv.Itoa(s) + "\n" +
		"[Total forks:](fg:magenta) " + strconv.Itoa(f) + "\n" +
		"[Total commits (2021):](fg:magenta) " + strconv.Itoa(c) + "\n" +
		"[Total issues (2021):](fg:magenta) " + strconv.Itoa(i) + "\n" +
		"[Total PRs (2021):](fg:magenta) " + strconv.Itoa(p) + "\n" +
		"[Total repos:](fg:magenta) " + strconv.Itoa(user.Repositories.TotalCount) + "\n" +
		"[Total gists:](fg:magenta) " + strconv.Itoa(user.Gists.TotalCount) + "\n" +
		"[Total packages:](fg:magenta) " + strconv.Itoa(user.Packages.TotalCount) + "\n" +
		"[Total projects:](fg:magenta) " + strconv.Itoa(user.Projects.TotalCount) + "\n" +
		"[Organizations:](fg:magenta) " + strconv.Itoa(len(user.Organizations.Nodes)) + "\n" +
		"[Sponsors:](fg:magenta) " + strconv.Itoa(len(user.SponsorshipsAsMaintainer.Nodes)) + "\n" +
		"[Sponsoring:](fg:magenta) " + strconv.Itoa(len(user.SponsorshipsAsSponsor.Nodes)) + "\n" +
		"[Watching:](fg:magenta) " + strconv.Itoa(user.Watching.TotalCount) + " repos\n" +
		"[Favorite licenses:](fg:magenta) " + usedLicenses[0]

	return baseStats
}

func buildReposStats(user g.User) string {
	var baseRepos string

	allRepos := r.AllRepos(ctx, client, user.Login)

	starredRepos, starredNums := r.MostStarredRepos(client, allRepos)
	forkedRepos, forkedNums := r.MostForkedRepos(client, allRepos)

	baseRepos = "Most starred repos" + "\n" +
		"[" + starredRepos[0] + ":](fg:red) " + strconv.Itoa(starredNums[0]) + "\n" +
		"[" + starredRepos[1] + ":](fg:red) " + strconv.Itoa(starredNums[1]) + "\n" +
		"[" + starredRepos[2] + ":](fg:red) " + strconv.Itoa(starredNums[2]) + "\n" +
		"Most forked repos" + "\n" +
		"[" + forkedRepos[0] + ":](fg:red) " + strconv.Itoa(forkedNums[0]) + "\n" +
		"[" + forkedRepos[1] + ":](fg:red) " + strconv.Itoa(forkedNums[1]) + "\n" +
		"[" + forkedRepos[2] + ":](fg:red) " + strconv.Itoa(forkedNums[2])

	return baseRepos
}
