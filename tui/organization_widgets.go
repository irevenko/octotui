package tui

import (
	gh "../github"
	"github.com/gizak/termui/widgets"
	"github.com/google/go-github/github"
	g "github.com/irevenko/octostats/graphql"
)

func SetupOrgInfo(org g.Organization) *widgets.Paragraph {
	p := widgets.NewParagraph()
	p.WrapText = true
	p.Border = true
	p.Text = gh.BuildOrgInfo(org)
	p.SetRect(0, 35, 35, 14)

	return p
}

func SetupOrgStats(org g.Organization, allRepos []*github.Repository) *widgets.Paragraph {
	p := widgets.NewParagraph()
	p.WrapText = true
	p.Text = gh.BuildOrgStats(ctx, restClient, qlClient, org, allRepos)
	p.Border = true
	p.SetRect(35, 0, 70, 12)

	return p
}

func SetupOrgRepos(org g.Organization, allRepos []*github.Repository) *widgets.Paragraph {
	p := widgets.NewParagraph()
	p.WrapText = true
	p.Border = true
	p.Text = gh.BuildOrgRepos(restClient, allRepos)
	p.SetRect(35, 12, 70, 35)

	return p
}
