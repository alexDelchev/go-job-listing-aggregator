package stackoverflow

import (
	"io/ioutil"
	"log"
	"net/http"
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

func executeRequest(request *http.Request) ([]byte, error) {
	var result []byte
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Println(err)
		return result, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return result, err
	}

	result = body
	return result, nil
}
