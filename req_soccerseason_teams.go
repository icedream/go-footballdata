package footballdata

type SoccerSeasonTeamsRequest struct{ request }

// Do executes the request.
func (r SoccerSeasonTeamsRequest) Do() (s TeamList, err error) {
	d, _, err := r.doJson("GET")
	if err != nil {
		return
	}

	err = d.Decode(&s)
	return
}

func (c *client) TeamsOfSoccerSeason(soccerSeasonId uint64) SoccerSeasonTeamsRequest {
	return SoccerSeasonTeamsRequest{c.req("soccerseasons/%d/leagueTable", soccerSeasonId)}
}
