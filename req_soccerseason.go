package footballdata

type SoccerSeasonRequest struct{ request }

// Do executes the request.
func (r SoccerSeasonRequest) Do() (s SoccerSeason, err error) {
	d, _, err := r.doJson("GET")
	if err != nil {
		return
	}

	err = d.Decode(&s)
	return
}

// SoccerSeason prepares a request to fetch the complete list of soccer seasons.
func (c *client) SoccerSeason(id uint64) SoccerSeasonRequest {
	return SoccerSeasonRequest{c.req("soccerseasons/%d", id)}
}
