package stackoverflow

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// SourceName defines stackoverflow as the data source
const SourceName string = "stackoverflow"

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
