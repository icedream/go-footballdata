package footballdata

import (
	"fmt"
	"time"
)

// DEPRECATED.
//
type SoccerSeasonFixturesRequest struct{ request }

// DEPRECATED.
//
// Matchday modifies the request to specify a match day.
func (r SoccerSeasonFixturesRequest) Matchday(matchday uint16) SoccerSeasonFixturesRequest {
	r.urlValues.Set("matchday", fmt.Sprintf("%d", matchday))
	return r
}

// DEPRECATED.
//
// TimeFrame modifies the request to specify a specific time frame.
func (r SoccerSeasonFixturesRequest) TimeFrame(timeframe time.Duration) SoccerSeasonFixturesRequest {
	r.urlValues.Set("timeFrame", durationToTimeFrame(timeframe))
	return r
}

// DEPRECATED.
//
// Do executes the request.
func (r SoccerSeasonFixturesRequest) Do() (s FixtureList, err error) {
	d, _, err := r.doJson("GET")
	if err != nil {
		return
	}

	err = d.Decode(&s)
	return
}

// DEPRECATED.
//
// FixturesOfSoccerSeason prepares a request to fetch the fixtures of a soccer season.
func (c *Client) FixturesOfSoccerSeason(soccerSeasonId uint64) SoccerSeasonFixturesRequest {
	return SoccerSeasonFixturesRequest{c.req("soccerseasons/%d/fixtures", soccerSeasonId)}
}
