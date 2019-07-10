package nzbget

import (
	"github.com/kolo/xmlrpc"
	"github.com/pkg/errors"
)

// TODO: combine hi/low's

type listNZBsParameters struct {
	NumberOfLogEntries int `xmlrpc:"NumberOfLogEntries"`
}

type NZB struct {
	RemainingSizeLo    int64  `xmlrpc:"RemainingSizeLo"`
	RemainingSizeHi    int    `xmlrpc:"RemainingSizeHi"`
	RemainingSizeMB    int    `xmlrpc:"RemainingSizeMB"`
	PausedSizeLo       int    `xmlrpc:"PausedSizeLo"`
	PausedSizeHi       int    `xmlrpc:"PausedSizeHi"`
	PausedSizeMB       int    `xmlrpc:"PausedSizeMB"`
	RemainingFileCount int    `xmlrpc:"RemainingFileCount"`
	RemainingParCount  int    `xmlrpc:"RemainingParCount"`
	MinPriority        int    `xmlrpc:"MinPriority"`
	MaxPriority        int    `xmlrpc:"MaxPriority"`
	ActiveDownloads    int    `xmlrpc:"ActiveDownloads"`
	Status             string `xmlrpc:"Status"`
	ID                 int    `xmlrpc:"NZBID"`
	Name               string `xmlrpc:"NZBName"`
	Kind               string `xmlrpc:"Kind"`
	URL                string `xmlrpc:"URL"`
	Filename           string `xmlrpc:"NZBFilename"`
	DestDir            string `xmlrpc:"DestDir"`
	FinalDir           string `xmlrpc:"FinalDir"`
	Category           string `xmlrpc:"Category"`
	ParStatus          string `xmlrpc:"ParStatus"`
	ExParStatus        string `xmlrpc:"ExParStatus"`
	UnpackStatus       string `xmlrpc:"UnpackStatus"`
	MoveStatus         string `xmlrpc:"MoveStatus"`
	ScriptStatus       string `xmlrpc:"ScriptStatus"`
	DeleteStatus       string `xmlrpc:"DeleteStatus"`
	MarkStatus         string `xmlrpc:"MarkStatus"`
	URLStatus          string `xmlrpc:"UrlStatus"`
	FileSizeLo         int64  `xmlrpc:"FileSizeLo"`
	FileSizeHi         int    `xmlrpc:"FileSizeHi"`
	FileSizeMB         int    `xmlrpc:"FileSizeMB"`
	FileCount          int    `xmlrpc:"FileCount"`
	MinPostTime        int    `xmlrpc:"MinPostTime"`
	MaxPostTime        int    `xmlrpc:"MaxPostTime"`
	TotalArticles      int    `xmlrpc:"TotalArticles"`
	SuccessArticles    int    `xmlrpc:"SuccessArticles"`
	FailedArticles     int    `xmlrpc:"FailedArticles"`
	Health             int    `xmlrpc:"Health"`
	CriticalHealth     int    `xmlrpc:"CriticalHealth"`
	DuplicateKey       string `xmlrpc:"DupeKey"`
	DuplicateScore     int    `xmlrpc:"DupeScore"`
	DuplicateMode      string `xmlrpc:"DupeMode"`
	DownloadedSizeLo   int    `xmlrpc:"DownloadedSizeLo"`
	DownloadedSizeHi   int    `xmlrpc:"DownloadedSizeHi"`
	DownloadedSizeMB   int    `xmlrpc:"DownloadedSizeMB"`
	DownloadTimeSec    int    `xmlrpc:"DownloadTimeSec"`
	PostTotalTimeSec   int    `xmlrpc:"PostTotalTimeSec"`
	ParTimeSec         int    `xmlrpc:"ParTimeSec"`
	RepairTimeSec      int    `xmlrpc:"RepairTimeSec"`
	UnpackTimeSec      int    `xmlrpc:"UnpackTimeSec"`
	MessageCount       int    `xmlrpc:"MessageCount"`
	ExtraParBlocks     int    `xmlrpc:"ExtraParBlocks"`
	Parameters         []struct {
		Name  string `xmlrpc:"Name"`
		Value string `xmlrpc:"Value"`
	} `xmlrpc:"Parameters"`
	ScriptStatuses []interface{} `xmlrpc:"ScriptStatuses"`
	ServerStats    []struct {
		ServerID        int `xmlrpc:"ServerID"`
		SuccessArticles int `xmlrpc:"SuccessArticles"`
		FailedArticles  int `xmlrpc:"FailedArticles"`
	} `xmlrpc:"ServerStats"`
	PostInfoText      string `xmlrpc:"PostInfoText"`
	PostStageProgress int    `xmlrpc:"PostStageProgress"`
	PostStageTimeSec  int    `xmlrpc:"PostStageTimeSec"`
}

func (c *Nzbget) ListNZBs() ([]NZB, error) {

	args := listNZBsParameters{
		// NumberOfLogEntries is deprecated, must be zero (https://nzbget.net/api/listgroups)
		NumberOfLogEntries: 0,
	}

	response := []NZB{}

	client, _ := xmlrpc.NewClient(c.endpoint, nil)
	err := client.Call("listgroups", args, &response)
	if err != nil {
		return nil, errors.Wrap(err, "unable to list groups")
	}

	return response, nil

}
