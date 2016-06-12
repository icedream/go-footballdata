package footballdata

import "net/url"

const (
	baseUrl = "http://api.football-data.org/v1/"
)

func resolveRelativeUrl(path string, values url.Values) *url.URL {
	if values == nil {
		values = url.Values{}
	}

	u, err := url.Parse(baseUrl)
	if err != nil {
		panic(err)
	}

	ru := &url.URL{Path: path, RawQuery: values.Encode()}

	return u.ResolveReference(ru)
}
