package nzbget

import "github.com/pkg/errors"

type versionRequest struct {
	Method string `json:"method"`
}

type versionResponse struct {
	APIVersion string `json:"version"`
	ID         string `json:"id"`
	Result     string `json:"result"`
}

func (c *Nzbget) Version() (string, error) {

	request := versionRequest{
		Method: "version",
	}

	response := versionResponse{}

	err := c.doCall(request, &response)
	if err != nil {
		return "", errors.Wrap(err, "unable to get version")
	}

	return response.Result, nil

}
