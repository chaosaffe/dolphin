package nzbget

import (
	"github.com/kolo/xmlrpc"
	"github.com/pkg/errors"
)

type versionResponse struct {
	Version string `xmlrpc:"value"`
}

func (c *Nzbget) Version() (string, error) {

	response := ""

	client, _ := xmlrpc.NewClient(c.endpoint, nil)
	err := client.Call("version", nil, &response)
	if err != nil {
		return "", errors.Wrap(err, "unable to get version")
	}

	return response, nil

}

//<?xml version="1.0"?>
//<methodResponse>
//<params><param><value><string>21.0</string></value></param></params>
//</methodResponse>
