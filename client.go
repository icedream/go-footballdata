package footballdata

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

/*
Provides a high-level client implementation to talk to the API that football-data.org offers.

A new instance of Client will by default use the default HTTP client and no
authentication token. To configure this, Client provides methods to set the
token and the HTTP client. For more information, see the respective documentation
of SetHttpClient and SetToken, or take a look at the fluent-style companion
methods WithHttpClient and WithToken.
*/
type Client struct {
	httpClient http.Client

	// Insert an API token here if you have one. It will be sent across with all requests.
	authToken string
}

// NewClient creates a new Client instance that wraps around the given HTTP client.
//
// A call to this method is not necessary in order to create a working instance
// of Client. `new(footballdata.Client)` works just as fine.
func NewClient(h *http.Client) *Client {
	return &Client{httpClient: *h}
}

// SetToken sets the authentication token.
// Calling this method is *optional*.
func (c *Client) SetToken(authToken string) {
	c.authToken = authToken
}

// WithToken sets the authentication token on a copy of the current Client
// instance.
//
// This method allows for easy fluent-style usage.
func (c Client) WithToken(authToken string) *Client {
	c.authToken = authToken
	return &c
}

// SetHttpClient sets the client that should be used to send out requests.
// Calling this method is *optional*.
func (c *Client) SetHttpClient(client *http.Client) {
	if client == nil {
		panic("client must not be nil")
	}
	c.httpClient = *client
}

// WithHttpClient sets the client that should be used to send out requests on
// a copy of the current Client instance.
//
// This method allows for easy fluent-style usage.
func (c Client) WithHttpClient(client *http.Client) *Client {
	if client == nil {
		panic("client must not be nil")
	}
	c.httpClient = *client
	return &c
}

func (c *Client) req(path string, pathValues ...interface{}) request {
	return request{
		fdClient:  c,
		path:      fmt.Sprintf(path, pathValues...),
		urlValues: make(url.Values),
		headers:   map[string]string{},
	}
}

// Executes an HTTP request with given parameters and on success returns the response wrapped in a JSON decoder.
func (c *Client) doJson(method string, path string, values url.Values, headers map[string]string) (j *json.Decoder, meta ResponseMeta, err error) {
	// Create request
	req := &http.Request{
		Method: method,
		URL:    resolveRelativeUrl(path, values),
		Header: http.Header{},
	}

	// Set request headers
	if len(c.authToken) > 0 {
		req.Header.Set("X-Auth-Token", c.authToken)
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	req.Header.Set("User-Agent", "go-footballdata/0.1")

	// Execute request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// Save metadata from headers
	meta = responseMetaFromHeaders(resp.Header)

	// Expect the request to be successful, otherwise throw an error
	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
		return
	}

	// Download to buffer to allow passing back a fully prepared decoder
	buf := new(bytes.Buffer)
	io.Copy(buf, resp.Body)

	// Wrap JSON decoder around buffered data
	j = json.NewDecoder(buf)
	return
}
