package footballdata

import "fmt"

type SoccerSeasonsRequest struct{ request }

// Modifies the request to specify a season.
func (r SoccerSeasonsRequest) Season(num uint32) SoccerSeasonsRequest {
	r.v.Set("season", fmt.Sprintf("%i", num))
	return r
}

// Executes the request.
func (r SoccerSeasonsRequest) Do() (s SoccerSeasonList, err error) {
	d, _, err := r.doJson("GET")
	if err != nil {
		return
	}

	err = d.Decode(&s)
	return
}

// Prepares a request to fetch the complete list of soccer seasons.
func (c *Client) SoccerSeasons() SoccerSeasonsRequest {
	return SoccerSeasonsRequest{c.req("soccerseasons")}
}
