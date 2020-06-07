package stackoverflow

import (
	"strings"
)

const rssURL = "https://stackoverflow.com/jobs/feed?q=::query::&l=::location::&u=Km&d=20"

func generateSearchURL(keywords []string, location string) string {
	var result string
	description := strings.Join(keywords[:], "+")

	result = strings.Replace(rssURL, "::query::", description, -1)
	result = strings.Replace(result, "::location::", location, -1)

	return result
}
