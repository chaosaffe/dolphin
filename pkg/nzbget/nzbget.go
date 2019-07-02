package nzbget

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/pkg/errors"
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
		endpoint: address + "/jsonrpc",
	}

}

func (c *Nzbget) doCall(request interface{}, response interface{}) error {
	data, err := json.Marshal(request)
	if err != nil {
		return errors.Wrap(err, "unable to marshal request")
	}
	resp, err := http.Post(c.endpoint,
		"application/json", strings.NewReader(string(data)))
	if err != nil {
		return errors.Wrap(err, "posting request failed")
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	//decoder.DisallowUnknownFields()
	decoder.UseNumber()
	err = decoder.Decode(&response)
	if err != nil {
		return errors.Wrap(err, "unable to unmarshal response")
	}
	return nil
}
