# Developer Guide

## Updating the schema definition

It is possible that the schema definition changes and queries no longer work as intended.
To fix this we need to download the SDL schema definition file from the [W&B GraphQL API](https://api.wandb.ai/graphql)
and save it as `schema.graphql`
More info here: <https://stackoverflow.com/questions/37397886/get-graphql-whole-schema-query>
You can also use a tool like Insomnia to download the schema definition file.

## Add new queries

If you take a look at genqlients config file (`genqlient.yaml`) you will see that it contains a list named operations.
It parses all `.graphql` files in the directory and generates Go code for each query.
To add a new query, simply create a new `.graphql` file and add the query to the operations list in `genqlient.yaml`.
Then run `go run github.com/Khan/genqlient` to generate the Go code for the new query.
