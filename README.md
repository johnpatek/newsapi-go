# NewsAPI Go Client

[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](http://pkg.go.dev/github.com/johnpatek/newsapi-go)

[![codecov](https://codecov.io/gh/johnpatek/newsapi-go/branch/master/graph/badge.svg)](https://codecov.io/gh/johnpatek/newsapi-go)

Go client package for [NewsAPI](https://newsapi.org).

## Usage

The following [example](_example/main.go) demonstrates basic usage.

```go
package main

import (
	"fmt"
	"os"

	"github.com/johnpatek/newsapi-go"
	"github.com/kr/pretty"
)

func main() {
	if len(os.Args) > 1 {
		// get the top headlines about Lionel Messi from US based outlets
		messiHeadlines, err := newsapi.GetTopHeadlines(os.Args[1], newsapi.TopHeadlinesParameters{
			Q:        "messi",
			Category: newsapi.Sports,
			Country:  newsapi.USA,
		})
		if err != nil {
			fmt.Println(err)
			return
		}
		// print the top articles as a formatted JSON
		for _, headline := range messiHeadlines.Articles {
			pretty.Println(headline)
		}
	} else {
		fmt.Printf("Usage: example <API KEY>")
	}
}
```

This example can be built and run using the following commands:
```bash
make build
./bin/example <your API key>
```