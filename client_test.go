// Copyright (c) 2024 John R Patek Sr
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package newsapi

import (
	"errors"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	// this key will return an error from the API
	invalidKey string = "00000000000000000000000000000000"

	invalidRootURI string = "invalidrooturi"
)

var validKey string = os.Getenv("NEWS_API_KEY")

func TestParams(t *testing.T) {
	values := url.Values{}
	setCategory(&values, Business)
	assert.NotEmpty(t, values.Encode())

	values = url.Values{}
	setCountry(&values, Japan)
	assert.NotEmpty(t, values.Encode())

	values = url.Values{}
	setDomains(&values, []string{"espn.com", "clickhole.com"})
	assert.NotEmpty(t, values.Encode())

	values = url.Values{}
	setExcludeDomains(&values, []string{"espn.com", "clickhole.com"})
	assert.NotEmpty(t, values.Encode())

	values = url.Values{}
	setFrom(&values, time.Now())
	assert.NotEmpty(t, values.Encode())

	values = url.Values{}
	setLanguage(&values, Swedish)
	assert.NotEmpty(t, values.Encode())

	values = url.Values{}
	setPage(&values, 3)
	assert.NotEmpty(t, values.Encode())

	values = url.Values{}
	setPageSize(&values, 2)
	assert.NotEmpty(t, values.Encode())

	values = url.Values{}
	setQ(&values, "golang")
	assert.NotEmpty(t, values.Encode())

	values = url.Values{}
	setSearchIn(&values, Title)
	assert.NotEmpty(t, values.Encode())

	values = url.Values{}
	setSortBy(&values, Popularity)
	assert.NotEmpty(t, values.Encode())

	values = url.Values{}
	setSources(&values, []string{"abc-news", "abc-news-au"})
	assert.NotEmpty(t, values.Encode())

	values = url.Values{}
	setTo(&values, time.Now())
	assert.NotEmpty(t, values.Encode())
}

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

type invalidResponseType struct {
	Channel chan int `json:"invalid"`
}

func (invalid *invalidResponseType) UnmarshalJSON(data []byte) error {
	return errors.New("unmarshal error")
}

func TestBadRequest(t *testing.T) {
	response, err := request[invalidResponseType](validKey, everythingEndpoint, EverythingParameters{
		Q: "golang",
	})
	assert.Nil(t, response)
	assert.Error(t, err)
	rootURI = invalidRootURI
	_, err = request[EverythingResponse](validKey, everythingEndpoint, EverythingParameters{
		Q: "golang",
	})
	assert.Error(t, err)
}
