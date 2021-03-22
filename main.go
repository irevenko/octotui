package main

import (
	gh "./github"
	tui "./tui"
)

func main() {
	results := gh.SearchUser(tui.Ctx, tui.RestClient, "irevenko")
	tui.RenderList(results)
}
