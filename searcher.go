package main

import (
	"context"
	"strings"

	"github.com/google/go-github/github"

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
func (searcher *Searcher) Run() ([]github.Issue, error) {
	queryComponents := make(map[string]string)
	queryComponents["repo"] = searcher.repo
	queryComponents["type"] = "pr"
	queryComponents["is"] = "open"

	buf := []string{}
	for k, v := range queryComponents {
		queryComponent := strings.Join([]string{k, v}, ":")
		buf = append(buf, queryComponent)
	}

	query := strings.Join(buf, " ")
	result, _, err := searcher.service.Issues(searcher.ctx, query, nil)

	return result.Issues, err
}
