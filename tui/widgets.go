package tui

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"net/http"

	"github.com/gizak/termui/widgets"
)

func SetupProfileInfo(username string) *widgets.Paragraph {
	p := widgets.NewParagraph()
	p.WrapText = true
	p.Text = username + " profile stats"
	p.SetRect(0, 25, 30, 15)

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
