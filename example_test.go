package footballdata_test

import (
	"fmt"
	"net/http"

	"github.com/icedream/go-footballdata"
)

func Example() {
	// Create client
	client := footballdata.NewClient(http.DefaultClient)

	// Tell it to use our API token
	client.SetToken("<insert your api token here>")

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

func ExampleClient() {
	// Create client
	client := footballdata.NewClient(http.DefaultClient)

	/* Do something with the client instance... */
	seasons, err := client.SoccerSeasons().Do()
	if err != nil {
		panic(err)
	}
	for _, season := range seasons {
		fmt.Println(season.Id, season.Caption)
	}
}

func ExampleClient_withToken() {
	// Create client
	client := footballdata.NewClient(http.DefaultClient)

	// Tell it to use our API token
	client.SetToken("<insert your api token here>")
}
