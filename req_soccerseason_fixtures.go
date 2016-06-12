package footballdata

import (
	"fmt"
	"time"
)

type SoccerSeasonFixturesRequest struct{ request }

// Modifies the request to specify a match day.
func (r SoccerSeasonFixturesRequest) Matchday(matchday uint16) SoccerSeasonFixturesRequest {
	r.v.Set("matchday", fmt.Sprintf("%i", matchday))
	return r
}

// Modifies the request to specify a specific time frame.
func (r SoccerSeasonFixturesRequest) TimeFrame(timeframe time.Duration) SoccerSeasonFixturesRequest {
	r.v.Set("timeFrame", durationToTimeFrame(timeframe))
	return r
}

// Executes the request.
func (r SoccerSeasonFixturesRequest) Do() (s SoccerSeason, err error) {
	d, _, err := r.doJson("GET")
	if err != nil {
		return
	}

	err = d.Decode(&s)
	return
}

// Prepares a request to fetch the fixtures of a soccer season.
func (c *Client) FixturesOfSoccerSeason(soccerSeasonId uint64) SoccerSeasonFixturesRequest {
	return SoccerSeasonFixturesRequest{c.req("soccerseasons/%i/fixtures", soccerSeasonId)}
}
