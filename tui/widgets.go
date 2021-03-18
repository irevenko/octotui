package tui

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"

	gh "../github"
	g "github.com/irevenko/octostats/graphql"
	r "github.com/irevenko/octostats/rest"

	"github.com/gizak/termui/widgets"
)

func SetupProfileInfo(user g.User) *widgets.Paragraph {
	p := widgets.NewParagraph()
	p.WrapText = true
	text := buildProfileInfo(user)
	p.Text = text
	p.Border = false
	p.SetRect(0, 35, 35, 14)

	return p
}

func SetupProfileStats(user g.User) *widgets.Paragraph {
	p := widgets.NewParagraph()
	p.WrapText = true
	text := buildProfileStats(user)
	p.Text = text
	p.Border = true
	p.SetRect(35, 0, 70, 20)

	return p
}

func SetupImage(profileImg string, login string) (*widgets.Image, []image.Image) {
	var images []image.Image

	resp, err := http.Get(profileImg)
	if err != nil {
		log.Fatalf("failed to fetch image: %v", err)
	}

	image, _, err := image.Decode(resp.Body)
	if err != nil {
		log.Fatalf("failed to decode fetched image: %v", err)
	}

	images = append(images, image)

	img := widgets.NewImage(nil)
	img.SetRect(0, 0, 30, 14)
	img.Title = login + " profile stats"

	return img, images
}

func SetupLangsByCommits(user g.User) *widgets.PieChart {
	pc := widgets.NewPieChart()
	pc.Title = "Languages by commit"
	pc.SetRect(35, 35, 75, 20)

	langs, commits := g.LanguagesByCommit(qlClient, user.Login, 2020, 2021)

	var data []float64

	for _, v := range commits {
		data = append(data, float64(v))
	}

	pc.Data = data[:6]
	pc.AngleOffset = .15 * math.Pi
	pc.LabelFormatter = func(i int, v float64) string {
		return fmt.Sprintf("%.00f"+" %s", v, langs[i])
	}

	return pc
}

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

	baseStats = "[Total stars:](fg:green) " + strconv.Itoa(s) + "\n" +
		"[Total forks:](fg:green) " + strconv.Itoa(f) + "\n" +
		"[Total commits (2021):](fg:green) " + strconv.Itoa(c) + "\n" +
		"[Total issues (2021):](fg:green) " + strconv.Itoa(i) + "\n" +
		"[Total PRs (2021):](fg:green) " + strconv.Itoa(p) + "\n" +
		"[Total repos:](fg:green) " + strconv.Itoa(user.Repositories.TotalCount) + "\n" +
		"[Total gists:](fg:green) " + strconv.Itoa(user.Gists.TotalCount) + "\n" +
		"[Total packages:](fg:green) " + strconv.Itoa(user.Packages.TotalCount) + "\n" +
		"[Total projects:](fg:green) " + strconv.Itoa(user.Projects.TotalCount) + "\n" +
		"[Organizations:](fg:green) " + strconv.Itoa(len(user.Organizations.Nodes)) + "\n" +
		"[Sponsors:](fg:green) " + strconv.Itoa(len(user.SponsorshipsAsMaintainer.Nodes)) + "\n" +
		"[Sponsoring:](fg:green) " + strconv.Itoa(len(user.SponsorshipsAsSponsor.Nodes)) + "\n" +
		"[Watching:](fg:green) " + strconv.Itoa(user.Watching.TotalCount) + " repos\n" +
		"[Favorite license:](fg:green) " + usedLicenses[0]

	return baseStats
}
