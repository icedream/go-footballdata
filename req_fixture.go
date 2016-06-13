package footballdata

import "fmt"

type FixtureRequest struct{ request }

// Modifies the request to specify the number of former games to be analyzed (normally 10).
func (r FixtureRequest) Head2Head(num uint16) FixtureRequest {
	r.v.Set("head2head", fmt.Sprintf("%d", num))
	return r
}

// Executes the request.
func (r FixtureRequest) Do() (s Fixture, err error) {
	d, _, err := r.doJson("GET")
	if err != nil {
		return
	}

	err = d.Decode(&s)
	return
}

// Prepares a request to fetch the fixtures of a soccer season.
func (c *Client) Fixture(id uint64) FixtureRequest {
	return FixtureRequest{c.req("fixture/%d", id)}
}
