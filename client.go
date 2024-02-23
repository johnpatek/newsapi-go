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
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// The fields to search for keywords and phrases.
type SearchInType string

// The category to get headlines for. Cannot be mixed with Sources parameter.
type CategoryType string

// The 2-letter ISO-639-1 code of languages to get headlines for.
type LanguageType string

// The 2-letter ISO 3166-1 code of countries to get headlines for.
type CountryType string

// The order to sort the articles in.
type SortByType string

const (
	AllCategories CategoryType = ""
	Business      CategoryType = "business"
	Entertainment CategoryType = "entertainment"
	General       CategoryType = "general"
	Health        CategoryType = "health"
	Science       CategoryType = "science"
	Sports        CategoryType = "sports"
	Technology    CategoryType = "technology"

	AllLanguages LanguageType = ""
	Arabic       LanguageType = "ar"
	German       LanguageType = "de"
	English      LanguageType = "en"
	Spanish      LanguageType = "es"
	French       LanguageType = "fr"
	Hebrew       LanguageType = "he"
	Italian      LanguageType = "it"
	Dutch        LanguageType = "nl"
	Norwegian    LanguageType = "no"
	Portuguese   LanguageType = "pt"
	Russian      LanguageType = "ru"
	Swedish      LanguageType = "sv"
	Undefined    LanguageType = "ud"
	Chinese      LanguageType = "zh"

	AllCountries CountryType = ""
	UAE          CountryType = "ae"
	Argentina    CountryType = "ar"
	Austria      CountryType = "at"
	Australia    CountryType = "au"
	Belgium      CountryType = "be"
	Bulgaria     CountryType = "bg"
	Brazil       CountryType = "br"
	Canada       CountryType = "ca"
	Switzerland  CountryType = "ch"
	China        CountryType = "cn"
	Colombia     CountryType = "co"
	Cuba         CountryType = "cu"
	Czechia      CountryType = "cz"
	Germany      CountryType = "de"
	Egypt        CountryType = "eg"
	France       CountryType = "fr"
	UK           CountryType = "gb"
	Greece       CountryType = "gr"
	HongKong     CountryType = "hk"
	Hungary      CountryType = "hu"
	Indonesia    CountryType = "id"
	Ireland      CountryType = "ie"
	Israel       CountryType = "il"
	India        CountryType = "in"
	Italy        CountryType = "it"
	Japan        CountryType = "jp"
	SouthKorea   CountryType = "kr"
	Lithuania    CountryType = "lt"
	Latvia       CountryType = "lv"
	Morocco      CountryType = "ma"
	Mexico       CountryType = "mx"
	Malaysia     CountryType = "my"
	Nigeria      CountryType = "ng"
	Netherlands  CountryType = "nl"
	Norway       CountryType = "no"
	NewZealand   CountryType = "nz"
	Philippines  CountryType = "ph"
	Poland       CountryType = "pl"
	Portugal     CountryType = "pt"
	Romania      CountryType = "ro"
	Serbia       CountryType = "rs"
	Russia       CountryType = "ru"
	SaudiaArabia CountryType = "sa"
	Sweden       CountryType = "se"
	Singapore    CountryType = "sg"
	Slovenia     CountryType = "si"
	Slovakia     CountryType = "sk"
	Thailand     CountryType = "th"
	Turkey       CountryType = "tr"
	Taiwan       CountryType = "tw"
	Ukraine      CountryType = "ua"
	USA          CountryType = "us"
	Venezuela    CountryType = "ve"
	SouthAfrica  CountryType = "za"

	SearchInDefault SearchInType = ""
	Title           SearchInType = "title"
	Description     SearchInType = "description"
	Content         SearchInType = "content"

	SortByDefault SortByType = ""
	Relevancy     SortByType = "relevancy"
	Popularity    SortByType = "popularity"
	PublishedAt   SortByType = "publishedAt"

	everythingEndpoint   string = "/everything"
	topHeadlinesEndpoint string = "/top-headlines"
	sourcesEndpoint      string = "/top-headlines/sources"
)

var rootURI string = "https://newsapi.org/v2"

// ArticleSource contains the identifier id and a display name for the source an article came from.
type ArticleSource struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Article represents an entry in the list of results from a search
type Article struct {
	Source      ArticleSource `json:"source"`
	Author      string        `json:"author"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	URL         string        `json:"url"`
	ImageURL    string        `json:"urlToImage"`
	PublishedAt string        `json:"publishedAt"`
	Content     string        `json:"content"`
}

// Source contains information for a single news publisher
type Source struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Category    string `json:"category"`
	Language    string `json:"language"`
	Country     string `json:"country"`
}

type parameters interface {
	toString() string
}

// EverythingParameters provides a URL query for the /everything endpoint
type EverythingParameters struct {
	Q              string
	SearchIn       SearchInType
	Sources        []string
	Domains        []string
	ExcludeDomains []string
	From           time.Time
	To             time.Time
	Language       LanguageType
	SortBy         SortByType
	PageSize       int
	Page           int
}

// TopHeadlinesParameters provides a URL query for the /top-headlines endpoint
type TopHeadlinesParameters struct {
	Country  CountryType
	Category CategoryType
	Sources  []string
	Q        string
	PageSize int
	Page     int
}

// SourcesParameters provides a URL query for the /top-headlines/sources endpoint
type SourcesParameters struct {
	Category CategoryType
	Language LanguageType
	Country  CountryType
}

// EverythingResponse contains information from calls to the /everything endpoint
type EverythingResponse struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []Article `json:"Articles"`
}

// TopHeadlinesResponse contains information from calls to the /top-headlines endpoint
type TopHeadlinesResponse struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []Article `json:"Articles"`
}

// SourcesResponse contains information from calls to the /top-headlines/sources endpoint
type SourcesResponse struct {
	Status  string   `json:"status"`
	Sources []Source `json:"sources"`
}

// ErrorResponse contains information from invalid API calls
type ErrorResponse struct {
	Status  string `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (params EverythingParameters) toString() string {
	values := url.Values{}
	setQ(&values, params.Q)
	setSearchIn(&values, params.SearchIn)
	setSources(&values, params.Sources)
	setDomains(&values, params.Domains)
	setExcludeDomains(&values, params.ExcludeDomains)
	setFrom(&values, params.From)
	setTo(&values, params.To)
	setLanguage(&values, params.Language)
	setSortBy(&values, params.SortBy)
	setPageSize(&values, params.PageSize)
	setPage(&values, params.Page)
	return values.Encode()
}

func (params TopHeadlinesParameters) toString() string {
	values := url.Values{}
	setCountry(&values, params.Country)
	setCategory(&values, params.Category)
	setSources(&values, params.Sources)
	setQ(&values, params.Q)
	setPageSize(&values, params.PageSize)
	setPage(&values, params.Page)
	return values.Encode()
}

func (params SourcesParameters) toString() string {
	values := url.Values{}
	setCategory(&values, params.Category)
	setLanguage(&values, params.Language)
	setCountry(&values, params.Country)
	return values.Encode()
}

// GetEverything requests the /everything endpoint to retrieve entries that match a specific criteria
func GetEverything(apiKey string, params EverythingParameters) (*EverythingResponse, error) {
	response, err := request[EverythingResponse](apiKey, everythingEndpoint, params)
	if err != nil {
		return nil, fmt.Errorf("newsapi.GetEverything: %v", err)
	}
	return response, nil
}

func GetTopHeadlines(apiKey string, params TopHeadlinesParameters) (*TopHeadlinesResponse, error) {
	response, err := request[TopHeadlinesResponse](apiKey, topHeadlinesEndpoint, params)
	if err != nil {
		return nil, fmt.Errorf("newsapi.GetTopHeadlines: %v", err)
	}
	return response, nil
}

func GetSources(apiKey string, params SourcesParameters) (*SourcesResponse, error) {
	response, err := request[SourcesResponse](apiKey, sourcesEndpoint, params)
	if err != nil {
		return nil, fmt.Errorf("newsapi.GetSources: %v", err)
	}
	return response, nil
}

func request[ResponseType any](apiKey, endpoint string, params parameters) (*ResponseType, error) {
	responseBody := new(ResponseType)

	paramsString := ""

	url := fmt.Sprintf("%s%s", rootURI, endpoint)

	paramsString = params.toString()

	if paramsString != "" {
		url = fmt.Sprintf("%s?%s", url, paramsString)
	}

	request, _ := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Add("X-Api-Key", apiKey)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	} else if response.StatusCode != http.StatusOK {
		errorDecoder := json.NewDecoder(response.Body)
		errorResponse := ErrorResponse{}
		_ = errorDecoder.Decode(&errorResponse)
		return nil, fmt.Errorf("%v", errorResponse)
	}
	responseDecoder := json.NewDecoder(response.Body)
	err = responseDecoder.Decode(responseBody)
	if err != nil {
		return nil, err
	}
	return responseBody, nil
}

func setCategory(values *url.Values, category CategoryType) {
	if category != AllCategories {
		values.Add("category", fmt.Sprintf("%v", category))
	}
}

func setCountry(values *url.Values, country CountryType) {
	if country != AllCountries {
		values.Add("country", fmt.Sprintf("%v", country))
	}
}

func setDomains(values *url.Values, domains []string) {
	if len(domains) > 0 {
		values.Add("domains", strings.Join(domains, ","))
	}
}

func setExcludeDomains(values *url.Values, excludeDomains []string) {
	if len(excludeDomains) > 0 {
		values.Add("excludeDomains", strings.Join(excludeDomains, ","))
	}
}

func setFrom(values *url.Values, from time.Time) {
	if !from.IsZero() {
		values.Add("from", from.Format(time.RFC3339))
	}
}

func setLanguage(values *url.Values, language LanguageType) {
	if language != AllLanguages {
		values.Add("language", fmt.Sprintf("%v", language))
	}
}

func setPage(values *url.Values, page int) {
	if page != 0 {
		values.Add("page", fmt.Sprintf("%d", page))
	}
}

func setPageSize(values *url.Values, pageSize int) {
	if pageSize != 0 {
		values.Add("page", fmt.Sprintf("%d", pageSize))
	}
}

func setQ(values *url.Values, q string) {
	if q != "" {
		values.Add("q", q)
	}
}

func setSearchIn(values *url.Values, searchIn SearchInType) {
	if searchIn != SearchInDefault {
		values.Add("searchIn", fmt.Sprintf("%v", searchIn))
	}
}

func setSortBy(values *url.Values, sortBy SortByType) {
	if sortBy != SortByDefault {
		values.Add("sortBy", fmt.Sprintf("%v", sortBy))
	}
}

func setSources(values *url.Values, sources []string) {
	if len(sources) > 0 {
		values.Add("sources", strings.Join(sources, ","))
	}
}

func setTo(values *url.Values, to time.Time) {
	if !to.IsZero() {
		values.Add("to", to.Format(time.RFC3339))
	}
}
