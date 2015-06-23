package main

import (
	alerts "github.com/opsgenie/opsgenie-go-sdk/alerts"
	ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
	"fmt"
	samples "github.com/opsgenie/opsgenie-go-sdk/samples"
)

var API_KEY string = "YOUR API KEY HERE"
var ACTIONS []string = []string{"ping", "pong"}
var ACTION_TO_EXEC string = "pong"

func main() {
	
	cli := new (ogcli.OpsGenieClient)
	cli.SetApiKey(API_KEY)
	
	alertCli, cliErr := cli.Alert()
	
	if cliErr != nil {
		panic(cliErr)
	}

	// create the alert
	req := alerts.CreateAlertRequest{Message: samples.RandStringWithPrefix("Test - ",10), Actions: ACTIONS,}
	response, alertErr := alertCli.Create(req)
	
	if alertErr != nil {
		panic(alertErr)
	}

	fmt.Println("message:", response.Message)
	fmt.Println("alert id:",response.AlertId)
	fmt.Println("status:", 	response.Status)
	fmt.Println("code:", 	response.Code)

	// execute sample 'pong' action for the alert
	execActionReq := alerts.ExecuteActionAlertRequest{ Id: response.AlertId, Action: ACTION_TO_EXEC, Note: "Action <b>pong</b> executed by the Go API" }
	execActionResponse, alertErr := alertCli.ExecuteAction(execActionReq)
	
	if alertErr != nil {
		panic(alertErr)
	}

	fmt.Println("status:", 	execActionResponse.Result)
	fmt.Println("code:", 	execActionResponse.Code)
}
