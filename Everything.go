package newsapi

import (
	"encoding/json"
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

type EverythingResponse struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []Article `json:"articles"`
	Code         string    `json:"code"`
	Message      string    `json:"message"`
}

func GetEverything(everything Everything) (EverythingResponse, error) {

	if len(everything.ApiKey) == 0 {
		fmt.Println("Missing api key")
	}

	response, err := http.Get(everything.buildQuery())
	if err != nil {
		return EverythingResponse{}, err
	}

	body, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		return EverythingResponse{}, readErr
	}

	everythingResponse := EverythingResponse{}
	everythingErr := json.Unmarshal(body, &everythingResponse)

	if everythingErr != nil {
		return everythingResponse, everythingErr
	}

	if everythingResponse.Status == "error" {
		return everythingResponse, everythingErr
	}

	return everythingResponse, nil
}

func (e Everything) buildQuery() string {
	query := baseUrl + "/everything?apiKey=" + e.ApiKey

	if e == (Everything{}) {
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
