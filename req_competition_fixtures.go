package footballdata

import (
	"fmt"
	"time"
)

type CompetitionFixturesRequest struct{ request }

// Matchday modifies the request to specify a match day.
func (r CompetitionFixturesRequest) Matchday(matchday uint16) CompetitionFixturesRequest {
	r.urlValues.Set("matchday", fmt.Sprintf("%d", matchday))
	return r
}

// TimeFrame modifies the request to specify a specific time frame.
func (r CompetitionFixturesRequest) TimeFrame(timeframe time.Duration) CompetitionFixturesRequest {
	r.urlValues.Set("timeFrame", durationToTimeFrame(timeframe))
	return r
}

// Do executes the request.
func (r CompetitionFixturesRequest) Do() (s FixtureList, err error) {
	d, _, err := r.doJson("GET")
	if err != nil {
		return
	}

	err = d.Decode(&s)
	return
}

// FixturesOfCompetition prepares a request to fetch the fixtures of a soccer season.
func (c *Client) FixturesOfCompetition(soccerSeasonId uint64) CompetitionFixturesRequest {
	return CompetitionFixturesRequest{c.req("competitions/%d/fixtures", soccerSeasonId)}
}
