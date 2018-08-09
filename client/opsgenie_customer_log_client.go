package client

import (
	"github.com/opsgenie/opsgenie-go-sdk/customerlog"
)

// OpsGenieCustomerLogClient is the data type to make Customer Log rule API requests.
type OpsGenieCustomerLogClient struct {
	RestClient
}

// SetOpsGenieClient sets the embedded OpsGenieClient type of the OpsGenieScheduleRotationV2Client.
func (cli *OpsGenieCustomerLogClient) SetOpsGenieClient(ogCli OpsGenieClient) {
	cli.OpsGenieClient = ogCli
}

func (cli *OpsGenieCustomerLogClient) GetLink(req customerlog.CustomerLogGetLinkRequest) (*customerlog.CustomerLogGetLinkResponse, error) {
	var response customerlog.CustomerLogGetLinkResponse
	err := cli.sendGetRequest(&req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (cli *OpsGenieCustomerLogClient) DownloadableList(req customerlog.CustomerLogListDownloadablesRequest) (*customerlog.CustomerLogListDownloadablesResponse, error) {
	var response customerlog.CustomerLogListDownloadablesResponse
	err := cli.sendGetRequest(&req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
