# WandB-Go

Wrapping the public Weights &amp; Biases GraphQL API in Go

## Setup

First, make sure you are signed in to Weights & Biases.
You can do this by using the python client and running `wandb login` in your terminal.
Or you can manually create a file called `.netrc` in your home directory and add the following lines to it:

```
machine api.wandb.ai
  login user
  password <your api key>
```

You can find your api key in your Weights & Biases settings or [here](https://wandb.ai/authorize).

Then, you can use the `wandb-go` package to query the Weights & Biases GraphQL API.

Once signed in we can generate the Go code for the GraphQL API by running `go run github.com/Khan/genqlient`.

## Usage

With everything in place we can start using it.
As the public GraphQL API is authorization protected, we need to set up the client to use the stored api key.
Take a look at `main.go` for this.

## Disclaimer

This is an unofficial wrapper around the Weights & Biases GraphQL API.
It is not affiliated with Weights & Biases in any way.

It is still work in progress and some queries might not work as expected.

## Roadmap

Most important features:

- init
- finish
- config
- log

Complete feature list:

- [ ] System metrics tracking
- [ ] Code Saving
- [ ] Storing system info in metadata
- [ ] resuming/reinitializing runs
- [ ] Main functions:
    - [ ] init
    - [ ] log
    - [ ] log_artifact
    - [ ] log_code
    - [ ] config
    - [ ] metadata:
        - [ ] summary
        - [ ] notes
        - [ ] tags
        - [ ] name
    - [ ] finish run (automatically?)
- [ ] offline mode
- [ ] connect to self-hosted servers

First try to mimic [javascript client](https://github.com/wandb/wandb-js). simpler and easier.
