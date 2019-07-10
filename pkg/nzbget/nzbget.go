package nzbget

import (
	"net/http"
)

// Nzbget represents a client for an Nzbget server
type Nzbget struct {
	endpoint      string
	httpClient    *http.Client
	customHeaders map[string]string
}

// TODO: build url with host, port, username, password,
// TODO: category for download
// TODO: priority for download

func NewClient(address string) Nzbget {

	return Nzbget{
		endpoint: address + "/xmlrpc",
	}

}
