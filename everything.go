package newsapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type EverythingOpts struct {
	Q              string
	QInTitle       string
	Sources        string
	Domains        string
	ExcludeDomains string
	From           string
	To             string
	Language       string
	SortBy         string
	PageSize       int
	Page           int
}

type everythingResponse struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []article `json:"articles"`
	Code         string    `json:"code"`
	Message      string    `json:"message"`
}

func (c *ClientOpts) GetEverything(everythingOpts EverythingOpts) (everythingResponse, error) {

	if len(c.ApiKey) == 0 {
		fmt.Println("Missing api key")
	}

	response, err := http.Get(everythingOpts.buildQuery(c))
	if err != nil {
		return everythingResponse{}, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return everythingResponse{}, err
	}

	everythingResponse := everythingResponse{}
	err = json.Unmarshal(body, &everythingResponse)

	if err != nil {
		return everythingResponse, err
	}

	if everythingResponse.Status == "error" {
		return everythingResponse, err
	}

	return everythingResponse, nil
}

func (e EverythingOpts) buildQuery(c *ClientOpts) string {
	query := baseUrl + "/everything?apiKey=" + c.ApiKey

	if e == (EverythingOpts{}) {
		return query
	}
	if len(e.Q) > 0 {
		query += "&q=" + e.Q
	}
	if len(e.QInTitle) > 0 {
		query += "&qInTitle=" + e.QInTitle
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
	if len(e.From) > 0 {
		query += "&from=" + e.From
	}
	if len(e.To) > 0 {
		query += "&to=" + e.To
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
