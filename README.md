# Football-Data API for Golang

[![Build Status](https://travis-ci.org/icedream/go-footballdata.svg?branch=master)](https://travis-ci.org/icedream/go-footballdata)
[![GoDoc](https://godoc.org/github.com/icedream/go-footballdata?status.svg)](https://godoc.org/github.com/icedream/go-footballdata)

This library provides functionality to communicate with the API provided by http://football-api.org. This way programs can use data provided by the API in order to show information about football/soccer games from various seasons for free.

## How to use this library?

Before you use this library please [register for a free API key](http://api.football-data.org/register) in order to increase your usage limits. The library also works without an API key.

You can install this library by running:

	go get github.com/icedream/go-footballdata

Afterwards you can use this library like this:

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/icedream/go-footballdata"
)

func main() {
	// Create client (optionally with auth token)
	client := new(footballdata.Client).
		WithToken("<insert your api token here>")

	// Get list of seasons...
	seasons, err := client.SoccerSeasons().Do()
	if err != nil {
		panic(err)
	}

	// ...and print them
	for _, season := range seasons {
		fmt.Println(season.Id, season.Caption)
	}
}

```
