package main

import (
	"fmt"
	"os"

	samples "github.com/vhugo/opsgenie-go-sdk/_samples"
	"github.com/vhugo/opsgenie-go-sdk/_samples/constants"
	alerts "github.com/vhugo/opsgenie-go-sdk/alerts"
	ogcli "github.com/vhugo/opsgenie-go-sdk/client"
)

func main() {
	cli := new(ogcli.OpsGenieClient)
	cli.SetAPIKey(constants.APIKey)

	alertCli, cliErr := cli.Alert()

	if cliErr != nil {
		panic(cliErr)
	}

	// create the alert
	req := alerts.CreateAlertRequest{Message: samples.RandStringWithPrefix("Test", 8)}
	response, alertErr := alertCli.Create(req)

	if alertErr != nil {
		panic(alertErr)
	}

	fmt.Printf("message: %s\n", response.Message)
	fmt.Printf("alert id: %s\n", response.AlertID)
	fmt.Printf("status: %s\n", response.Status)
	fmt.Printf("code: %d\n", response.Code)

	file, err := os.OpenFile(constants.PathToFile, os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	attachFileReq := alerts.AttachFileAlertRequest{ID: response.AlertID, Attachment: file}
	attachFileResp, attachFileErr := alertCli.AttachFile(attachFileReq)

	if attachFileErr != nil {
		panic(attachFileErr)
	}

	fmt.Printf("Status: %s\n", attachFileResp.Status)
	fmt.Printf("Code: %d\n", attachFileResp.Code)
}
