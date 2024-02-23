package newsapi

import (
	"net/url"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	// this key will return an error from the API
	invalidKey string = "00000000000000000000000000000000"
)

var validKey string = os.Getenv("NEWS_API_KEY")

func TestGetEverything(t *testing.T) {
	_, err := GetEverything(invalidKey, EverythingParameters{})
	assert.Error(t, err)
	_, err = GetEverything(validKey, EverythingParameters{
		Q: "golang",
	})
	assert.NoError(t, err)
}

func TestGetTopHeadlines(t *testing.T) {
	_, err := GetTopHeadlines(invalidKey, TopHeadlinesParameters{})
	assert.Error(t, err)
	_, err = GetTopHeadlines(validKey, TopHeadlinesParameters{
		Q: "golang",
	})
	assert.NoError(t, err)
}

func TestGetSources(t *testing.T) {
	_, err := GetSources(invalidKey, SourcesParameters{})
	assert.Error(t, err)
	_, err = GetSources(validKey, SourcesParameters{})
	assert.NoError(t, err)
}

func TestQuery(t *testing.T) {
	values := new(url.Values)
	queryString := values.Encode()
	assert.Empty(t, queryString)
}
