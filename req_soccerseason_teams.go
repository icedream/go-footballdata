package footballdata

// DEPRECATED.
//
type SoccerSeasonTeamsRequest struct{ request }

// DEPRECATED.
//
// Do executes the request.
func (r SoccerSeasonTeamsRequest) Do() (s TeamList, err error) {
	d, _, err := r.doJson("GET")
	if err != nil {
		return
	}

	err = d.Decode(&s)
	return
}

// DEPRECATED.
//
// TeamsOfSoccerSeason prepares a new request to fetch the league table of a given soccer season.
func (c *Client) TeamsOfSoccerSeason(soccerSeasonId uint64) SoccerSeasonTeamsRequest {
	return SoccerSeasonTeamsRequest{c.req("soccerseasons/%d/teams", soccerSeasonId)}
}
