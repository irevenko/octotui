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
	"time"

	"github.com/google/go-github/v33/github"
	g "github.com/irevenko/octostats/graphql"
	r "github.com/irevenko/octostats/rest"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	gh "github.com/irevenko/octotui/github"
	h "github.com/irevenko/octotui/helpers"
)

func SetupProfileInfo(user g.User) *widgets.Paragraph {
	p := widgets.NewParagraph()
	p.WrapText = true
	p.Border = true
	p.Text = gh.BuildUserInfo(user)

	return p
}

func SetupProfileStats(user g.User, allRepos []*github.Repository) *widgets.Paragraph {
	p := widgets.NewParagraph()
	p.WrapText = true
	p.Text = gh.BuildUserStats(Ctx, RestClient, qlClient, user, allRepos)
	p.Border = true

	return p
}

func SetupReposStats(user g.User, allRepos []*github.Repository) *widgets.Paragraph {
	p := widgets.NewParagraph()
	p.WrapText = true
	p.Border = true
	p.Text = gh.BuildUserRepos(RestClient, user, allRepos)

	return p
}

func SetupOrgInfo(org g.Organization) *widgets.Paragraph {
	p := widgets.NewParagraph()
	p.WrapText = true
	p.Border = true
	p.Text = gh.BuildOrgInfo(org)

	return p
}

func SetupOrgStats(org g.Organization, allRepos []*github.Repository) *widgets.Paragraph {
	p := widgets.NewParagraph()
	p.WrapText = true
	p.Text = gh.BuildOrgStats(Ctx, RestClient, qlClient, org, allRepos)
	p.Border = true

	return p
}

func SetupOrgRepos(org g.Organization, allRepos []*github.Repository) *widgets.Paragraph {
	p := widgets.NewParagraph()
	p.WrapText = true
	p.Border = true
	p.Text = gh.BuildOrgRepos(RestClient, allRepos)

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
	img.Title = login + "'s GitHub"

	return img, images
}

func SetupLangsByCommits(user g.User) *widgets.PieChart {
	year, _, _ := time.Now().Date()
	langs, commits, err := g.LanguagesByCommit(qlClient, user.Login, year-1, year)
	if err != nil {
		log.Fatalf("Couldn't get langs by commit for: %v: %v", user.Login, err)
	}

	pc := widgets.NewPieChart()
	pc.Title = "Languages by commit"

	boundNum := h.UserDataBound(langs)
	if boundNum == 0 {
		pc.Title = "Languages by commit (no commits)"
	} else {
		pc.Data = commits[:boundNum]
	}

	pc.AngleOffset = .15 * math.Pi
	pc.LabelFormatter = func(i int, v float64) string {
		return fmt.Sprintf("%.00f"+" %s", v, langs[i])
	}

	return pc
}

func SetupLangsByRepo(allRepos []*github.Repository, accType string) *widgets.PieChart {
	usedLangs, langsNum := r.LanguagesByRepo(RestClient, allRepos)

	pc := widgets.NewPieChart()
	pc.Title = "Languages by repo"

	if accType == "user" {
		boundNum := h.UserDataBound(usedLangs)
		if boundNum == 0 {
			pc.Title = "Langs by repo (no repos)"
		} else {
			pc.Data = langsNum[:boundNum]
		}
	}

	if accType == "organization" {
		boundNum := h.OrgDataBound(usedLangs)
		if boundNum == 0 {
			pc.Title = "Langs by repo (no repos)"
		} else {
			pc.Data = langsNum[:boundNum]
		}
	}

	pc.AngleOffset = .15 * math.Pi
	pc.LabelFormatter = func(i int, v float64) string {
		return fmt.Sprintf("%.00f"+" %s", v, usedLangs[i])
	}

	return pc
}

func SetupStarsPerLangs(allRepos []*github.Repository, accType string) *widgets.BarChart {
	starsPerL, starsNum := r.StarsPerLanguage(RestClient, allRepos)

	bc := widgets.NewBarChart()
	bc.Title = "Stars per language"

	if accType == "user" {
		bound := h.UserDataBound(starsPerL)
		if bound == 0 || h.AllZero(starsNum[:bound]) {
			bc.Title = "Stars per language (no stars)"
		} else {
			bc.Data = starsNum[:bound]
			bc.Labels = starsPerL[:bound]
		}
	}

	if accType == "organization" {
		bound := h.OrgDataBound(starsPerL)
		if bound == 0 || h.AllZero(starsNum[:bound]) {
			bc.Title = "Stars per language (no stars)"
		} else {
			bc.Data = starsNum[:bound]
			bc.Labels = starsPerL[:bound]
		}
	}

	bc.BarColors = []ui.Color{ui.ColorMagenta, ui.ColorGreen, ui.ColorYellow, ui.ColorBlue}
	bc.LabelStyles = []ui.Style{ui.NewStyle(ui.ColorCyan)}
	bc.NumStyles = []ui.Style{ui.NewStyle(ui.ColorWhite)}
	bc.BarWidth = 8
	bc.BarGap = 3

	return bc
}

func SetupForksPerLangs(allRepos []*github.Repository, accType string) *widgets.BarChart {
	forksPerL, forksNum := r.ForksPerLanguage(RestClient, allRepos)

	bc := widgets.NewBarChart()
	bc.Title = "Forks per language"

	if accType == "user" {
		bound := h.UserDataBound(forksPerL)
		if bound == 0 || h.AllZero(forksNum[:bound]) {
			bc.Title = "Forks per language (no forks)"
		} else {
			bc.Data = forksNum[:bound]
			bc.Labels = forksPerL[:bound]
		}
	}

	if accType == "organization" {
		bound := h.OrgDataBound(forksPerL)
		if bound == 0 || h.AllZero(forksNum[:bound]) {
			bc.Title = "Forks per language (no forks)"
		} else {
			bc.Data = forksNum[:bound]
			bc.Labels = forksPerL[:bound]
		}
	}

	bc.BarColors = []ui.Color{ui.ColorMagenta, ui.ColorGreen, ui.ColorYellow, ui.ColorBlue}
	bc.LabelStyles = []ui.Style{ui.NewStyle(ui.ColorCyan)}
	bc.NumStyles = []ui.Style{ui.NewStyle(ui.ColorWhite)}
	bc.BarWidth = 8
	bc.BarGap = 3

	return bc
}

func SetupContribsSparkline(user g.User) *widgets.SparklineGroup {
	_, contribs, err := g.YearActivity(qlClient, user.Login)
	if err != nil {
		log.Fatalf("Couldn't get year activity for: %v: %v", user.Login, err)
	}
	timeSpan := 75

	sl := widgets.NewSparkline()
	sl.Data = contribs[len(contribs)-timeSpan : len(contribs)]
	sl.TitleStyle.Fg = ui.ColorWhite
	sl.LineColor = ui.ColorCyan

	slg := widgets.NewSparklineGroup(sl)
	slg.Title = "Activity for the last " + strconv.Itoa(timeSpan) + " days"

	return slg
}
