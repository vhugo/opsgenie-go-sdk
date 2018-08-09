package customerlog

type CustomerLogListDownloadablesResponse struct {
	ResponseMeta
	Downloadables []string `json:"data"` // todo ask mustafak
}
