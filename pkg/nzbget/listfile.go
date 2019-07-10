package nzbget

import (
	"github.com/kolo/xmlrpc"
	"github.com/pkg/errors"
)

type listFilesParameters struct {
	IDFrom int `xmlrpc:"IDFrom"`
	IDTo   int `xmlrpc:"IDTo"`
	NZBID  int `xmlrpc:"NZBID"`
}

type File struct {
	ID                int    `xmlrpc:"ID"`
	FileSizeLo        int    `xmlrpc:"FileSizeLo"`
	FileSizeHi        int    `xmlrpc:"FileSizeHi"`
	RemainingSizeLo   int    `xmlrpc:"RemainingSizeLo"`
	RemainingSizeHi   int    `xmlrpc:"RemainingSizeHi"`
	PostTime          int    `xmlrpc:"PostTime"`
	FilenameConfirmed bool   `xmlrpc:"FilenameConfirmed"`
	Paused            bool   `xmlrpc:"Paused"`
	NZBID             int    `xmlrpc:"NZBID"`
	NZBName           string `xmlrpc:"NZBName"`
	NZBNicename       string `xmlrpc:"NZBNicename"`
	NZBFilename       string `xmlrpc:"NZBFilename"`
	Subject           string `xmlrpc:"Subject"`
	Filename          string `xmlrpc:"Filename"`
	DestDir           string `xmlrpc:"DestDir"`
	Category          string `xmlrpc:"Category"`
	Priority          int    `xmlrpc:"Priority"`
	ActiveDownloads   int    `xmlrpc:"ActiveDownloads"`
	Progress          int    `xmlrpc:"Progress"`
}

func (c *Nzbget) ListFiles(nzbID int) ([]File, error) {

	args := listFilesParameters{
		IDFrom: 0,
		IDTo:   0,
		NZBID:  nzbID,
	}

	response := []File{}

	client, _ := xmlrpc.NewClient(c.endpoint, nil)
	err := client.Call("listfiles", args, &response)
	if err != nil {
		return nil, errors.Wrap(err, "unable to list groups")
	}

	return response, nil

}
