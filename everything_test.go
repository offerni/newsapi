package newsapi

import (
	"testing"
)

func TestEverythingBuildQuery(t *testing.T) {

	everything := EverythingOpts{}
	c := &ClientOpts{}

	result := everything.buildQuery(c)
	expected := "https://newsapi.org/v2/everything?apiKey="
	assertEqual(t, result, expected)

	everything = EverythingOpts{
		Q:              "the oscars",
		QInTitle:       "streaming",
		Sources:        "rotten tomatoes",
		Domains:        "rottentomatoes.com",
		ExcludeDomains: "www.imdb.com",
		From:           "2020-05-31T00:34:00Z",
		To:             "2020-06-24T00:15:15Z",
		Language:       "en",
		SortBy:         "publishedAt",
		PageSize:       10,
		Page:           2,
	}
	c = &ClientOpts{ApiKey: "qwerty123"}

	result = everything.buildQuery(c)
	expected = "https://newsapi.org/v2/everything?apiKey=qwerty123&q=the oscars&" +
		"qInTitle=streaming&sources=rotten tomatoes&domains=rottentomatoes.com&excludeDomains=www.imdb.com&" +
		"from=2020-05-31T00:34:00Z&to=2020-06-24T00:15:15Z&language=en&sortBy=publishedAt&pageSize=10&page=2"
	assertEqual(t, result, expected)

	everything = EverythingOpts{
		Q:        "games",
		QInTitle: "overwatch",
		Sources:  "ign",
		Language: "en",
		SortBy:   "relevancy",
		PageSize: 5,
		Page:     1,
	}
	c = &ClientOpts{ApiKey: "qwerty123"}

	result = everything.buildQuery(c)
	expected = "https://newsapi.org/v2/everything?apiKey=qwerty123&q=games&" +
		"qInTitle=overwatch&sources=ign&language=en&sortBy=relevancy&pageSize=5&page=1"
	assertEqual(t, result, expected)

	everything = EverythingOpts{}
	c = &ClientOpts{ApiKey: ""}

	result = everything.buildQuery(c)
	expected = "https://newsapi.org/v2/everything?apiKey="
	assertEqual(t, result, expected)
}
