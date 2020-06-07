package weworkremotely

import (
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
