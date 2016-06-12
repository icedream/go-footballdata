package footballdata

import (
	"fmt"
	"time"
)

type TeamFixturesRequest struct{ request }

// Modifies the request to specify a specific time frame.
func (r TeamFixturesRequest) TimeFrame(timeframe time.Duration) TeamFixturesRequest {
	r.v.Set("timeFrame", durationToTimeFrame(timeframe))
	return r
}

// Modifies the request to specify a list of leagues by their code.
func (r TeamFixturesRequest) Season(season uint64) TeamFixturesRequest {
	r.v.Set("season", fmt.Sprintf("%i", season))
	return r
}

// Modifies the request to specify a venue.
func (r TeamFixturesRequest) Venue(venue Venue) TeamFixturesRequest {
	r.v.Set("venue", string(venue))
	return r
}

// Executes the request.
func (r TeamFixturesRequest) Do() (s FixturesResponse, err error) {
	d, _, err := r.doJson("GET")
	if err != nil {
		return
	}

	err = d.Decode(&s)
	return
}

// Prepares a request to fetch the fixtures of a soccer season.
func (c *Client) FixturesOfTeam(id uint64) TeamFixturesRequest {
	return TeamFixturesRequest{c.req("teams/%i/fixtures", id)}
}
