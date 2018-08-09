package command

import (
	"os"
	"fmt"
	gcli "github.com/codegangsta/cli"
	cl "github.com/opsgenie/opsgenie-go-sdk/customerlog"
	"bytes"
)

func GetLogLink(c *gcli.Context) {
	cli, err := NewCustomerLogClient(c)
	if err != nil {
		os.Exit(1)
	}

	req := cl.CustomerLogGetLinkRequest{}
	if val, success := getVal("logFile", c); success {
		req.LogFile = val
	}

	printVerboseMessage("Get Link request prepared from flags, sending request to OpsGenie..")

	response, err := cli.GetLink(req)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}
	fmt.Printf("URL=%d\n", response.URL)
}

func ListDownloadableLogs(c *gcli.Context) {
	cli, err := NewCustomerLogClient(c)
	if err != nil {
		os.Exit(1)
	}

	req := cl.CustomerLogListDownloadablesRequest{}
	if val, success := getVal("after", c); success {
		req.After = val
	}

	printVerboseMessage("List Downloadable Logs request prepared from flags, sending request to OpsGenie..")

	response, err := cli.DownloadableList(req)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}
	var buffer bytes.Buffer
	downloadables := response.Downloadables
	for _, downloadable := range downloadables {
		buffer.WriteString(downloadable+", ")
	}
	fmt.Printf(buffer.String())
}