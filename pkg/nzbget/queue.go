package nzbget

import (
	"github.com/pkg/errors"
)

type listGroupsRequest struct {
	Method     string               `json:"method"`
	ID         interface{}          `json:"id,omitempty"`
	Parameters listGroupsParameters `json:"param"`
}

type listGroupsResponse struct {
	APIVersion string `json:"version"`
	ID         string `json:"id"`
	Result     []NZB  `json:"result"`
}

type listGroupsParameters struct {
	ItemCount int `json:"NumberOfLogEntries"`
}

// TODO: combine hi/low's

type NZB struct {
	RemainingSizeLo    int64  `json:"RemainingSizeLo"`
	RemainingSizeHi    int    `json:"RemainingSizeHi"`
	RemainingSizeMB    int    `json:"RemainingSizeMB"`
	PausedSizeLo       int    `json:"PausedSizeLo"`
	PausedSizeHi       int    `json:"PausedSizeHi"`
	PausedSizeMB       int    `json:"PausedSizeMB"`
	RemainingFileCount int    `json:"RemainingFileCount"`
	RemainingParCount  int    `json:"RemainingParCount"`
	MinPriority        int    `json:"MinPriority"`
	MaxPriority        int    `json:"MaxPriority"`
	ActiveDownloads    int    `json:"ActiveDownloads"`
	Status             string `json:"Status"`
	ID                 int    `json:"NZBID"`
	Name               string `json:"NZBName"`
	Kind               string `json:"Kind"`
	URL                string `json:"URL"`
	Filename           string `json:"NZBFilename"`
	DestDir            string `json:"DestDir"`
	FinalDir           string `json:"FinalDir"`
	Category           string `json:"Category"`
	ParStatus          string `json:"ParStatus"`
	ExParStatus        string `json:"ExParStatus"`
	UnpackStatus       string `json:"UnpackStatus"`
	MoveStatus         string `json:"MoveStatus"`
	ScriptStatus       string `json:"ScriptStatus"`
	DeleteStatus       string `json:"DeleteStatus"`
	MarkStatus         string `json:"MarkStatus"`
	URLStatus          string `json:"UrlStatus"`
	FileSizeLo         int64  `json:"FileSizeLo"`
	FileSizeHi         int    `json:"FileSizeHi"`
	FileSizeMB         int    `json:"FileSizeMB"`
	FileCount          int    `json:"FileCount"`
	MinPostTime        int    `json:"MinPostTime"`
	MaxPostTime        int    `json:"MaxPostTime"`
	TotalArticles      int    `json:"TotalArticles"`
	SuccessArticles    int    `json:"SuccessArticles"`
	FailedArticles     int    `json:"FailedArticles"`
	Health             int    `json:"Health"`
	CriticalHealth     int    `json:"CriticalHealth"`
	DuplicateKey       string `json:"DupeKey"`
	DuplicateScore     int    `json:"DupeScore"`
	DuplicateMode      string `json:"DupeMode"`
	DownloadedSizeLo   int    `json:"DownloadedSizeLo"`
	DownloadedSizeHi   int    `json:"DownloadedSizeHi"`
	DownloadedSizeMB   int    `json:"DownloadedSizeMB"`
	DownloadTimeSec    int    `json:"DownloadTimeSec"`
	PostTotalTimeSec   int    `json:"PostTotalTimeSec"`
	ParTimeSec         int    `json:"ParTimeSec"`
	RepairTimeSec      int    `json:"RepairTimeSec"`
	UnpackTimeSec      int    `json:"UnpackTimeSec"`
	MessageCount       int    `json:"MessageCount"`
	ExtraParBlocks     int    `json:"ExtraParBlocks"`
	Parameters         []struct {
		Name  string `json:"Name"`
		Value string `json:"Value"`
	} `json:"Parameters"`
	ScriptStatuses []interface{} `json:"ScriptStatuses"`
	ServerStats    []struct {
		ServerID        int `json:"ServerID"`
		SuccessArticles int `json:"SuccessArticles"`
		FailedArticles  int `json:"FailedArticles"`
	} `json:"ServerStats"`
	PostInfoText      string `json:"PostInfoText"`
	PostStageProgress int    `json:"PostStageProgress"`
	PostStageTimeSec  int    `json:"PostStageTimeSec"`
}

func (c *Nzbget) ListNZBs() ([]NZB, error) {

	request := listGroupsRequest{
		Method: "listgroups",
		Parameters: listGroupsParameters{
			// ItemCount is deprecated, must be zero (https://nzbget.net/api/listgroups)
			ItemCount: 0,
		},
	}

	response := listGroupsResponse{}

	err := c.doCall(request, &response)
	if err != nil {
		return nil, errors.Wrap(err, "unable to list groups")
	}

	return response.Result, nil

}
