package footballdata

import "fmt"

type SoccerSeasonsRequest struct{ request }

// Season Modifies the request to specify a season.
func (r SoccerSeasonsRequest) Season(num uint32) SoccerSeasonsRequest {
	r.urlValues.Set("season", fmt.Sprintf("%d", num))
	return r
}

// Do Executes the request.
func (r SoccerSeasonsRequest) Do() (s SoccerSeasonList, err error) {
	d, _, err := r.doJson("GET")
	if err != nil {
		return
	}

	err = d.Decode(&s)
	return
}

// SoccerSeasons Prepares a request to fetch the complete list of soccer seasons.
func (c *client) SoccerSeasons() SoccerSeasonsRequest {
	return SoccerSeasonsRequest{c.req("soccerseasons")}
}
