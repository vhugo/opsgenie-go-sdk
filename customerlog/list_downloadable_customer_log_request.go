package customerlog

import (
	"net/url"
)

// CustomerLogListDownloadablesRequest is a struct of request to download customerLog.
type CustomerLogListDownloadablesRequest struct {
	After  string
	ApiKey string
}

// GetApiKey returns api key.
func (r *CustomerLogListDownloadablesRequest) GetApiKey() string {
	return r.ApiKey
}

// GenerateUrl generates url to API endpoint.
func (r *CustomerLogListDownloadablesRequest) GenerateUrl() (string, url.Values, error) {
	// todo kagan fix here
	//baseUrl, params, err := r.ScheduleIdentifier.GenerateUrl()
	//
	//if err != nil {
	//	return "", nil, err
	//}
	//
	//baseUrl += "/overrides"
	//return baseUrl, params, nil
	return "", nil, nil
}
