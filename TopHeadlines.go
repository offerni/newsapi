package newsApiSdk

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

const baseUrl string = "https://newsapi.org/v2"

type Options struct {
	Country  string
	Category string
	Sources  string
	Keyword  string
	PageSize int
	Page     int
	ApiKey   string
}

func GetTopHeadlines(options Options) ([]byte, error) {

	if len(options.ApiKey) == 0 {
		fmt.Println("Missing api key")
	}

	response, err := http.Get(buildQuery(options))
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func buildQuery(options Options) string {
	query := baseUrl + "/top-headlines?apiKey=" + options.ApiKey

	if (Options{}) == options {
		return query
	}
	// see if it's possible to keep it DRY by adding a for loop
	if len(options.Country) > 0 {
		query += "&country=" + options.Country
	}
	if len(options.Category) > 0 {
		query += "&category=" + options.Category
	}
	if len(options.Sources) > 0 {
		query += "&sources=" + options.Sources
	}
	if len(options.Keyword) > 0 {
		query += "&q=" + options.Keyword
	}
	if options.PageSize > 0 {
		query += "&pageSize=" + strconv.Itoa(options.PageSize)
	}
	if options.Page > 0 {
		query += "&page=" + strconv.Itoa(options.Page)
	}

	return query
}
