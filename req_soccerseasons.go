package footballdata

import "fmt"

// DEPRECATED.
//
type SoccerSeasonsRequest struct{ request }

// DEPRECATED.
//
// Season modifies the request to specify a season.
func (r SoccerSeasonsRequest) Season(num uint32) SoccerSeasonsRequest {
	r.urlValues.Set("season", fmt.Sprintf("%d", num))
	return r
}

// DEPRECATED.
//
// Do executes the request.
func (r SoccerSeasonsRequest) Do() (s SoccerSeasonList, err error) {
	d, _, err := r.doJson("GET")
	if err != nil {
		return
	}

	err = d.Decode(&s)
	return
}

// DEPRECATED.
//
// SoccerSeasons prepares a request to fetch the list of soccer seasons for the current season or for any specified season (via the "Season" submethod).
func (c *Client) SoccerSeasons() SoccerSeasonsRequest {
	return SoccerSeasonsRequest{c.req("soccerseasons")}
}
