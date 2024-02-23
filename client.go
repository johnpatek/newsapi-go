package newsapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type SearchInType string

type CategoryType string

type LanguageType string

type CountryType string

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

	rootURI              string = "https://newsapi.org/v2"
	everythingEndpoint   string = "/everything"
	topHeadlinesEndpoint string = "/top-headlines"
	sourcesEndpoint      string = "/top-headlines/sources"
)

type ArticleSource struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

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

type Source struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Category    string `json:"category"`
	Language    string `json:"language"`
	Country     string `json:"country"`
}

type Parameters interface {
	toString() string
}

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

type TopHeadlinesParameters struct {
	Country  CountryType
	Category CategoryType
	Sources  []string
	Q        string
	PageSize int
	Page     int
}

type SourcesParameters struct {
	Category CategoryType
	Language LanguageType
	Country  CountryType
}

type EverythingResponse struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []Article `json:"Articles"`
}

type TopHeadlinesResponse struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []Article `json:"Articles"`
}

type SourcesResponse struct {
	Status  string   `json:"status"`
	Sources []Source `json:"sources"`
}

type ErrorResponse struct {
	Status  string `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (params EverythingParameters) toString() string {
	values := url.Values{}

	if params.Q != "" {
		values.Add("q", params.Q)
	}

	if params.SearchIn != SearchInDefault {
		values.Add("searchIn", fmt.Sprintf("%v", params.SearchIn))
	}

	if len(params.Sources) > 0 {
		sources := strings.Join(params.Sources, ",")
		values.Add("sources", sources)
	}

	if len(params.Domains) > 0 {
		domains := strings.Join(params.Domains, ",")
		values.Add("domains", domains)
	}

	if !params.From.IsZero() {
		values.Add("from", params.From.Format(time.RFC3339))
	}

	if !params.To.IsZero() {
		values.Add("to", params.To.Format(time.RFC3339))
	}

	if params.Language != AllLanguages {
		values.Add("language", fmt.Sprintf("%v", params.Language))
	}

	if params.SortBy != SortByDefault {
		values.Add("sortBy", fmt.Sprintf("%v", params.SortBy))
	}

	if params.PageSize != 0 {
		values.Add("pageSize", fmt.Sprintf("%d", params.PageSize))
	}

	if params.Page != 0 {
		values.Add("page", fmt.Sprintf("%d", params.Page))
	}

	return values.Encode()
}

func (params TopHeadlinesParameters) toString() string {
	values := url.Values{}

	setCountry(&values, params.Country)

	setCategory(&values, params.Category)

	if len(params.Sources) > 0 {
		sources := strings.Join(params.Sources, ",")
		values.Add("sources", sources)
	}

	if params.Q != "" {
		values.Add("q", params.Q)
	}

	if params.PageSize != 0 {
		values.Add("pageSize", fmt.Sprintf("%d", params.PageSize))
	}

	if params.Page != 0 {
		values.Add("page", fmt.Sprintf("%d", params.Page))
	}

	return values.Encode()
}

func (params SourcesParameters) toString() string {
	values := url.Values{}

	if params.Category != AllCategories {
		values.Add("category", fmt.Sprintf("%v", params.Category))
	}

	if params.Language != AllLanguages {
		values.Add("language", fmt.Sprintf("%v", params.Language))
	}

	if params.Country != AllCountries {
		values.Add("country", fmt.Sprintf("%v", params.Country))
	}

	return values.Encode()
}

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

func request[ResponseType any](apiKey, endpoint string, params Parameters) (*ResponseType, error) {
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
