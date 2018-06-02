package footballdata

type CompetitionMatchesRequest struct{ request }

// Do executes the request.
func (r CompetitionMatchesRequest) Do() (s FixtureList, err error) {
	d, _, err := r.doJson("GET")
	if err != nil {
		return
	}

	err = d.Decode(&s)
	return
}

// MatchesOfCompetition prepares a request to fetch the matches of a soccer season.
func (c *Client) MatchesOfCompetition(competitionId uint64) CompetitionMatchesRequest {
	return CompetitionMatchesRequest{c.req("competitions/%d/matches", competitionId)}
}
