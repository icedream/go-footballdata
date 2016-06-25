package footballdata

import (
	"encoding/json"
	"net/url"
)

type request struct {
	fdClient  *client
	path      string
	urlValues url.Values
}

func (r request) doJson(method string) (*json.Decoder, ResponseMeta, error) {
	return r.fdClient.doJson(method, r.path, r.urlValues)
}
