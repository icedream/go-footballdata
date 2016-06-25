package footballdata

import "fmt"

type SoccerSeasonsRequest struct{ request }

// Season modifies the request to specify a season.
func (r SoccerSeasonsRequest) Season(num uint32) SoccerSeasonsRequest {
	r.urlValues.Set("season", fmt.Sprintf("%d", num))
	return r
}

// Do executes the request.
func (r SoccerSeasonsRequest) Do() (s SoccerSeasonList, err error) {
	d, _, err := r.doJson("GET")
	if err != nil {
		return
	}

	err = d.Decode(&s)
	return
}

func (c *client) SoccerSeasons() SoccerSeasonsRequest {
	return SoccerSeasonsRequest{c.req("soccerseasons")}
}
