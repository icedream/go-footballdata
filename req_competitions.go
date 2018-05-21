package footballdata

import (
	"fmt"
	"strings"
)

type CompetitionsRequest struct{ request }

// Areas modifies the request to specify a filter for specific area information that should be returned.
func (r AreasRequest) Areas(areaIDs ...uint32) AreasRequest {
	areasStr := []string{}
	for _, areaID := range areaIDs {
		areasStr = append(areasStr, fmt.Sprintf("%d", areaID))
	}
	r.urlValues.Set("areas", strings.Join(areasStr, ","))
	return r
}

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
