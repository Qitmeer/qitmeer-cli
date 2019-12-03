package commands

import (
	"fmt"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {

	blockCmds := []*cobra.Command{
		GetBlockCountCmd,
		GetBlockTotalCmd,
		GetOrphansTotalCmd,
		GetBlockHashCmd,
		GetBlockCmd,
		GetBlockhashByRangeCmd,
		GetBlockByOrderCmd,
		GetBlockByIDCmd,
		GetBestBlockHashCmd,
		GetBlockHeaderCmd,
		IsOnMainChainCmd,
		GetMainChainHeightCmd,
		GetBlockWeightCmd,
		IsBlueCmd,
	}

	RootCmd.AddCommand(blockCmds...)
	RootSubCmdGroups["block"] = blockCmds
}

//GetBlockTotalCmd The total number blocks that this dag currently owned
var GetBlockTotalCmd = &cobra.Command{
	Use:     "getBlockTotal",
	Short:   "getBlockTotal; The total number blocks that this dag currently owned",
	Aliases: []string{"GetBlockTotal", "getblocktotal"},
	Example: `
	getBlockTotal 
	`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		params := []interface{}{}
		blockCount, err := getResString("getBlockTotal", params)
		if err != nil {
			log.Error(cmd.Use+" err: ", err)
		} else {
			output(blockCount)
		}
	},
}

//GetBlockCountCmd get block count
var GetBlockCountCmd = &cobra.Command{
	Use:     "getBlockCount",
	Short:   "getBlockCount; The order of main chain tip",
	Aliases: []string{"getblockcount", "GetBlockCount"},
	Example: `
	getBlockCount 
	`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		params := []interface{}{}
		blockCount, err := getResString("getBlockCount", params)
		if err != nil {
			log.Error(cmd.Use+" err: ", err)
		} else {
			output(blockCount)
		}
	},
}

//GetOrphansTotalCmd Get the total of all orphans
var GetOrphansTotalCmd = &cobra.Command{
	Use:     "getOrphansTotal",
	Short:   "getOrphansTotal; Get the total of all orphans",
	Aliases: []string{"GetOrphansTotal", "getorphanstotal", "getorphansblockcount"},
	Example: `
	getOrphansTotal 
	`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		params := []interface{}{}
		blockCount, err := getResString("getOrphansTotal", params)
		if err != nil {
			log.Error(cmd.Use+" err: ", err)
		} else {
			output(blockCount)
		}
	},
}

//GetBlockHashCmd get block hash by number
var GetBlockHashCmd = &cobra.Command{
	Use:     "getBlockHash {number}",
	Short:   "getBlockHash {number}; get block hash by number",
	Aliases: []string{"getblockhash", "GetBlockHash", "getBlockhash"},
	Example: `
	getBlockHash 100 
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
	Use:     "getBlock {number|hash} [verbose] [inclTx], [fullTx]",
	Aliases: []string{"getblock", "GetBlock"},
	Short:   "getBlock {number|hash} [verbose]; verbose: defalut false,show block detail,get block by number or hash",
	Example: `
getBlock 100 false

getBlock 100

getBlock 000000e4c6b7f5b89827711d412957bfff5c51730df05c2eedd1352468313eca

getBlock 000000e4c6b7f5b89827711d412957bfff5c51730df05c2eedd1352468313eca true
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

		verbose := false
		if len(args) > 1 {
			verbose, err = strconv.ParseBool(args[1])
			if err != nil {
				fmt.Println("verbose bool true or false", err)
				return
			}
		}

		inclTx := true
		if len(args) > 2 {
			inclTx, err = strconv.ParseBool(args[2])
			if err != nil {
				fmt.Println("inclTx bool true or false", err)
				return
			}
		}

		fullTx := true
		if len(args) > 3 {
			fullTx, err = strconv.ParseBool(args[3])
			if err != nil {
				fmt.Println("fullTx bool true or false", err)
				return
			}
		}

		getBlockParam := []interface{}{blockHash, verbose, inclTx, fullTx}

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
	Use:     "getBlockHashByRange {start} {end}",
	Aliases: []string{"getblockhashbyrange", "GetBlockHashByRange", "getBlockhashByRange", "gethash"},
	Short:   "getBlockHashByRange {start} {end};Return the hash range of block from 'start' to 'end'(exclude self)",
	Long: `
	getBlockHashByRange {start} {end};Return the hash range of block from 'start' to 'end'(exclude self)
	if 'end' is equal to zero, 'start' is the number that from the last block to the Gen
	if 'start' is greater than or equal to 'end', it will just return the hash of 'start'`,
	Example: `
	getBlockHashByRange 5 22
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
	Use:     "getBlockByOrder {order} {verbose} {inclTx} {fullTx}",
	Aliases: []string{"getblockbyorder", "GetBlockByOrder"},
	Short:   "getBlockByOrder {order} {verbose,default false} {inclTx,default true} {fullTx,default true}",
	Example: `
	getBlockByOrder 10 true
	`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		order, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Error(cmd.Use + " err: block order is not int")
			return
		}

		verbose := false
		if len(args) > 1 {
			verbose, err = strconv.ParseBool(args[1])
			if err != nil {
				log.Error(cmd.Use + " err: verbose bool")
				return
			}
		}

		inclTx := true
		if len(args) > 2 {
			inclTx, err = strconv.ParseBool(args[2])
			if err != nil {
				log.Error(cmd.Use + " err: inclTx bool")
				return
			}
		}

		fullTx := true
		if len(args) > 3 {
			fullTx, err = strconv.ParseBool(args[3])
			if err != nil {
				log.Error(cmd.Use + " err: fullTx bool")
				return
			}
		}

		params := []interface{}{order, verbose, inclTx, fullTx}

		blockInfo, err := getResString("getBlockByOrder", params)
		if err != nil {
			log.Error(cmd.Use+" err: ", err)
		} else {
			output(blockInfo)
		}
	},
}

//GetBlockByIDCmd get block hash by number
var GetBlockByIDCmd = &cobra.Command{
	Use:     "getBlockByID {id} {verbose} {inclTx} {fullTx}",
	Aliases: []string{"GetBlockByID", "getblockbyid"},
	Short:   "getBlockByID {id} {verbose,default false} {inclTx,default true} {fullTx,default true}",
	Example: `
	getBlockByID 10
	`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ID, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Error(cmd.Use + " err: block ID is not int")
			return
		}

		verbose := false
		if len(args) > 1 {
			verbose, err = strconv.ParseBool(args[1])
			if err != nil {
				log.Error(cmd.Use + " err: verbose bool")
				return
			}
		}

		inclTx := true
		if len(args) > 2 {
			inclTx, err = strconv.ParseBool(args[2])
			if err != nil {
				log.Error(cmd.Use + " err: inclTx bool")
				return
			}
		}

		fullTx := true
		if len(args) > 3 {
			fullTx, err = strconv.ParseBool(args[3])
			if err != nil {
				log.Error(cmd.Use + " err: fullTx bool")
				return
			}
		}

		params := []interface{}{ID, verbose, inclTx, fullTx}

		blockInfo, err := getResString("getBlockByID", params)
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
	Aliases: []string{"getbestblockhash", "GetBestBlockHash"},
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
	Use:     "getBlockHeader {number|hash} [verbose]",
	Aliases: []string{"getblockheader", "GetBlockHeader"},
	Short:   "getBlockHeader {number|hash} [verbose];verbose:bool,show detail,defalut true; get block by number or hash",
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
				log.Error("block number is not int or hash wrong: ", err)
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
				log.Error("verbose bool true or false", err)
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
	Short:   "isOnMainChain {hash}; query whether a given block is on the main chain",
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
	Short:   "getMainChainHeight: Return the current height of DAG main chain",
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
	Short:   "getBlockWeight: the weight of block",
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

//IsBlueCmd 0:not blue;  1：blue  2：Cannot confirm
var IsBlueCmd = &cobra.Command{
	Use:     "isBlue {hash}",
	Short:   "isBlue {hash},0:not blue;  1：blue  2：Cannot confirm",
	Aliases: []string{"isBlue", "isblue", "IsBlue"},
	Example: `
	isBlue 0000006c77a308846e0e0759bef5ebe0dbf4d49f345b08bdda24642efcc0cb91
	`,
	Run: func(cmd *cobra.Command, args []string) {

		params := []interface{}{args[0]}

		isOn, err := getResString("isBlue", params)
		if err != nil {
			log.Error(cmd.Use+" err: ", err)
		} else {
			fmt.Println(strings.Trim(isOn, "\""))
		}
	},
}
