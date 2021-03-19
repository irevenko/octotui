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

	g "github.com/irevenko/octostats/graphql"
	r "github.com/irevenko/octostats/rest"

	ui "github.com/gizak/termui"
	"github.com/gizak/termui/widgets"
)

func SetupProfileInfo(user g.User) *widgets.Paragraph {
	p := widgets.NewParagraph()
	p.WrapText = true
	text := buildProfileInfo(user)
	p.Text = text
	p.Border = true
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

func SetupReposStats(user g.User) *widgets.Paragraph {
	p := widgets.NewParagraph()
	p.WrapText = true
	p.Border = true
	p.Text = buildReposStats(user)
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
	pc := widgets.NewPieChart()
	pc.Title = "Languages by commit"
	pc.SetRect(35, 35, 70, 20)

	langs, commits := g.LanguagesByCommit(qlClient, user.Login, 2020, 2021)

	var data []float64

	for _, v := range commits {
		data = append(data, float64(v))
	}

	pc.Data = data[:4]
	pc.AngleOffset = .15 * math.Pi
	pc.LabelFormatter = func(i int, v float64) string {
		return fmt.Sprintf("%.00f"+" %s", v, langs[i])
	}

	return pc
}

func SetupLangsByRepo(user g.User) *widgets.PieChart {
	pc := widgets.NewPieChart()
	pc.Title = "Languages by repo"
	pc.SetRect(105, 35, 70, 20)

	allRepos := r.AllRepos(ctx, client, user.Login)
	var data []float64

	usedLangs, langsNum := r.LanguagesByRepo(client, allRepos)
	for _, v := range langsNum {
		data = append(data, float64(v))
	}

	pc.Data = data[:4]
	pc.AngleOffset = .15 * math.Pi
	pc.LabelFormatter = func(i int, v float64) string {
		return fmt.Sprintf("%.00f"+" %s", v, usedLangs[i])
	}

	return pc
}

func SetupStarsPerLangs(user g.User) *widgets.BarChart {
	var data []float64
	allRepos := r.AllRepos(ctx, client, user.Login)

	starsPerL, starsNum := r.StarsPerLanguage(client, allRepos)

	for _, v := range starsNum {
		data = append(data, float64(v))
	}

	bc := widgets.NewBarChart()
	bc.Data = data[:4]
	bc.Labels = starsPerL[:4]
	bc.Title = "Stars per language"
	bc.SetRect(150, 35, 105, 20)
	bc.BarWidth = 5
	bc.BarColors = []ui.Color{ui.ColorMagenta, ui.ColorGreen, ui.ColorYellow, ui.ColorBlue}
	bc.LabelStyles = []ui.Style{ui.NewStyle(ui.ColorCyan)}
	bc.NumStyles = []ui.Style{ui.NewStyle(ui.ColorWhite)}
	bc.BarWidth = 8
	bc.BarGap = 3

	return bc
}

func SetupForksPerLangs(user g.User) *widgets.BarChart {
	var data []float64
	allRepos := r.AllRepos(ctx, client, user.Login)

	forksPerL, forksNum := r.ForksPerLanguage(client, allRepos)

	for _, v := range forksNum {
		data = append(data, float64(v))
	}

	bc := widgets.NewBarChart()
	bc.Data = data[:4]
	bc.Labels = forksPerL[:4]
	bc.Title = "Forks per language"
	bc.SetRect(150, 10, 105, 20)
	bc.BarWidth = 5
	bc.BarColors = []ui.Color{ui.ColorMagenta, ui.ColorGreen, ui.ColorYellow, ui.ColorBlue}
	bc.LabelStyles = []ui.Style{ui.NewStyle(ui.ColorCyan)}
	bc.NumStyles = []ui.Style{ui.NewStyle(ui.ColorWhite)}
	bc.BarWidth = 8
	bc.BarGap = 3

	return bc
}

func SetupContribsSparkline(user g.User) *widgets.SparklineGroup {
	_, contribs := g.YearActivity(qlClient, user.Login)

	sl := widgets.NewSparkline()
	sl.Data = contribs[len(contribs)-75 : len(contribs)]
	sl.TitleStyle.Fg = ui.ColorWhite
	sl.LineColor = ui.ColorCyan

	slg := widgets.NewSparklineGroup(sl)
	slg.Title = "Activity for the last 75 days"
	slg.SetRect(150, 0, 70, 10)
	return slg
}
