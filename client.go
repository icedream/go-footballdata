package footballdata

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

/*
Provides a high-level client to talk to the API that football-data.org offers.

To create an instance please use NewClient(h).
*/
type Client struct {
	h *http.Client

	// Insert an API token here if you have one. It will be sent across with all requests.
	AuthToken string
}

// Creates a new Client instance that wraps around the given HTTP client.
func NewClient(h *http.Client) *Client {
	return &Client{h: h}
}

func (c *Client) req(path string, pathValues ...interface{}) request {
	return request{c, fmt.Sprintf(path, pathValues...), url.Values{}}
}

// Executes an HTTP request with given parameters and on success returns the response wrapped in a JSON decoder.
func (c *Client) doJson(method string, path string, values url.Values) (j *json.Decoder, meta ResponseMeta, err error) {
	// Create request
	req := &http.Request{
		Method: method,
		URL:    resolveRelativeUrl(path, values),
	}

	// Set request headers
	if len(c.AuthToken) > 0 {
		req.Header.Set("X-Auth-Token", c.AuthToken)
	}
	req.Header.Set("X-Response-Control", "minified")
	req.Header.Set("User-Agent", "go-footballdata/0.0")

	// Execute request
	resp, err := c.h.Do(req)
	if err != nil {
		return
	}

	// Save metadata from headers
	meta = responseMetaFromHeaders(resp.Header)

	// Download to buffer to allow passing back a fully prepared decoder
	defer resp.Body.Close()
	buf := new(bytes.Buffer)
	io.Copy(buf, resp.Body)

	// Wrap JSON decoder around buffered data
	j = json.NewDecoder(buf)
	return
}
