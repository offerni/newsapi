package newsApiSdk

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Everything struct {
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

func GetEverything(everything Everything) ([]byte, error) {

	if len(everything.ApiKey) == 0 {
		fmt.Println("Missing api key")
	}

	response, err := http.Get(everything.buildQuery())
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (e Everything) buildQuery() string {
	query := baseUrl + "/everything?apiKey=" + e.ApiKey

	if (Everything{}) == e {
		return query
	}
	// see if it's possible to keep it DRY by adding a for loop
	if len(e.KeywordBody) > 0 {
		query += "&q=" + e.KeywordBody
	}
	if len(e.KeywordTitle) > 0 {
		query += "&qInTitle=" + e.KeywordTitle
	}
	if len(e.Sources) > 0 {
		query += "&sources=" + e.Sources
	}
	if len(e.Domains) > 0 {
		query += "&domains=" + e.Domains
	}
	if len(e.ExcludeDomains) > 0 {
		query += "&excludeDomains=" + e.ExcludeDomains
	}
	if len(e.FromDate) > 0 {
		query += "&from=" + e.FromDate
	}
	if len(e.ToDate) > 0 {
		query += "&to=" + e.ToDate
	}
	if len(e.Language) > 0 {
		query += "&language=" + e.Language
	}
	if len(e.SortBy) > 0 {
		query += "&sortBy=" + e.SortBy
	}
	if e.PageSize > 0 {
		query += "&pageSize=" + strconv.Itoa(e.PageSize)
	}
	if e.Page > 0 {
		query += "&page=" + strconv.Itoa(e.Page)
	}

	return query
}
