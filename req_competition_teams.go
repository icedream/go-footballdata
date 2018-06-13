package footballdata

type CompetitionTeamsRequest struct{ request }

// Do executes the request.
func (r CompetitionTeamsRequest) Do() (s TeamList, err error) {
	d, _, err := r.doJson("GET")
	if err != nil {
		return
	}

	err = d.Decode(&s)
	return
}

// TeamsOfCompetition prepares a new request to fetch the league table of a given soccer season.
func (c *Client) TeamsOfCompetition(soccerSeasonId uint64) CompetitionTeamsRequest {
	req := CompetitionTeamsRequest{c.req("competitions/%d/teams", soccerSeasonId)}
	req.headers["X-Response-Control"] = "minified"
	return req
}
