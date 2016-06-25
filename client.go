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
Provides a high-level client to talk to the API that football-data.org offers.

To create an instance please use NewClient(h).
*/
type Client interface {
	Fixture(id uint64) FixtureRequest
	Fixtures() FixturesRequest
	SoccerSeason(id uint64) SoccerSeasonRequest
	FixturesOfSoccerSeason(soccerSeasonId uint64) SoccerSeasonFixturesRequest
	LeagueTableOfSoccerSeason(soccerSeasonId uint64) SoccerSeasonLeagueTableRequest
	TeamsOfSoccerSeason(soccerSeasonId uint64) SoccerSeasonTeamsRequest
	SoccerSeasons() SoccerSeasonsRequest
	Team(id uint64) TeamRequest
	FixturesOfTeam(id uint64) TeamFixturesRequest

	SetToken(authToken string)
}

/*
Provides a high-level client implementation to talk to the API that football-data.org offers.

To create an instance please use NewClient(h).
*/
type client struct {
	httpClient *http.Client

	// Insert an API token here if you have one. It will be sent across with all requests.
	AuthToken string
}

// NewClient Creates a new Client instance that wraps around the given HTTP client.
//
// Call SetToken to add your token.
func NewClient(h *http.Client) Client {
	return &client{httpClient: h}
}

// SetToken Set the authentication token
// Calling this method is *optional*
func (c *client) SetToken(authToken string) {
	c.AuthToken = authToken
}

func (c *client) req(path string, pathValues ...interface{}) request {
	return request{c, fmt.Sprintf(path, pathValues...), url.Values{}}
}

// Executes an HTTP request with given parameters and on success returns the response wrapped in a JSON decoder.
func (c *client) doJson(method string, path string, values url.Values) (j *json.Decoder, meta ResponseMeta, err error) {
	// Create request
	req := &http.Request{
		Method: method,
		URL:    resolveRelativeUrl(path, values),
		Header: http.Header{},
	}

	// Set request headers
	if len(c.AuthToken) > 0 {
		req.Header.Set("X-Auth-Token", c.AuthToken)
	}
	req.Header.Set("X-Response-Control", "minified")
	req.Header.Set("User-Agent", "go-footballdata/0.0")

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
