package newsApiSdk

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Headlines struct {
	Country  string
	Category string
	Sources  string
	Keyword  string
	PageSize int
	Page     int
	ApiKey   string
}

func GetTopHeadlines(options Headlines) ([]byte, error) {

	if len(options.ApiKey) == 0 {
		fmt.Println("Missing api key")
	}

	response, err := http.Get(options.buildQuery())
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (h Headlines) buildQuery() string {
	query := baseUrl + "/top-headlines?apiKey=" + h.ApiKey

	if h == (Headlines{}) {
		return query
	}
	// see if it's possible to keep it DRY by adding a for loop
	if len(h.Country) > 0 {
		query += "&country=" + h.Country
	}
	if len(h.Category) > 0 {
		query += "&category=" + h.Category
	}
	if len(h.Sources) > 0 {
		query += "&sources=" + h.Sources
	}
	if len(h.Keyword) > 0 {
		query += "&q=" + h.Keyword
	}
	if h.PageSize > 0 {
		query += "&pageSize=" + strconv.Itoa(h.PageSize)
	}
	if h.Page > 0 {
		query += "&page=" + strconv.Itoa(h.Page)
	}

	return query
}
