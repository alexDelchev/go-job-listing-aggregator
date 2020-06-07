package jobsbg

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const searchPageURL = "https://www.jobs.bg/front_job_search.php?add_sh=1&location_sid=&distance=&categories[]=0&type_all=0&position_level_all=0&company_type[]=0&::keywords::keyword=&job_languages_all=0&salary_from=0&last=0"

func generateKeywordsQueryParameter(keywords []string) string {
	var builder strings.Builder

	for _, word := range keywords {
		parameterRepresentation := fmt.Sprintf("keywords[]=%s&", word)
		builder.WriteString(parameterRepresentation)
	}

	return builder.String()
}

func generateSearchURL(keywords []string) string {
	keywordsQueryParameter := generateKeywordsQueryParameter(keywords)
	return strings.Replace(searchPageURL, "::keywords::", keywordsQueryParameter, -1)
}

func executeRequest(request *http.Request) (string, error) {
	var result string
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

	result = string(body)
	return result, nil
}

func searchListings(keywords []string) (string, error) {
	var result string
	url := generateSearchURL(keywords)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		return result, err
	}

	response, err := executeRequest(request)
	if err != nil {
		return result, err
	}

	result = response
	return result, err
}

func getListingPage(url string) (string, error) {
	var result string

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		return result, err
	}

	response, err := executeRequest(request)
	if err != nil {
		return result, err
	}

	result = response
	return response, nil
}
