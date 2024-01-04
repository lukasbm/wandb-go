# WandB-Go

Wrapping the public Weights &amp; Biases GraphQL API in Go

## First steps

1. Download the SDL schema definition file from the [W&B GraphQL API](https://api.wandb.ai/graphql) and save it as `schema.graphql`
    More info here: <https://stackoverflow.com/questions/37397886/get-graphql-whole-schema-query>
2. Run `go run genqlient --init` to generate the go definitions
3. 