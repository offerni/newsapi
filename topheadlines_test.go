package newsapi

import (
	"testing"
)

func TestTopHeadlinesBuildQuery(t *testing.T) {
	headlines := Headlines{}
	result := headlines.buildQuery()
	expected := "https://newsapi.org/v2/top-headlines?apiKey="

	assertEqual(t, result, expected)

	headlines = Headlines{
		Country:  "br",
		Category: "politics",
		Sources:  "cnn",
		Q:        "covid",
		PageSize: 20,
		Page:     1,
		ApiKey:   "123456789",
	}
	result = headlines.buildQuery()
	expected = "https://newsapi.org/v2/top-headlines?apiKey=123456789&country=br&category=politics&sources=cnn&q=covid&pageSize=20&page=1"

	assertEqual(t, result, expected)

	headlines = Headlines{
		Country:  "us",
		Category: "",
		Sources:  "nytimes",
		Q:        "games",
		PageSize: 10,
		Page:     2,
		ApiKey:   "",
	}
	result = headlines.buildQuery()
	expected = "https://newsapi.org/v2/top-headlines?apiKey=&country=us&sources=nytimes&q=games&pageSize=10&page=2"

	assertEqual(t, result, expected)

	headlines = Headlines{
		ApiKey:  "qwerty123",
		Country: "super invalid country",
	}
	result = headlines.buildQuery()
	expected = "https://newsapi.org/v2/top-headlines?apiKey=qwerty123&country=super invalid country"

	assertEqual(t, result, expected)
}
