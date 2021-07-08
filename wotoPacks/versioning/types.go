package versioning

type Version struct {
	Num1 uint8
	Num2 uint8
	Num3 uint8
	Num4 uint8
}

type VersionResp struct {
	// IsAcceptable will be true if and only if the
	IsAcceptable bool `json:"is_acceptable"`

	DataDownloadLink *string `json:"data_download_link"`

	AppDownloadLink *string `json:"app_download_link"`
}
