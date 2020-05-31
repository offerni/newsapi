package newsApiSdk

import (
	"encoding/json"
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

type HeadlineResponse struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []Article `json:"articles"`
	Code         string    `json:"code"`
	Message      string    `json:"message"`
}

type Article struct {
	Source      ArticleSource `json:"source"`
	Author      string        `json:"author"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Url         string        `json:"url"`
	UrlToImage  string        `json:"urlToImage"`
	PublishedAt string        `json:"publishedAt"`
	Content     string        `json:"content"`
}

type ArticleSource struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func GetTopHeadlines(headlines Headlines) (HeadlineResponse, error) {
	if len(headlines.ApiKey) == 0 {
		fmt.Println("Missing api key")
	}

	response, err := http.Get(headlines.buildQuery())
	if err != nil { // response error handling
		return HeadlineResponse{}, err
	}

	data, readErr := ioutil.ReadAll(response.Body)
	headlinesResponse := HeadlineResponse{}
	headlinesErr := json.Unmarshal(data, &headlinesResponse)

	if readErr != nil {
		return headlinesResponse, readErr
	}

	if headlinesErr != nil {
		return headlinesResponse, headlinesErr
	}

	if headlinesResponse.Status == "error" {
		return headlinesResponse, headlinesErr
	}

	return headlinesResponse, headlinesErr
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
