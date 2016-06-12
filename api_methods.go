package footballdata

import (
	"encoding/json"
	"net/url"
)

type request struct {
	c *Client
	p string
	v url.Values
}

func (r request) doJson(method string) (*json.Decoder, ResponseMeta, error) {
	return r.c.doJson(method, r.p, r.v)
}
