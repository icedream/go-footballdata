package footballdata

type CompetitionRequest struct{ request }

// Do executes the request.
func (r CompetitionRequest) Do() (s Competition, err error) {
	d, _, err := r.doJson("GET")
	if err != nil {
		return
	}

	err = d.Decode(&s)
	return
}

// Competition prepares a request to fetch the complete list of soccer seasons.
func (c *Client) Competition(id uint64) CompetitionRequest {
	return CompetitionRequest{c.req("competitions/%d", id)}
}
