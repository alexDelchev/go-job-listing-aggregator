package github

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// SourceName defines github as the data source
const SourceName string = "github"

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
