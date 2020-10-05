package artifactory

// Artifact represents details of an artifact with a repository
type Artifact struct {
	Repo       string `json:"repo"`
	Path       string `json:"path"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Size       int    `json:"size"`
	Created    string `json:"created"` // should be timestamp
	CreatedBy  string `json:"created_by"`
	Modified   string `json:"modified"` // should be timestamp
	ModifiedBy string `json:"modified_by"`
	Updated    string `json:"updated"` // should be timestamp
}

// AQLSearchResponse represents response obtained from AQL repo search endpoint
type AQLSearchResponse struct {
	Results []Artifact `json:"results"`
	Range   struct {
		StartPos int `json:"start_pos"`
		EndPos   int `json:"end_pos"`
		Total    int `json:"total"`
	} `json:"range"`
}

// FileStatisticsResponse represents response obtained from file statistic endpoint
type FileStatisticsResponse struct {
	URI                  string `json:"uri"`
	DownloadCount        int    `json:"downloadCount"`
	LastDownloaded       int    `json:"lastDownloaded"`
	LastDownloadedBy     string `json:"lastDownloadedBy"`
	RemoteDownloadCount  int    `json:"remoteDownloadCount"`
	RemoteLastDownloaded int    `json:"remoteLastDownloaded"`
}
