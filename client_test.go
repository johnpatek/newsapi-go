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

func TestParams(t *testing.T) {
	values := url.Values{}
	setCategory(&values, Business)
	assert.NotEmpty(t, values.Encode())

	values = url.Values{}
	setCountry(&values, Japan)
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

func TestQuery(t *testing.T) {
	values := new(url.Values)
	queryString := values.Encode()
	assert.Empty(t, queryString)
}
