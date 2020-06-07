package github

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
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

func searchPositions(keywords []string, location string) ([]jobListingAPIModel, error) {
	var result []jobListingAPIModel
	searchURL := generateSearchURL(keywords, location)

	request, err := http.NewRequest("GET", searchURL, nil)
	if err != nil {
		log.Println(err)
		return result, err
	}

	responseBody, err := executeRequest(request)
	if err != nil {
		return result, err
	}

	var models []jobListingAPIModel
	if err := json.Unmarshal(responseBody, &models); err != nil {
		log.Println(err)
		return result, err
	}

	result = models
	return result, nil
}
