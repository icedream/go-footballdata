package footballdata

import (
	"strings"
	"time"
)

type FixturesRequest struct{ request }

// TimeFrame modifies the request to specify a specific relative time frame.
func (r FixturesRequest) TimeFrame(timeframe time.Duration) FixturesRequest {
	r.urlValues.Set("timeFrame", durationToTimeFrame(timeframe))
	return r
}

// TimeFrameStart modifies the request to specify the beginning of the time frame filter for the returned results.
//
// Only the year, month and day of the Time value will be used for the request.
func (r FixturesRequest) TimeFrameStart(date time.Time) FixturesRequest {
	r.urlValues.Set("timeFrameStart", date.Format(timeFrameLayout))
	return r
}

// TimeFrameEnd modifies the request to specify the end of the time frame filter for the returned results.
//
// Only the year, month and day of the Time value will be used for the request.
func (r FixturesRequest) TimeFrameStart(date time.Time) FixturesRequest {
	r.urlValues.Set("timeFrameEnd", date.Format(timeFrameLayout))
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
