package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/bradleyfalzon/ghinstallation/v2"
	"github.com/google/go-github/v44/github"
)

const (
	AppID             = 1
	InstallationID    = 99
	PrivateKeyFile    = "2022-05-17.private-key.pem"
	MinRequestSpacing = 5 * time.Second
)

func main() {
	ctx := context.Background()

	transport, err := ghinstallation.NewKeyFromFile(http.DefaultTransport, AppID, InstallationID, PrivateKeyFile)
	try(err)
	client := github.NewClient(&http.Client{Transport: transport})

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
			log.Printf("[Page %d] %d results in %s; sleeping %s\n", opt.Page, len(results.CodeResults), respDuration.Round(time.Millisecond), spacing.Round(time.Millisecond))
			time.Sleep(spacing)
		} else {
			log.Printf("[Page %d] %d results in %s; done\n", opt.Page, len(results.CodeResults), respDuration.Round(time.Millisecond))
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
