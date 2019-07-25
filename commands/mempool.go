package commands

import (
	"fmt"
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	mempoolCmds := []*cobra.Command{
		GetMempoolCmd,
	}
	RootCmd.AddCommand(mempoolCmds...)
	RootSubCmdGroups["mempool"] = mempoolCmds
}

//GetMempoolCmd get mempool
var GetMempoolCmd = &cobra.Command{
	Use:     "getMempool [type] [verbose]",
	Short:   "getMempool [type] [verbose]; type: defalut regular; verbose: bool ; get mempool info",
	Aliases: []string{"getmempool", "GetMempool"},
	Example: `
getMempool
getMempool regular false
getMempool false
	`,
	Run: func(cmd *cobra.Command, args []string) {

		var err error
		var gtype string = "regular"
		var verbose bool = false

		if len(args) == 1 {
			if args[0] == "true" || args[0] == "false" {
				verbose, _ = strconv.ParseBool(args[0])
			} else {
				gtype = args[0]
			}
		} else if len(args) > 1 {
			if verbose, err = strconv.ParseBool(args[1]); err != nil {
				fmt.Println("verbose true or false", err)
				return
			}
		}

		getBlockParam := []interface{}{}
		getBlockParam = append(getBlockParam, gtype)
		getBlockParam = append(getBlockParam, verbose)

		var info string
		info, err = getResString("getMempool", getBlockParam)
		if err != nil {
			log.Error(cmd.Use+" err: ", err)
		} else {
			output(info)
		}
	},
}
