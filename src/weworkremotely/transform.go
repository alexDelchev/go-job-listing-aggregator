package weworkremotely

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// SourceName defines WeWorkRemotely as the data source
const SourceName string = "weworkremotely"

func extractTextFromHTML(data string) string {
	var html string

	// Package text in a html tag if not present
	if strings.Contains(data, "<html>") {
		html = data
	} else {
		html = fmt.Sprintf("<html><body>%s</body></html>", data)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Println(err)
		return data
	}

	return doc.Text()
}

// Returns the (company name, position title) wich construct the title property
func deconstructTitle(title string) (string, string) {
	tokens := strings.SplitN(title, ":", 2)

	return strings.TrimSpace(tokens[0]), strings.TrimSpace(tokens[1])
}
