package newsapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type TopHeadlinesOpts struct {
	Country  string
	Category string
	Sources  string
	Q        string
	PageSize int
	Page     int
}

type headlinesResponse struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []article `json:"articles"`
	Code         string    `json:"code"`
	Message      string    `json:"message"`
}

func (c *ClientOpts) GetTopHeadlines(topHeadlinesOpts TopHeadlinesOpts) (headlinesResponse, error) {

	if len(c.ApiKey) == 0 {
		fmt.Println("Missing api key")
	}

	response, err := http.Get(topHeadlinesOpts.buildQuery(c))
	if err != nil { // response error handling
		return headlinesResponse{}, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return headlinesResponse{}, err
	}

	headlinesResponse := headlinesResponse{}
	err = json.Unmarshal(body, &headlinesResponse)

	if err != nil {
		return headlinesResponse, err
	}

	if headlinesResponse.Status == "error" {
		return headlinesResponse, err
	}

	return headlinesResponse, nil
}

func (t TopHeadlinesOpts) buildQuery(c *ClientOpts) string {
	query := baseUrl + "/top-headlines?apiKey=" + c.ApiKey

	if t == (TopHeadlinesOpts{}) {
		return query
	}
	if len(t.Country) > 0 {
		query += "&country=" + t.Country
	}
	if len(t.Category) > 0 {
		query += "&category=" + t.Category
	}
	if len(t.Sources) > 0 {
		query += "&sources=" + t.Sources
	}
	if len(t.Q) > 0 {
		query += "&q=" + t.Q
	}
	if t.PageSize > 0 {
		query += "&pageSize=" + strconv.Itoa(t.PageSize)
	}
	if t.Page > 0 {
		query += "&page=" + strconv.Itoa(t.Page)
	}

	return query
}
