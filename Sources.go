package newsapi

import (
	"encoding/json"
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

type sourcesResponse struct {
	Status         string         `json:"status"`
	SourceResponse []sourceStruct `json:"sources"`
	Code           string         `json:"code"`
	Message        string         `json:"message"`
}

type sourceStruct struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Url         string `json:"url"`
	Category    string `json:"category"`
	Language    string `json:"language"`
	Country     string `json:"country"`
}

func GetSources(sources Sources) (sourcesResponse, error) {

	if len(sources.ApiKey) == 0 {
		fmt.Println("Missing api key")
	}

	response, err := http.Get(sources.buildQuery())
	if err != nil {
		return sourcesResponse{}, err
	}

	body, readErr := ioutil.ReadAll(response.Body)

	if readErr != nil {
		return sourcesResponse{}, readErr
	}

	sourcesResponse := sourcesResponse{}
	sourcesErr := json.Unmarshal(body, &sourcesResponse)

	if sourcesErr != nil {
		return sourcesResponse, sourcesErr
	}

	return sourcesResponse, sourcesErr
}

func (s Sources) buildQuery() string {
	query := baseUrl + "/sources?apiKey=" + s.ApiKey

	if (s == Sources{}) {
		return query
	}
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
