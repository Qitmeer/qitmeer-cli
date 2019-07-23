package commands

import (
	"fmt"
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(GenerateCmd)
	RootCmd.AddCommand(GetBlockTemplateCmd)
	RootCmd.AddCommand(SubmitBlockCmd)

}

//GenerateCmd cpu mine block
var GenerateCmd = &cobra.Command{
	Use:   "generate {number,default latest}",
	Short: "generate {n}, cpu mine n blocks",
	Example: `
		generate
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
				fmt.Println("number error:", err)
				return
			}
			params = append(params, number)
		}

		var rs string
		rs, err = getResString("miner_generate", params)
		if err != nil {
			log.Error(cmd.Use+" err: ", err)
		} else {
			fmt.Println(rs)
		}
	},
}

//GetBlockTemplateCmd get block template to mine
var GetBlockTemplateCmd = &cobra.Command{
	Use:   "getblocktemplate",
	Short: "getblocktemplate",
	Example: `
		getblocktemplate 
	`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {

		params := []interface{}{}
		blockCount, err := getResString("getBlockTemplate", params)
		if err != nil {
			log.Error(cmd.Use+" err: ", err)
		} else {
			fmt.Println(blockCount)
		}
	},
}

//SubmitBlockCmd cpu mine block
var SubmitBlockCmd = &cobra.Command{
	Use:   "submitBlock {blockHex}",
	Short: "submitBlock blockHex",
	Example: `
		SubmitBlock  {blockHex}	
	`,
	Aliases: []string{"submitBlock", "submitblock"},
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
