package newsapi

import (
	"testing"
)

func TestTopHeadlinesBuildQuery(t *testing.T) {
	headlines := TopHeadlinesOpts{}
	c := &ClientOpts{}

	result := headlines.buildQuery(c)
	expected := "https://newsapi.org/v2/top-headlines?apiKey="
	assertEqual(t, result, expected)

	headlines = TopHeadlinesOpts{
		Country:  "br",
		Category: "politics",
		Sources:  "cnn",
		Q:        "covid",
		PageSize: 20,
		Page:     1,
	}
	c = &ClientOpts{ApiKey: "qwerty123"}

	result = headlines.buildQuery(c)
	expected = "https://newsapi.org/v2/top-headlines?apiKey=qwerty123&country=br&category=politics&sources=cnn&q=covid&pageSize=20&page=1"
	assertEqual(t, result, expected)

	headlines = TopHeadlinesOpts{
		Country:  "us",
		Category: "",
		Sources:  "nytimes",
		Q:        "games",
		PageSize: 10,
		Page:     2,
	}
	c = &ClientOpts{ApiKey: ""}

	result = headlines.buildQuery(c)
	expected = "https://newsapi.org/v2/top-headlines?apiKey=&country=us&sources=nytimes&q=games&pageSize=10&page=2"

	assertEqual(t, result, expected)

	headlines = TopHeadlinesOpts{
		Country: "super invalid country",
	}
	c = &ClientOpts{ApiKey: "qwerty123"}

	result = headlines.buildQuery(c)
	expected = "https://newsapi.org/v2/top-headlines?apiKey=qwerty123&country=super invalid country"
	assertEqual(t, result, expected)
}
