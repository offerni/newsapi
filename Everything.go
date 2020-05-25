package newsApiSdk

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type EverythingOptions struct {
	KeywordBody    string
	KeywordTitle   string
	Sources        string
	Domains        string
	ExcludeDomains string
	FromDate       string
	ToDate         string
	Language       string
	SortBy         string
	PageSize       int
	Page           int
	ApiKey         string
}

func GetEverything(everythingOptions EverythingOptions) ([]byte, error) {

	if len(everythingOptions.ApiKey) == 0 {
		fmt.Println("Missing api key")
	}

	response, err := http.Get(buildEverythingQuery(everythingOptions))
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func buildEverythingQuery(everythingOptions EverythingOptions) string {
	query := baseUrl + "/everything?apiKey=" + everythingOptions.ApiKey

	if (EverythingOptions{}) == everythingOptions {
		return query
	}
	// see if it's possible to keep it DRY by adding a for loop
	if len(everythingOptions.KeywordBody) > 0 {
		query += "&q=" + everythingOptions.KeywordBody
	}
	if len(everythingOptions.KeywordTitle) > 0 {
		query += "&qInTitle=" + everythingOptions.KeywordTitle
	}
	if len(everythingOptions.Sources) > 0 {
		query += "&sources=" + everythingOptions.Sources
	}
	if len(everythingOptions.Domains) > 0 {
		query += "&domains=" + everythingOptions.Domains
	}
	if len(everythingOptions.ExcludeDomains) > 0 {
		query += "&excludeDomains=" + everythingOptions.ExcludeDomains
	}
	if len(everythingOptions.FromDate) > 0 {
		query += "&from=" + everythingOptions.FromDate
	}
	if len(everythingOptions.ToDate) > 0 {
		query += "&to=" + everythingOptions.ToDate
	}
	if len(everythingOptions.Language) > 0 {
		query += "&language=" + everythingOptions.Language
	}
	if len(everythingOptions.SortBy) > 0 {
		query += "&sortBy=" + everythingOptions.SortBy
	}
	if everythingOptions.PageSize > 0 {
		query += "&pageSize=" + strconv.Itoa(everythingOptions.PageSize)
	}
	if everythingOptions.Page > 0 {
		query += "&page=" + strconv.Itoa(everythingOptions.Page)
	}

	return query
}
