package newsApiSdk

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type HeadlineOptions struct {
	Country  string
	Category string
	Sources  string
	Keyword  string
	PageSize int
	Page     int
	ApiKey   string
}

func GetTopHeadlines(options HeadlineOptions) ([]byte, error) {

	if len(options.ApiKey) == 0 {
		fmt.Println("Missing api key")
	}

	response, err := http.Get(buildTopHeadlinesQuery(options))
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func buildTopHeadlinesQuery(options HeadlineOptions) string {
	query := baseUrl + "/top-headlines?apiKey=" + options.ApiKey

	if (HeadlineOptions{}) == options {
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
