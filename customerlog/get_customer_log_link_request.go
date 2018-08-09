package customerlog

import (
	"net/url"
)

// CustomerLogGetLinkRequest is a struct of request to list new customerLog.
type CustomerLogGetLinkRequest struct {
	LogFile string
	ApiKey  string
}

// GetApiKey returns api key.
func (r *CustomerLogGetLinkRequest) GetApiKey() string {
	return r.ApiKey
}

func (r *CustomerLogGetLinkRequest) GetLogFile() string {
	return r.LogFile
}

// GenerateUrl generates url to API endpoint.
func (r *CustomerLogGetLinkRequest) GenerateUrl() (string, url.Values, error) {
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
