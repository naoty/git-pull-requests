package main

import (
	"context"
	"fmt"
	"go-github/github"
	"strings"

	"golang.org/x/oauth2"
)

// Searcher searches pull requests.
type Searcher struct {
	repo    string
	ctx     context.Context
	service *github.SearchService
}

// NewSearcher generates a new Searcher with a given token.
func NewSearcher(repo, token string) *Searcher {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	return &Searcher{repo: repo, ctx: ctx, service: client.Search}
}

// Run searches pull requests.
func (searcher *Searcher) Run() error {
	queryComponents := make(map[string]string)
	queryComponents["repo"] = searcher.repo
	queryComponents["type"] = "pr"

	buf := []string{}
	for k, v := range queryComponents {
		queryComponent := strings.Join([]string{k, v}, ":")
		buf = append(buf, queryComponent)
	}

	query := strings.Join(buf, " ")
	result, _, err := searcher.service.Issues(searcher.ctx, query, nil)

	for _, issue := range result.Issues {
		fmt.Printf("%v\n", issue.GetTitle())
	}

	return err
}
