package footballdata

type MatchRequest struct{ request }

// Do executes the request.
func (r MatchRequest) Do() (s Match, err error) {
	d, _, err := r.doJson("GET")
	if err != nil {
		return
	}

	err = d.Decode(&s)
	return
}

// Match prepares a request to fetch the matchs of a soccer season.
func (c *Client) Match(id uint64) MatchRequest {
	return MatchRequest{c.req("match/%d", id)}
}
