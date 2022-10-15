# Advent of Code in Go

This is set up as a Visual Studio Code devcontainer, to make it easy to set up the
Go environment.

## Use

1. Clone repo
1. Open in Visual Studio code
1. Allow it to establish a 'Dev Container' using the Dev Container extension (which you'll be prompted to install if not already present)
1. It'll take a few minutes to create the container.
1. Bring up a terminal in the container (Terminal->New Container)
1. Run `go run test.go`
1. You should see a message confirming that your Go container is working

# Working on Advent of Code challenges

## Setup

1. Log into Advent of Code in a browser and extract the session cookie.
1. Place it in the `.env` file, based on `.env.sample`.

## To solve a new day

* Inside the dev container, run `scrape.sh <year> <day>`
  * This will fetch the assignment details and input, into `<year>/<day>`
