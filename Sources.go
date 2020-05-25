package newsApiSdk

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type SourcesOptions struct {
	Category string
	Language string
	Country  string
	ApiKey   string
}

func GetSources(sourcesOptions SourcesOptions) ([]byte, error) {

	if len(sourcesOptions.ApiKey) == 0 {
		fmt.Println("Missing api key")
	}

	response, err := http.Get(buildSourcesQuery(sourcesOptions))
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func buildSourcesQuery(sourcesOptions SourcesOptions) string {
	query := baseUrl + "/sources?apiKey=" + sourcesOptions.ApiKey

	if (SourcesOptions{}) == sourcesOptions {
		return query
	}
	// see if it's possible to keep it DRY by adding a for loop
	if len(sourcesOptions.Category) > 0 {
		query += "&category=" + sourcesOptions.Category
	}
	if len(sourcesOptions.Language) > 0 {
		query += "&language=" + sourcesOptions.Language
	}
	if len(sourcesOptions.Country) > 0 {
		query += "&country=" + sourcesOptions.Country
	}

	return query
}
