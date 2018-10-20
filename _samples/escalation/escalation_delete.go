package main

import (
	"fmt"
	ogcli "github.com/vhugo/opsgenie-go-sdk/client"
	esc "github.com/vhugo/opsgenie-go-sdk/escalation"
	"github.com/vhugo/opsgenie-go-sdk/_samples/constants"
)

func main() {
	cli := new(ogcli.OpsGenieClient)
	cli.SetAPIKey(constants.APIKey)

	escCli, cliErr := cli.Escalation()

	if cliErr != nil {
		panic(cliErr)
	}

	req := esc.DeleteEscalationRequest{Name:""}
	response, escErr := escCli.Delete(req)

	if escErr != nil {
		panic(escErr)
	}

	fmt.Printf("status: %s\n", response.Status)
	fmt.Printf("code: %d\n", response.Code)
}
