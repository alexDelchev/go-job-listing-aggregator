package github

import (
	"strings"
)

const apiURL string = "https://jobs.github.com/positions.json?description=::description::&location=::location::&full_time=true"

func generateSearchURL(keywords []string, location string) string {
	var result string
	description := strings.Join(keywords[:], "+")

	result = strings.Replace(apiURL, "::description::", description, -1)
	result = strings.Replace(result, "::location::", location, -1)

	return result
}
