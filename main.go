package main

import (
	"context"
	"fmt"
	"github.com/Khan/genqlient/graphql"
	"net/http"
)

//go:generate go run github.com/Khan/genqlient

func main() {
	ctx := context.Background()
	client := graphql.NewClient("https://api.wandb.ai/graphql", http.DefaultClient)

	resp, err := FetchAProject(ctx, client)
	fmt.Println(resp.GetProject(), err)
}
