package commands

import (
	"fmt"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(GetBlockCountCmd)
	RootCmd.AddCommand(GetBlockHashCmd)
	RootCmd.AddCommand(GetBlockCmd)
	RootCmd.AddCommand(GetBlockhashByRangeCmd)
	RootCmd.AddCommand(GetBlockByOrderCmd)
	RootCmd.AddCommand(GetBestBlockHashCmd)
	RootCmd.AddCommand(GetBlockHeaderCmd)
	RootCmd.AddCommand(IsOnMainChainCmd)
	RootCmd.AddCommand(GetMainChainHeightCmd)
	RootCmd.AddCommand(GetBlockWeightCmd)
}

//GetBlockCountCmd get block count
var GetBlockCountCmd = &cobra.Command{
	Use:   "getblockcount",
	Short: "get block count",
	Example: `
		getblockcount 
	`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {

		params := []interface{}{}
		blockCount, err := getResString("getBlockCount", params)
		if err != nil {
			log.Error(cmd.Use+" err: ", err)
		} else {
			fmt.Println(blockCount)
		}
	},
}

//GetBlockHashCmd get block hash by number
var GetBlockHashCmd = &cobra.Command{
	Use:   "getblockhash {number}",
	Short: "get block hash by number",
	Example: `
		getblockhash 100 
	`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		blockNUmber, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Error(cmd.Use + " err: block number is not int")
			return
		}

		params := []interface{}{blockNUmber}

		blockHash, err := getResString("getBlockhash", params)
		if err != nil {
			log.Error(cmd.Use+" err: ", err)
		} else {
			// blockHash= " \"xxxx\" "
			blockHash = strings.Trim(blockHash, "\"")
			fmt.Println(blockHash)
		}
	},
}

//GetBlockCmd get block by number or hash
var GetBlockCmd = &cobra.Command{
	Use:   "getblock {number|hash} {bool,show detail,defalut true}",
	Short: "get block by number or hash",
	Example: `
		getblock 100 false
		getblock 100
		getblock 000000e4c6b7f5b89827711d412957bfff5c51730df05c2eedd1352468313eca
		getblock 000000e4c6b7f5b89827711d412957bfff5c51730df05c2eedd1352468313eca true
	`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		var err error
		var blockHash string

		if len(args[0]) != 64 {
			//block number
			var blockNUmber int64
			blockNUmber, err = strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				fmt.Println("block number is not int or hash wrong")
				return
			}

			blockHash, err = getResString("getBlockhash", []interface{}{blockNUmber})
			if err != nil {
				log.Error(cmd.Use+" err: ", err)
				return
			}
			// blockHash= " \"xxxx\" "
			blockHash = strings.Trim(blockHash, "\"")
		} else {
			blockHash = args[0]
		}

		var isDetail bool = true
		if len(args) > 1 {
			isDetail, err = strconv.ParseBool(args[1])
			if err != nil {
				fmt.Println("isDetail bool true or false", err)
				return
			}
		}

		getBlockParam := []interface{}{}
		getBlockParam = append(getBlockParam, blockHash)
		getBlockParam = append(getBlockParam, isDetail)

		var blockInfo string
		blockInfo, err = getResString("getBlock", getBlockParam)
		if err != nil {
			log.Error(cmd.Use+" err: ", err)
		} else {
			output(blockInfo)
		}
	},
}

//GetBlockhashByRangeCmd get block hash by number
var GetBlockhashByRangeCmd = &cobra.Command{
	Use:     "getBlockhashByRange {start} {end}",
	Aliases: []string{"getblockhashbyrange"},
	Short:   "getBlockhashByRange start end",
	Example: `
		getBlockhashByRange 10 90
		Return the hash range of block from 'start' to 'end'(exclude self)
		if 'end' is equal to zero, 'start' is the number that from the last block to the Gen
		if 'start' is greater than or equal to 'end', it will just return the hash of 'start'
	`,
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

		start, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Error(cmd.Use + " err: block number is not int")
			return
		}

		end, err := strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			log.Error(cmd.Use + " err: block number is not int")
			return
		}

		params := []interface{}{start, end}

		blockHash, err := getResString("getBlockhashByRange", params)
		if err != nil {
			log.Error(cmd.Use+" err: ", err)
		} else {
			// blockHash= " \"xxxx\" "
			blockHash = strings.Trim(blockHash, "\"")
			output(blockHash)
		}
	},
}

//GetBlockByOrderCmd get block hash by number
var GetBlockByOrderCmd = &cobra.Command{
	Use:     "getBlockByOrder {order} {fullTx}",
	Aliases: []string{"getblockbyorder"},
	Short:   "getblockbyorder uint64 bool",
	Example: `
		getblockbyorder 10 true
	`,
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

		order, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Error(cmd.Use + " err: block number is not int")
			return
		}

		fullTx, err := strconv.ParseBool(args[1])
		if err != nil {
			log.Error(cmd.Use + " err: fullTx bool")
			return
		}

		params := []interface{}{order, fullTx}

		blockInfo, err := getResString("getBlockByOrder", params)
		if err != nil {
			log.Error(cmd.Use+" err: ", err)
		} else {
			output(blockInfo)
		}
	},
}

//GetBestBlockHashCmd get block hash by number
var GetBestBlockHashCmd = &cobra.Command{
	Use:     "getBestBlockHash",
	Short:   "getBestBlockHash",
	Aliases: []string{"getbestblockhash"},
	Example: `
		getBestBlockHash 
	`,
	Run: func(cmd *cobra.Command, args []string) {
		blockHash, err := getResString("getBestBlockHash", nil)
		if err != nil {
			log.Error(cmd.Use+" err: ", err)
		} else {
			blockHash = strings.Trim(blockHash, "\"")
			fmt.Println(blockHash)
		}
	},
}

//GetBlockHeaderCmd  get block by number or hash
var GetBlockHeaderCmd = &cobra.Command{
	Use:     "getblockheader {number|hash} {bool,show detail,defalut true}",
	Aliases: []string{"GetBlockHeader", "getBlockHeader"},
	Short:   "get block by number or hash",
	Example: `
		getBlockHeader 100 false
		getBlockHeader 100
		getBlockHeader 000000e4c6b7f5b89827711d412957bfff5c51730df05c2eedd1352468313eca
		getBlockHeader 000000e4c6b7f5b89827711d412957bfff5c51730df05c2eedd1352468313eca true
	`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		var err error
		var blockHash string

		if len(args[0]) != 64 {
			//block number
			var blockNUmber int64
			blockNUmber, err = strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				fmt.Println("block number is not int or hash wrong")
				return
			}

			blockHash, err = getResString("getBlockhash", []interface{}{blockNUmber})
			if err != nil {
				log.Error(cmd.Use+" err: ", err)
				return
			}
			// blockHash= " \"xxxx\" "
			blockHash = strings.Trim(blockHash, "\"")
		} else {
			blockHash = args[0]
		}

		var isDetail bool = true
		if len(args) > 1 {
			isDetail, err = strconv.ParseBool(args[1])
			if err != nil {
				fmt.Println("isDetail bool true or false", err)
				return
			}
		}

		getBlockParam := []interface{}{}
		getBlockParam = append(getBlockParam, blockHash)
		getBlockParam = append(getBlockParam, isDetail)

		var blockInfo string
		blockInfo, err = getResString("getBlockHeader", getBlockParam)
		if err != nil {
			log.Error(cmd.Use+" err: ", err)
		} else {
			output(blockInfo)
		}
	},
}

//IsOnMainChainCmd .
var IsOnMainChainCmd = &cobra.Command{
	Use:     "isOnMainChain {hash}",
	Short:   "query whether a given block is on the main chain",
	Aliases: []string{"isOnMainChain", "isonmainchain"},
	Example: `
		isOnMainChain 0000006c77a308846e0e0759bef5ebe0dbf4d49f345b08bdda24642efcc0cb91
	`,
	Run: func(cmd *cobra.Command, args []string) {

		params := []interface{}{args[0]}

		isOn, err := getResString("isOnMainChain", params)
		if err != nil {
			log.Error(cmd.Use+" err: ", err)
		} else {
			fmt.Println(strings.Trim(isOn, "\""))
		}
	},
}

//GetMainChainHeightCmd .
var GetMainChainHeightCmd = &cobra.Command{
	Use:     "getMainChainHeight",
	Short:   "getMainChainHeight",
	Aliases: []string{"getMainChainHeight", "getmainchainheight"},
	Example: `
		GetMainChainHeight
	`,
	Run: func(cmd *cobra.Command, args []string) {

		height, err := getResString("getMainChainHeight", nil)
		if err != nil {
			log.Error(cmd.Use+" err: ", err)
		} else {
			fmt.Println(strings.Trim(height, "\""))
		}
	},
}

//GetBlockWeightCmd .
var GetBlockWeightCmd = &cobra.Command{
	Use:     "getBlockWeight {hash}",
	Short:   "getBlockWeight",
	Aliases: []string{"getBlockWeight", "getblockweight"},
	Example: `
		getBlockWeight 0000006c77a308846e0e0759bef5ebe0dbf4d49f345b08bdda24642efcc0cb91
	`,
	Run: func(cmd *cobra.Command, args []string) {

		params := []interface{}{args[0]}

		isOn, err := getResString("getBlockWeight", params)
		if err != nil {
			log.Error(cmd.Use+" err: ", err)
		} else {
			fmt.Println(strings.Trim(isOn, "\""))
		}
	},
}
