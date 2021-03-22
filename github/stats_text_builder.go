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

func BuildUserInfo(user g.User) string {
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

func BuildUserStats(ctx context.Context, restClient *github.Client, qlClient *githubv4.Client, user g.User, allRepos []*github.Repository) string {
	var baseStats string

	s, f, c, i, p, rp := FetchUserStats(ctx, restClient, qlClient, user.Login, allRepos)
	usedLicenses, _ := r.MostUsedLicenses(restClient, allRepos)

	baseStats = "Profile Statistics" + "\n" +
		"[Total repos:](fg:magenta) " + strconv.Itoa(user.Repositories.TotalCount) + "\n" +
		"[Total gists:](fg:magenta) " + strconv.Itoa(user.Gists.TotalCount) + "\n" +
		"[Total stars:](fg:magenta) " + strconv.Itoa(s) + "\n" +
		"[Total forks:](fg:magenta) " + strconv.Itoa(f) + "\n" +
		"[Total commits (last year):](fg:magenta) " + strconv.Itoa(c) + "\n" +
		"[Opened issues (last year):](fg:magenta) " + strconv.Itoa(i) + "\n" +
		"[Opened PRs (last year):](fg:magenta) " + strconv.Itoa(p) + "\n" +
		"[Reviewed PRs (last year):](fg:magenta) " + strconv.Itoa(rp) + "\n" +
		"[Total packages:](fg:magenta) " + strconv.Itoa(user.Packages.TotalCount) + "\n" +
		"[Total projects:](fg:magenta) " + strconv.Itoa(user.Projects.TotalCount) + "\n" +
		"[Organizations:](fg:magenta) " + strconv.Itoa(user.Organizations.TotalCount) + "\n" +
		"[Sponsors:](fg:magenta) " + strconv.Itoa(user.SponsorshipsAsMaintainer.TotalCount) + "\n" +
		"[Sponsoring:](fg:magenta) " + strconv.Itoa(user.SponsorshipsAsSponsor.TotalCount) + " people\n" +
		"[Watching:](fg:magenta) " + strconv.Itoa(user.Watching.TotalCount) + " repos\n"

	if len(usedLicenses) > 0 {
		baseStats += "[Favorite license:](fg:magenta) " + usedLicenses[0]
	}

	return baseStats
}

func BuildOrgInfo(org g.Organization) string {
	var baseOrg string
	joinedAt := strings.Split(org.CreatedAt.String(), " ")

	if org.Name == "" {
		baseOrg = org.Login + "\n" +
			"[People:](fg:yellow) " + strconv.Itoa(org.MembersWithRole.TotalCount) + "\n" +
			"[Joined:](fg:yellow) " + joinedAt[0] + "\n"
	} else {
		baseOrg = org.Name + "\n" +
			"[People:](fg:yellow) " + strconv.Itoa(org.MembersWithRole.TotalCount) + "\n" +
			"[Joined:](fg:yellow) " + joinedAt[0] + "\n"
	}

	if org.Description != "" {
		baseOrg += "[Description:](fg:yellow) " + org.Description + "\n"
	}

	if org.Email != "" {
		baseOrg += "[Email:](fg:yellow) " + org.Email + "\n"
	}

	if org.Location != "" {
		baseOrg += "[Location:](fg:yellow) " + org.Location + "\n"
	}

	if org.TwitterUsername != "" {
		baseOrg += "[Twitter:](fg:yellow) @" + org.TwitterUsername + "\n"
	}

	if org.WebsiteURL != "" {
		baseOrg += "[Site:](fg:yellow) " + org.WebsiteURL + "\n"
	}

	return baseOrg
}

func BuildOrgStats(ctx context.Context, restClient *github.Client, qlClient *githubv4.Client, org g.Organization, allRepos []*github.Repository) string {
	var baseStats string

	s, f := FetchOrgStats(ctx, restClient, qlClient, org.Login, allRepos)
	usedLicenses, _ := r.MostUsedLicenses(restClient, allRepos)

	baseStats = "Organization Statistics" + "\n" +
		"[Total repos:](fg:magenta) " + strconv.Itoa(org.Repositories.TotalCount) + "\n" +
		"[Total stars:](fg:magenta) " + strconv.Itoa(s) + "\n" +
		"[Total forks:](fg:magenta) " + strconv.Itoa(f) + "\n" +
		"[Total packages:](fg:magenta) " + strconv.Itoa(org.Packages.TotalCount) + "\n" +
		"[Total projects:](fg:magenta) " + strconv.Itoa(org.Projects.TotalCount) + "\n" +
		"[Sponsors:](fg:magenta) " + strconv.Itoa(org.SponsorshipsAsMaintainer.TotalCount) + "\n" +
		"[Sponsoring:](fg:magenta) " + strconv.Itoa(org.SponsorshipsAsSponsor.TotalCount) + "\n"

	if len(usedLicenses) > 0 {
		baseStats += "[Favorite license:](fg:magenta) " + usedLicenses[0]
	}

	return baseStats
}

func BuildUserRepos(restClient *github.Client, user g.User, allRepos []*github.Repository) string {
	var baseRepos string

	starredRepos, starredNums := r.MostStarredRepos(restClient, allRepos)
	forkedRepos, forkedNums := r.MostForkedRepos(restClient, allRepos)

	if len(allRepos) > 0 {
		baseRepos = "Most starred repos\n"
		for i, v := range starredRepos {
			if i > 2 {
				break
			}
			if v != "" {
				baseRepos += "[" + starredRepos[i] + ":](fg:red) " + strconv.FormatFloat(starredNums[i], 'f', -1, 64) + "\n"
			}
		}
		baseRepos += "Most forked repos\n"
		for i, v := range forkedRepos {
			if i > 2 {
				break
			}
			if v != "" {
				baseRepos += "[" + forkedRepos[i] + ":](fg:red) " + strconv.FormatFloat(forkedNums[i], 'f', -1, 64) + "\n"
			}
		}
	} else {
		baseRepos = "No repositories"
	}

	return baseRepos
}

func BuildOrgRepos(restClient *github.Client, allRepos []*github.Repository) string {
	var baseRepos string

	starredRepos, starredNums := r.MostStarredRepos(restClient, allRepos)
	forkedRepos, forkedNums := r.MostForkedRepos(restClient, allRepos)

	if len(allRepos) > 0 {
		baseRepos = "Most starred repos" + "\n"
		for i, v := range starredRepos {
			if i > 8 {
				break
			}
			if v != "" {
				baseRepos += "[" + starredRepos[i] + ":](fg:red) " + strconv.FormatFloat(starredNums[i], 'f', -1, 64) + "\n"
			}
		}
		baseRepos += "Most forked repos" + "\n"
		for i, v := range forkedRepos {
			if i > 8 {
				break
			}
			if v != "" {
				baseRepos += "[" + forkedRepos[i] + ":](fg:red) " + strconv.FormatFloat(forkedNums[i], 'f', -1, 64) + "\n"
			}
		}
	} else {
		baseRepos = "No repositories"
	}

	return baseRepos
}
