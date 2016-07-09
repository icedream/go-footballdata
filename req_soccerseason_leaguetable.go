package footballdata

import "fmt"

type SoccerSeasonLeagueTableRequest struct{ request }

// DEPRECATED.
//
// Matchday modifies the request to specify a match day.
func (r SoccerSeasonLeagueTableRequest) Matchday(matchday uint16) SoccerSeasonLeagueTableRequest {
	r.urlValues.Set("matchday", fmt.Sprintf("%d", matchday))
	return r
}

// DEPRECATED.
//
// Do executes the request.
func (r SoccerSeasonLeagueTableRequest) Do() (s LeagueTable, err error) {
	d, _, err := r.doJson("GET")
	if err != nil {
		return
	}

	err = d.Decode(&s)
	return
}

// DEPRECATED.
//
// LeagueTableOfSoccerSeason prepares a new request to fetch the league table of a given soccer season.
func (c *Client) LeagueTableOfSoccerSeason(soccerSeasonId uint64) SoccerSeasonLeagueTableRequest {
	return SoccerSeasonLeagueTableRequest{c.req("soccerseasons/%d/leagueTable", soccerSeasonId)}
}
