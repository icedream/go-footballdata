package footballdata

import "fmt"

type SoccerSeasonLeagueTableRequest struct{ request }

// Modifies the request to specify a match day.
func (r SoccerSeasonLeagueTableRequest) Matchday(matchday uint16) SoccerSeasonLeagueTableRequest {
	r.v.Set("matchday", fmt.Sprintf("%i", matchday))
	return r
}

// Executes the request.
func (r SoccerSeasonLeagueTableRequest) Do() (s SoccerSeason, err error) {
	d, _, err := r.doJson("GET")
	if err != nil {
		return
	}

	err = d.Decode(&s)
	return
}

// Prepares a new request to fetch the league table of a given soccer season.
func (c *Client) LeagueTableOfSoccerSeason(soccerSeasonId uint64) SoccerSeasonLeagueTableRequest {
	return SoccerSeasonLeagueTableRequest{c.req("soccerseasons/%i/leagueTable", soccerSeasonId)}
}
