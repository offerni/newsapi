package newsapi

import (
	"testing"
)

func TestSourcesBuildQuery(t *testing.T) {
	sources := Sources{}
	result := sources.buildQuery()
	expected := "https://newsapi.org/v2/sources?apiKey="

	assertEqual(t, result, expected)

	sources = Sources{
		Category: "entertainment",
		Language: "pt",
		Country:  "br",
		ApiKey:   "qwerty123",
	}
	result = sources.buildQuery()
	expected = "https://newsapi.org/v2/sources?apiKey=qwerty123&category=entertainment&language=pt&country=br"

	assertEqual(t, result, expected)

	sources = Sources{
		Category: "sports",
		Language: "en",
		Country:  "us",
	}
	result = sources.buildQuery()
	expected = "https://newsapi.org/v2/sources?apiKey=&category=sports&language=en&country=us"

	assertEqual(t, result, expected)
}
