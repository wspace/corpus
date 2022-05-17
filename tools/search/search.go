package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/go-github/v44/github"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/endpoints"
)

const MinRequestSpacing = 5 * time.Second

func main() {
	ctx := context.Background()

	httpClient, err := authenticate(ctx, "GITHUB_CLIENT_ID", "GITHUB_CLIENT_SECRET")
	try(err)
	client := github.NewClient(httpClient)

	opt := &github.SearchOptions{
		Sort:  "indexed",
		Order: "desc",
	}
	results, err := searchPaged(ctx, client, opt, func(search *github.SearchService, opt *github.SearchOptions) (*github.CodeSearchResult, *github.Response, error) {
		return search.Code(ctx, "filename:fact.ws extension:ws", opt)
	})
	log.Printf("Got %d results\n", len(results))
	try(err)
}

func authenticate(ctx context.Context, clientID, clientSecret string) (*http.Client, error) {
	conf := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     endpoints.GitHub,
		Scopes:       []string{},
	}

	// Redirect user to consent page to ask for permission for the scopes
	// specified above.
	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	fmt.Printf("Visit the URL to authorize:\n%v\n", url)

	// Use the authorization code that is pushed to the redirect URL. Exchange
	// will do the handshake to retrieve the initial access token. The HTTP Client
	// returned by conf.Client will refresh the token as necessary.
	fmt.Print("Input code from URL query parameters: ")
	var code string
	if _, err := fmt.Scan(&code); err != nil {
		return nil, err
	}
	tok, err := conf.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}

	return conf.Client(ctx, tok), nil
}

func searchPaged(ctx context.Context, client *github.Client, opt *github.SearchOptions,
	search func(*github.SearchService, *github.SearchOptions) (*github.CodeSearchResult, *github.Response, error)) ([]*github.CodeResult, error) {
	opt.ListOptions = github.ListOptions{
		Page:    1,
		PerPage: 10, // 100 is max
	}
	var allResults []*github.CodeResult
	for {
		before := time.Now()
		results, resp, err := search(client.Search, opt)
		if err != nil {
			return allResults, err
		}
		respDuration := time.Since(before)
		allResults = append(allResults, results.CodeResults...)

		if resp.NextPage != 0 {
			spacing := resp.Rate.Reset.Sub(before)/time.Duration(resp.Rate.Remaining+1) - respDuration
			if spacing < MinRequestSpacing {
				spacing = MinRequestSpacing
			}
			log.Printf("[Page %d] %d results in %s; sleeping %s\n", opt.Page, len(results.CodeResults), respDuration, spacing)
			time.Sleep(spacing)
		} else {
			log.Printf("[Page %d] %d results in %s; done\n", opt.Page, len(results.CodeResults), respDuration)
			return allResults, nil
		}
		opt.Page = resp.NextPage
	}
}

func try(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
