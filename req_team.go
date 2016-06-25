package footballdata

type TeamRequest struct {
	request
	id uint64
}

// Do executes the request.
func (r TeamRequest) Do() (s Team, err error) {
	d, _, err := r.doJson("GET")
	if err != nil {
		return
	}

	err = d.Decode(&s)
	if err != nil {
		// Fill up data not returned by the server
		s.Id = r.id
	}
	return
}

func (c *client) Team(id uint64) TeamRequest {
	return TeamRequest{c.req("teams/%d", id), id}
}
