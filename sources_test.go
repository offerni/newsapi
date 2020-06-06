package newsapi

import (
	"testing"
)

func TestSourcesBuildQuery(t *testing.T) {
	sources := SourcesOpts{}
	c := &ClientOpts{ApiKey: ""}

	result := sources.buildQuery(c)
	expected := "https://newsapi.org/v2/sources?apiKey="

	assertEqual(t, result, expected)

	sources = SourcesOpts{
		Category: "entertainment",
		Language: "pt",
		Country:  "br",
	}
	c = &ClientOpts{ApiKey: "qwerty123"}

	result = sources.buildQuery(c)
	expected = "https://newsapi.org/v2/sources?apiKey=qwerty123&category=entertainment&language=pt&country=br"

	assertEqual(t, result, expected)

	sources = SourcesOpts{
		Category: "sports",
		Language: "en",
		Country:  "us",
	}
	c = &ClientOpts{ApiKey: ""}

	result = sources.buildQuery(c)
	expected = "https://newsapi.org/v2/sources?apiKey=&category=sports&language=en&country=us"

	assertEqual(t, result, expected)
}
