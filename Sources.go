package newsApiSdk

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Sources struct {
	Category string
	Language string
	Country  string
	ApiKey   string
}

func GetSources(sources Sources) ([]byte, error) {

	if len(sources.ApiKey) == 0 {
		fmt.Println("Missing api key")
	}

	response, err := http.Get(sources.buildQuery())
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s Sources) buildQuery() string {
	query := baseUrl + "/sources?apiKey=" + s.ApiKey

	if (s == Sources{}) {
		return query
	}
	// see if it's possible to keep it DRY by adding a for loop
	if len(s.Category) > 0 {
		query += "&category=" + s.Category
	}
	if len(s.Language) > 0 {
		query += "&language=" + s.Language
	}
	if len(s.Country) > 0 {
		query += "&country=" + s.Country
	}

	return query
}
