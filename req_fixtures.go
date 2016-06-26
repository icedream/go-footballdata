package footballdata

import (
	"strings"
	"time"
)

type FixturesRequest struct{ request }

// TimeFrame modifies the request to specify a specific time frame.
func (r FixturesRequest) TimeFrame(timeframe time.Duration) FixturesRequest {
	r.urlValues.Set("timeFrame", durationToTimeFrame(timeframe))
	return r
}

// League modifies the request to specify a list of leagues by their code.
func (r FixturesRequest) League(leagueCodes ...string) FixturesRequest {
	r.urlValues.Set("league", strings.Join(leagueCodes, ","))
	return r
}

// Do executes the request.
func (r FixturesRequest) Do() (s FixturesResponse, err error) {
	d, _, err := r.doJson("GET")
	if err != nil {
		return
	}

	err = d.Decode(&s)
	return
}

// Fixtures prepares a request to fetch the fixtures of a soccer season.
func (c *Client) Fixtures() FixturesRequest {
	return FixturesRequest{c.req("fixtures")}
}
