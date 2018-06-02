package footballdata

import "fmt"

type CompetitionLeagueTableRequest struct{ request }

// Matchday modifies the request to specify a match day.
func (r CompetitionLeagueTableRequest) Matchday(matchday uint16) CompetitionLeagueTableRequest {
	r.urlValues.Set("matchday", fmt.Sprintf("%d", matchday))
	return r
}

// Do executes the request.
func (r CompetitionLeagueTableRequest) Do() (s LeagueTable, err error) {
	d, _, err := r.doJson("GET")
	if err != nil {
		return
	}

	err = d.Decode(&s)
	return
}

// LeagueTableOfCompetition prepares a new request to fetch the league table of a given soccer season.
func (c *Client) LeagueTableOfCompetition(competitionId uint64) CompetitionLeagueTableRequest {
	return CompetitionLeagueTableRequest{c.req("competitions/%d/leagueTable", competitionId)}
}
