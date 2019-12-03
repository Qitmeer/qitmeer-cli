package commands

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	nodeCmds := []*cobra.Command{

		GetNodeInfoCmd,
		GetPeerInfoCmd,
	}
	RootCmd.AddCommand(nodeCmds...)
	RootSubCmdGroups["node"] = nodeCmds
}

//GetNodeInfoCmd GetNodeInfo
var GetNodeInfoCmd = &cobra.Command{
	Use:     "getNodeInfo",
	Short:   "get node info",
	Aliases: []string{"getnodeinfo"},
	Example: `
	getNodeInfo
	`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		var info string

		info, err = getResString("getNodeInfo", nil)
		if err != nil {
			log.Error(cmd.Use+" err: ", err)
		} else {
			output(info)
		}
	},
}

//GetPeerInfoCmd GetPeerInfo
var GetPeerInfoCmd = &cobra.Command{
	Use:     "getPeerInfo",
	Short:   "get Peer Info",
	Aliases: []string{"getpeerinfo"},
	Example: `
	getPeerInfo
	`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		var info string

		info, err = getResString("getPeerInfo", nil)
		if err != nil {
			log.Error(cmd.Use+" err: ", err)
		} else {
			output(info)
		}
	},
}
