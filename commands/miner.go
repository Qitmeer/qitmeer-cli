package commands

import (
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	minerCmds := []*cobra.Command{
		GenerateCmd,
		GetBlockTemplateCmd,
		SubmitBlockCmd,
	}
	RootCmd.AddCommand(minerCmds...)
	RootSubCmdGroups["miner"] = minerCmds
}

//GenerateCmd cpu mine block
var GenerateCmd = &cobra.Command{
	Use:     "generate {number | default latest}",
	Short:   "generate {number}, cpu mine {number} blocks",
	Long:    "cpu mine block",
	Aliases: []string{"Generate"},
	Example: `
generate // generate latest 

generate 1
	`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		params := []interface{}{}
		if len(args) == 0 {
			params = append(params, "latest")
		} else {
			number, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				log.Error(cmd.Use+" number err: ", err)
				return
			}
			params = append(params, number)
		}

		var rs string
		rs, err = getResString("miner_generate", params)
		if err != nil {
			log.Error(cmd.Use+" err: ", err)
		} else {
			output(rs)
		}
	},
}

//GetBlockTemplateCmd get block template to mine
var GetBlockTemplateCmd = &cobra.Command{
	Use:     "getBlockTemplate",
	Short:   "getBlockTemplate; get new block work to mine",
	Long:    "get best block template to mine work",
	Aliases: []string{"getblocktemplate", "GetBlockTemplate"},
	Example: `
getBlockTemplate 
	`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		params := []interface{}{}
		rs, err := getResString("getBlockTemplate", params)
		if err != nil {
			log.Error(cmd.Use+" err: ", err)
		} else {
			output(rs)
		}
	},
}

//SubmitBlockCmd submit block
var SubmitBlockCmd = &cobra.Command{
	Use:   "submitBlock {blockHex}",
	Short: "submitBlock {blockHex}; broadcast mine block to network",
	Example: `
//submitBlock {blockHex}

submitBlock 01000000c76bc4356cf83757de0173e07696d602350fe38b5e9262843ee267a48389976c6f8fbaeee50332b3e8f7e78542a685fefac0c070bf7574997447578b3397ed560000000000000000000000000000000000000000000000000000000000000000ffff001e000000000000000057e51e5d00000000b24b19010000000001c76bc4356cf83757de0173e07696d602350fe38b5e9262843ee267a48389976c0101000000010000000000000000000000000000000000000000000000000000000000000000ffffffffffffffff0380b2e60e000000000000000000000000000e6a0c01000000524715a54582568580461c86000000001976a914c1777151516afe2b9f59bbd1479231aa2f250d2888ac00000000000000000100f902950000000000000000ffffffff0700002f6e6f782f
	`,
	Long: `
broadcast mine block to network

{blockHex}: block hex data
`,
	Aliases: []string{"submitblock", "Submitblock"},
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		params := []interface{}{args[0]}

		var rs string
		rs, err = getResString("submitBlock", params)
		if err != nil {
			log.Error(cmd.Use+" err: ", err)
		} else {
			output(rs)
		}
	},
}
