package commands

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/HalalChain/qitmeer-cli/rpc/client"
)

const (
	// DefaultServiceNameSpace default api
	DefaultServiceNameSpace = "qitmeer"
	// MinerNameSpace  miner api
	MinerNameSpace = "miner"
)

// RootCmd command obj
var RootCmd = &cobra.Command{
	Use:  "qitmeer-cli",
	Long: `qitmeer cli is a RPC tool for the qitmeer and qitmeer-wallet`,
	//PersistentPreRunE: rootCmdPreRun,
}

var RPCCfg = &client.Config{}
var RPCVersion = "1.0"
var Format = false

func getResString(method string, args []interface{}) (rs string, err error) {
	reqData, err := client.MakeRequestData(RPCVersion, 1, method, args)
	if err != nil {
		return
	}

	resResult, err := client.SendPostRequest(reqData, RPCCfg)
	if err != nil {
		return
	}

	rs = string(resResult)
	return
}

//
func output(dataStr string) {
	if Format {
		var str bytes.Buffer
		_ = json.Indent(&str, []byte(dataStr), "", "    ")
		fmt.Println(str.String())
	} else {
		fmt.Println(dataStr)
	}
}
