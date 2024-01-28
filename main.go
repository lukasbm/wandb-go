package main

import (
	"context"
	"fmt"
	"github.com/Khan/genqlient/graphql"
	"net/http"
)

//go:generate go run github.com/Khan/genqlient

const ENDPOINT = "https://api.wandb.ai/graphql"

type authedTransport struct {
	wrapped http.RoundTripper
}

func (t *authedTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	AddCredentials(req) // adds credentials from .netrc
	return t.wrapped.RoundTrip(req)
}

func main() {
	ctx := context.Background()
	client := graphql.NewClient(
		ENDPOINT,
		&http.Client{
			Transport: &authedTransport{http.DefaultTransport},
		},
	)

	// actually make a request
	resp, err := Viewer(ctx, client)   // fetches viewer information
	fmt.Println(resp.GetViewer(), err) // and shows it
}
