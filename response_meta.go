package footballdata

import (
	"net/http"
	"strconv"
)

// Contains additional information returned by the Football-Data API in the HTTP headers.
// This includes the currently authenticated user and information about the rate limitation.
type ResponseMeta struct {
	// Indicates the recognized user or returns "anonymous" if not authenticated.
	AuthenticatedClient string

	// Defines the seconds left to reset your request counter.
	RequestCounterReset uint64

	// Indicates the requests left.
	RequestsAvailable uint64
}

func responseMetaFromHeaders(h http.Header) (r ResponseMeta) {
	if v := h.Get("X-Authenticated-Client"); v != "" {
		r.AuthenticatedClient = v
	}
	if v := h.Get("X-RequestCounter-Reset"); v != "" {
		i, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			r.RequestCounterReset = i
		}
	}
	if v := h.Get("X-Requests-Available"); v != "" {
		i, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			r.RequestsAvailable = i
		}
	}
	return
}
