package footballdata

import (
	"fmt"
	"strings"
)

type AreasRequest struct{ request }

// Areas modifies the request to specify a filter for specific area information that should be returned.
func (r AreasRequest) Areas(areaIDs ...uint32) AreasRequest {
	areasStr := []string{}
	for _, areaID := range areaIDs {
		areasStr = append(areasStr, fmt.Sprintf("%d", areaID))
	}
	r.urlValues.Set("areas", strings.Join(areasStr, ","))
	return r
}

// Do executes the request.
func (r AreasRequest) Do() (s AreaList, err error) {
	d, _, err := r.doJson("GET")
	if err != nil {
		return
	}

	err = d.Decode(&s)
	return
}

// Areas prepares a request to fetch the list of areas.
func (c *Client) Areas() AreasRequest {
	return AreasRequest{c.req("areas")}
}
