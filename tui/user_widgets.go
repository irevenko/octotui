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

	"github.com/google/go-github/github"
	g "github.com/irevenko/octostats/graphql"
	r "github.com/irevenko/octostats/rest"

	gh "../github"
	ui "github.com/gizak/termui"
	"github.com/gizak/termui/widgets"
)

func SetupProfileInfo(user g.User) *widgets.Paragraph {
	p := widgets.NewParagraph()
	p.WrapText = true
	p.Border = true
	p.Text = gh.BuildProfileInfo(user)
	p.SetRect(0, 35, 35, 14)

	return p
}

func SetupProfileStats(user g.User, allRepos []*github.Repository) *widgets.Paragraph {
	p := widgets.NewParagraph()
	p.WrapText = true
	p.Text = gh.BuildProfileStats(ctx, restClient, qlClient, user, allRepos)
	p.Border = true
	p.SetRect(35, 0, 70, 20)

	return p
}

func SetupReposStats(user g.User, allRepos []*github.Repository) *widgets.Paragraph {
	p := widgets.NewParagraph()
	p.WrapText = true
	p.Border = true
	p.Text = gh.BuildReposStats(restClient, user, allRepos)
	p.SetRect(70, 10, 105, 20)

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
	img.Title = login + "'s GitHub"

	return img, images
}

func SetupLangsByCommits(user g.User) *widgets.PieChart {
	year, _, _ := time.Now().Date()
	langs, commits := g.LanguagesByCommit(qlClient, user.Login, year-1, year)

	pc := widgets.NewPieChart()
	pc.Title = "Languages by commit"
	pc.SetRect(35, 35, 70, 20)

	boundNum := defineDataBound(langs)
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

func SetupLangsByRepo(user g.User, allRepos []*github.Repository) *widgets.PieChart {
	usedLangs, langsNum := r.LanguagesByRepo(restClient, allRepos)

	pc := widgets.NewPieChart()
	pc.Title = "Languages by repo"
	pc.SetRect(105, 35, 70, 20)

	boundNum := defineDataBound(usedLangs)
	if boundNum == 0 {
		pc.Title = "Languages by repo (no repos)"
	} else {
		pc.Data = langsNum[:boundNum]
	}

	pc.AngleOffset = .15 * math.Pi
	pc.LabelFormatter = func(i int, v float64) string {
		return fmt.Sprintf("%.00f"+" %s", v, usedLangs[i])
	}

	return pc
}

func SetupStarsPerLangs(user g.User, allRepos []*github.Repository) *widgets.BarChart {
	starsPerL, starsNum := r.StarsPerLanguage(restClient, allRepos)

	bc := widgets.NewBarChart()
	bc.Title = "Stars per language"
	bc.SetRect(150, 35, 105, 20)

	boundNum := defineDataBound(starsPerL)
	if boundNum == 0 {
		bc.Title = "Stars per language (no languages)"
	} else {
		bc.Data = starsNum[:boundNum]
		bc.Labels = starsPerL[:boundNum]
	}

	bc.BarColors = []ui.Color{ui.ColorMagenta, ui.ColorGreen, ui.ColorYellow, ui.ColorBlue}
	bc.LabelStyles = []ui.Style{ui.NewStyle(ui.ColorCyan)}
	bc.NumStyles = []ui.Style{ui.NewStyle(ui.ColorWhite)}
	bc.BarWidth = 8
	bc.BarGap = 3

	return bc
}

func SetupForksPerLangs(user g.User, allRepos []*github.Repository) *widgets.BarChart {
	forksPerL, forksNum := r.ForksPerLanguage(restClient, allRepos)

	bc := widgets.NewBarChart()
	bc.Title = "Forks per language"
	bc.SetRect(150, 10, 105, 20)

	boundNum := defineDataBound(forksPerL)
	if boundNum == 0 {
		bc.Title = "Forks per language (no languages)"
	} else {
		bc.Data = forksNum[:boundNum]
		bc.Labels = forksPerL[:boundNum]
	}

	bc.BarColors = []ui.Color{ui.ColorMagenta, ui.ColorGreen, ui.ColorYellow, ui.ColorBlue}
	bc.LabelStyles = []ui.Style{ui.NewStyle(ui.ColorCyan)}
	bc.NumStyles = []ui.Style{ui.NewStyle(ui.ColorWhite)}
	bc.BarWidth = 8
	bc.BarGap = 3

	return bc
}

func SetupContribsSparkline(user g.User) *widgets.SparklineGroup {
	_, contribs := g.YearActivity(qlClient, user.Login)
	timeSpan := 75

	sl := widgets.NewSparkline()
	sl.Data = contribs[len(contribs)-timeSpan : len(contribs)]
	sl.TitleStyle.Fg = ui.ColorWhite
	sl.LineColor = ui.ColorCyan

	slg := widgets.NewSparklineGroup(sl)
	slg.Title = "Activity for the last " + strconv.Itoa(timeSpan) + " days"
	slg.SetRect(150, 0, 70, 10)

	return slg
}

func defineDataBound(slice []string) int {
	if len(slice) > 4 { // 4 is max amount of entries for pie/bar chart
		return 4
	} else if len(slice) < 1 {
		return 0
	} else if len(slice) < 2 {
		return 1
	} else if len(slice) < 3 {
		return 2
	} else if len(slice) < 4 {
		return 3
	} else if len(slice) < 5 {
		return 4
	}

	return 0
}
