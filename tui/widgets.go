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

	pc.Data = data[:4]
	pc.AngleOffset = .15 * math.Pi
	pc.LabelFormatter = func(i int, v float64) string {
		return fmt.Sprintf("%.00f"+" %s", v, langs[i])
	}

	return pc
}
