package weworkremotely

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
)

const rssURL string = "https://weworkremotely.com/categories/remote-programming-jobs.rss"

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

func searchPositions(keywords []string) ([]jobListingRSSModel, error) {
	var result []jobListingRSSModel

	request, err := http.NewRequest("GET", rssURL, nil)
	if err != nil {
		log.Println(err)
		return result, err
	}

	responseBody, err := executeRequest(request)

	var response jobsRSSFeed
	if err := xml.Unmarshal(responseBody, &response); err != nil {
		log.Println(err)
		return result, err
	}

	result = response.Channel.PositionListings
	return result, nil
}
