package main

import (
	"fmt"
	"github.com/kbarber/puppetdb-client-go"
	"net/url"
)

func main() {
	server := puppetdb.NewServer("http://localhost:8080/")

	// A blank query string - just an example
	var values url.Values
	values = map[string][]string{
		"query": []string{"[\"=\",\"certname\",\"foobar\"]"},
	}
	queryString := values.Encode()

	response, _ := server.QueryEvents(queryString)
	fmt.Printf("Events: %v\n", response)
}
