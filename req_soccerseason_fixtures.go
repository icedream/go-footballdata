package footballdata

import (
	"fmt"
	"time"
)

type SoccerSeasonFixturesRequest struct{ request }

// Matchday modifies the request to specify a match day.
func (r SoccerSeasonFixturesRequest) Matchday(matchday uint16) SoccerSeasonFixturesRequest {
	r.urlValues.Set("matchday", fmt.Sprintf("%d", matchday))
	return r
}

// TimeFrame modifies the request to specify a specific time frame.
func (r SoccerSeasonFixturesRequest) TimeFrame(timeframe time.Duration) SoccerSeasonFixturesRequest {
	r.urlValues.Set("timeFrame", durationToTimeFrame(timeframe))
	return r
}

// Do executes the request.
func (r SoccerSeasonFixturesRequest) Do() (s FixtureList, err error) {
	d, _, err := r.doJson("GET")
	if err != nil {
		return
	}

	err = d.Decode(&s)
	return
}

func (c *client) FixturesOfSoccerSeason(soccerSeasonId uint64) SoccerSeasonFixturesRequest {
	return SoccerSeasonFixturesRequest{c.req("soccerseasons/%d/fixtures", soccerSeasonId)}
}
