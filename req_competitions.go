package footballdata

import "fmt"

type CompetitionsRequest struct{ request }

// Season modifies the request to specify a season.
func (r CompetitionsRequest) Season(num uint32) CompetitionsRequest {
	r.urlValues.Set("season", fmt.Sprintf("%d", num))
	return r
}

// Do executes the request.
func (r CompetitionsRequest) Do() (s CompetitionList, err error) {
	d, _, err := r.doJson("GET")
	if err != nil {
		return
	}

	err = d.Decode(&s)
	return
}

// Competitions prepares a request to fetch the list of competitions for the current season or for any specified season (via the "Season" submethod).
func (c *Client) Competitions() CompetitionsRequest {
	return CompetitionsRequest{c.req("competitions")}
}
