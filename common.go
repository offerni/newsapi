package newsapi

import "testing"

// Shared between all endpoints
const baseUrl string = "https://newsapi.org/v2"

type article struct {
	Source      articleSource `json:"source"`
	Author      string        `json:"author"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Url         string        `json:"url"`
	UrlToImage  string        `json:"urlToImage"`
	PublishedAt string        `json:"publishedAt"`
	Content     string        `json:"content"`
}

type articleSource struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// Assertion functions for unit tests
func assertEqual(t *testing.T, result string, expected string) {
	if result != expected {
		t.Errorf("Incorrect Equals Assertion, expected: (%s), got: (%s)", expected, result)
	}
}
