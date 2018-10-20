package main

import (
	"github.com/vhugo/opsgenie-go-sdk/_samples/constants"
	"github.com/vhugo/opsgenie-go-sdk/alertsv2"
	"fmt"
	ogcli "github.com/vhugo/opsgenie-go-sdk/client"
	"encoding/json"
)

func main() {
	cli := new(ogcli.OpsGenieClient)
	cli.SetAPIKey(constants.APIKey)

	alertCli, _ := cli.AlertV2()

	response, err := alertCli.Get(alertsv2.GetAlertRequest{
		Identifier: &alertsv2.Identifier{
			TinyID: "2",
		},
	})

	if err != nil {
		fmt.Println(err.Error())
	} else {
		alert := response.Alert
		alertJson, _ := json.Marshal(alert)

		fmt.Println(string(alertJson))
	}
}
