package commands

import (
	"fmt"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(CreateRawTransactionCmd)
	RootCmd.AddCommand(GetRawTransactionCmd)
	RootCmd.AddCommand(DecodeRawTransactionCmd)
	RootCmd.AddCommand(SendRawTransactionCmd)
	RootCmd.AddCommand(TxSignCmd)
	RootCmd.AddCommand(GetUtxoCmd)
	RootCmd.AddCommand(GetNodeInfoCmd)
	RootCmd.AddCommand(GetPeerInfoCmd)

}

//GetRawTransactionCmd getrawtransaction
var GetRawTransactionCmd = &cobra.Command{
	Use:     "getrawtransaction {tx_hash} {verbose bool,show detail,defalut true}",
	Aliases: []string{"tx", "getrawtx", "getRawTransaction"},
	Short:   "getrawtransaction",
	Example: `
		getrawtransaction 000000e4c6b7f5b89827711d412957bfff5c51730df05c2eedd1352468313eca
		getrawtransaction 000000e4c6b7f5b89827711d412957bfff5c51730df05c2eedd1352468313eca true
	`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		var err error

		var txHash string
		var verbose bool = true

		txHash = args[0]

		if len(args) > 1 {
			verbose, err = strconv.ParseBool(args[1])
			if err != nil {
				fmt.Println("verbose bool true or false", err)
				return
			}
		}

		var txInfo string
		txInfo, err = getResString("getRawTransaction", []interface{}{txHash, verbose})
		if err != nil {
			log.Error(cmd.Use+" err: ", err)
		} else {
			fmt.Println(txInfo)
		}
	},
}

//CreateRawTransactionCmd CreateRawTransactionCmd
var CreateRawTransactionCmd = &cobra.Command{
	Use:     "createrawtransaction {inTxid:vout}... {toAddr:amount}...",
	Aliases: []string{"createrawtx", "createRawTransaction"},
	Short:   "createRawTransaction",
	Example: `
		createRawTransaction b203ff6ba4f39ecf846a103c17f15e35afcbd229f72ad1a9f0a90f07a7535dff:2 RmFFQV5FsuKFU5b4sBjGvpDd6P183iMZRcT:20.3
	`,
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

		var err error

		type fromTx struct {
			Txid string `json:"txid"`
			Vout int64  `json:"vout"`
		}

		fromTxs := []fromTx{}
		toAddr := make(map[string]float64)

		for _, k := range args {
			kParts := strings.Split(k, ":")
			if len(kParts[0]) > 40 {
				vout, _ := strconv.ParseInt(kParts[1], 10, 64)
				fromTxs = append(fromTxs, fromTx{Txid: kParts[0], Vout: vout})
			} else {
				amount, _ := strconv.ParseFloat(kParts[1], 64)
				toAddr[kParts[0]] = amount
			}
		}

		/*
					 [
			            {
			                "txid": "b203ff6ba4f39ecf846a103c17f15e35afcbd229f72ad1a9f0a90f07a7535dff",
			                "vout":2
			            }
			        ],
			        {
			            "RmFFQV5FsuKFU5b4sBjGvpDd6P183iMZRcT": 449.8
					}
		*/

		var rawTx string
		rawTx, err = getResString("createRawTransaction", []interface{}{fromTxs, toAddr})
		if err != nil {
			log.Error(cmd.Use+" err: ", err)
		} else {
			fmt.Println(rawTx)
		}
	},
}

//DecodeRawTransactionCmd DecodeRawTransactionCmd
var DecodeRawTransactionCmd = &cobra.Command{
	Use:     "decoderawtransaction {raw_tx}",
	Aliases: []string{"decoderawtx", "decodeRawTransaction"},
	Short:   "decodeRawTransaction",
	Example: `
		decoderawtx xx
	`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		var err error
		var tx string
		tx, err = getResString("decodeRawTransaction", []interface{}{args[0]})
		if err != nil {
			log.Error(cmd.Use+" err: ", err)
		} else {
			fmt.Println(tx)
		}
	},
}

//SendRawTransactionCmd SendRawTransactionCmdl
var SendRawTransactionCmd = &cobra.Command{
	Use:     "sendrawtransaction {raw_tx} {allow_high_fee bool,defalut false}",
	Aliases: []string{"sendRawTx", "sendrawtx", "sendRawTransaction"},
	Short:   "sendRawTransaction",
	Example: `
		sendRawTransaction raw_tx
		sendRawTransaction raw_tx true
	`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		var allowHighFee bool = false
		if len(args) > 1 {
			allowHighFee, err = strconv.ParseBool(args[1])
			if err != nil {
				fmt.Println("allowHighFee bool true or false", err)
				return
			}
		}

		var rs string
		rs, err = getResString("sendRawTransaction", []interface{}{args[0], allowHighFee})
		if err != nil {
			log.Error(cmd.Use+" err: ", err)
		} else {
			fmt.Println(rs)
		}
	},
}

//TxSignCmd TxSignCmd
var TxSignCmd = &cobra.Command{
	Use:   "txSign {private_key} {raw_tx}",
	Short: "txSign private_key raw_tx",
	Example: `
	txSign private_key raw_tx
	`,
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		var rs string
		rs, err = getResString("txSign", []interface{}{args[0], args[1]})
		if err != nil {
			log.Error(cmd.Use+" err: ", err)
		} else {
			fmt.Println(rs)
		}
	},
}

//GetUtxoCmd GetUtxoCmd
var GetUtxoCmd = &cobra.Command{
	Use:     "getUtxo {tx_hash} {vout index} {include_mempool,bool,defalut true}",
	Short:   "getUtxo tx_hash vout include_mempool,",
	Aliases: []string{"getutxo"},
	Example: `
		getUtxo xx
	`,
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		params := []interface{}{}
		params = append(params, args[0])

		var vout int64
		vout, err = strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			fmt.Println("vout not number", err)
			return
		}
		params = append(params, vout)

		var includeMempool bool = true
		if len(args) > 2 {
			includeMempool, err = strconv.ParseBool(args[2])
			if err != nil {
				fmt.Println("include_mempool true or false", err)
				return
			}
		}
		params = append(params, includeMempool)

		var tx string

		tx, err = getResString("getUtxo", params)
		if err != nil {
			log.Error(cmd.Use+" err: ", err)
		} else {
			output(tx)
		}
	},
}

//GetNodeInfoCmd GetNodeInfo
var GetNodeInfoCmd = &cobra.Command{
	Use:     "getNodeInfo",
	Short:   "getNodeInfo",
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
	Short:   "getPeerInfo",
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
