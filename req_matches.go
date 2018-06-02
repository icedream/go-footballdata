package footballdata

type MatchesRequest struct{ request }

// Do executes the request.
func (r MatchesRequest) Do() (s MatchesResponse, err error) {
	d, _, err := r.doJson("GET")
	if err != nil {
		return
	}

	err = d.Decode(&s)
	return
}

// Matches prepares a request to fetch the matches of a soccer season.
func (c *Client) Matches() MatchesRequest {
	return MatchesRequest{c.req("matches")}
}
