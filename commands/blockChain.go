package commands

import (
	"fmt"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	blockChainCmds := []*cobra.Command{
		CreateRawTransactionCmd,
		GetRawTransactionCmd,
		DecodeRawTransactionCmd,
		SendRawTransactionCmd,
		TxSignCmd,
		GetUtxoCmd,
		GetNodeInfoCmd,
		GetPeerInfoCmd,
	}
	RootCmd.AddCommand(blockChainCmds...)
	RootSubCmdGroups["blockChain"] = blockChainCmds

}

//GetRawTransactionCmd getrawtransaction
var GetRawTransactionCmd = &cobra.Command{
	Use:     "getRawTransaction {tx_hash} [verbose]",
	Aliases: []string{"getrawtransaction", "GetRawTransaction", "getRawTx", "getrawtx", "GetRawTx"},
	Short:   "getRawTransaction {tx_hash} [verbose]; verbose: bool,show detail,defalut true",
	Example: `
getRawTransaction 000000e4c6b7f5b89827711d412957bfff5c51730df05c2eedd1352468313eca

getRawTransaction 000000e4c6b7f5b89827711d412957bfff5c51730df05c2eedd1352468313eca true
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
				log.Error("verbose bool true or false", err)
				return
			}
		}

		var txInfo string
		txInfo, err = getResString("getRawTransaction", []interface{}{txHash, verbose})
		if err != nil {
			log.Error(cmd.Use+" err: ", err)
		} else {
			output(txInfo)
		}
	},
}

//CreateRawTransactionCmd CreateRawTransactionCmd
var CreateRawTransactionCmd = &cobra.Command{
	Use:     "createRawTransaction {inTxid:vout}... {toAddr:amount}... {lockTime}",
	Aliases: []string{"createrawtransaction", "CreateRawTransaction", "createRawTx", "createrawtx", "CreateRawTx"},
	Short:   "createRawTx {inTxid:vout}... {toAddr:amount}... {lockTime},crate raw transaction",
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

		var lockTime int64 = -1

		for _, k := range args {
			kParts := strings.Split(k, ":")

			//lockTime
			if len(kParts) == 1 {
				lockTime, err = strconv.ParseInt(k, 10, 64)
				if err != nil {
					log.Error("lockTime err: ", err)
					return
				}
				continue
			}

			//from
			if len(kParts[0]) > 40 {
				vout, _ := strconv.ParseInt(kParts[1], 10, 64)
				fromTxs = append(fromTxs, fromTx{Txid: kParts[0], Vout: vout})
				continue
			}

			//to
			amount, _ := strconv.ParseFloat(kParts[1], 64)
			toAddr[kParts[0]] = amount
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

		params := []interface{}{fromTxs, toAddr}
		if lockTime != -1 {
			params = append(params, lockTime)
		}

		var rawTx string
		rawTx, err = getResString("createRawTransaction", params)
		if err != nil {
			log.Error(cmd.Use+" err: ", err)
		} else {
			output(rawTx)
		}
	},
}

//DecodeRawTransactionCmd DecodeRawTransactionCmd
var DecodeRawTransactionCmd = &cobra.Command{
	Use:     "decodeRawTransaction {raw_tx}",
	Aliases: []string{"decoderawtransaction", "DecodeRawTransaction", "decoderawtx", "DecodeRawTx", "decodeRawTx"},
	Short:   "decodeRawTransaction {raw_tx}",
	Example: `
decodeRawTransaction 0100000002ecca2b71379753f64cd57ac611835272f0381142b4f290affbbf21a4544b4f3c000000006b483045022100c042dfc287ca6aa02fac67dd291b6e0df67af7e1328ba0c932e7bf21b5b2b050022030d8e528d62412085b1f3d9bd1887a80b277b5d6e0d31d9efddf91e5ca207016012103eb2609d195f15b5976d50b119796e9448afd5503f051aaf085aadd46b29ec6a0ffffffff08aa3d2bdfec7da453f1b61cd0f991bc5c2bdd6e5d471f092029581ca8bc1d53000000006a4730440220778d685685d65d3866863a009dd4c62fc7f799825ef9835987dcd08453a401e1022023cf2073d3bb9552871d7fee777785f28c6693e4c0bfa4db1ebd870c57a58546012103fb7863f5d5c8ade2d1dfb1e2171765c5ffeb79012cf034d79d7fe6bb90b32f12ffffffff0400562183000000001976a91420db62bb6e4907083a524df9620bf37922d4c29a88ac005c995a000000001976a914f0ab9022524a632e55beb282a3b766923702367d88ac00f15365000000001976a91459534ae5dde9008b4b53ee38a24baceafef2c8ac88ac3ea4186a000000001976a9142b35121f5554181ec0e7c2a84cb6fa6fe115ac6788ac00000000
	`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		var err error
		var tx string
		tx, err = getResString("decodeRawTransaction", []interface{}{args[0]})
		if err != nil {
			log.Error(cmd.Use+" err: ", err)
		} else {
			output(tx)
		}
	},
}

//SendRawTransactionCmd SendRawTransactionCmdl
var SendRawTransactionCmd = &cobra.Command{
	Use:     "sendRawTransaction {sign_raw_tx} {allow_high_fee}",
	Aliases: []string{"sendrawtransaction", "SendRawTransaction", "sendRawTx", "sendrawtx", "SendRawTx"},
	Short:   "sendRawTransaction {sign_raw_tx} {allow_high_fee}; allow_high_fee: default false; send sing_raw_tx to network",
	Example: `
sendRawTransaction 0100000001ff5d53a7070fa9f0a9d12af729d2cbaf355ef1173c106a84cf9ef3a46bff03b202000000ffffffff01005504790a0000001976a914627777996288556166614462639988446255776688ac000000000000000001000000000000000000000000ffffffff6b483045022100dced4d67dd74647d0036077ee5b435838934377b1d296dd9da852772911e3be2022063dd346bd812a894968b8acacead7e7beff48947657a82f1e8f9c38876d4c905012103aba0a09f5b44138a46a2e5d26b8659923d84c4ba9437e22c3828cac43d0edb49

sendRawTransaction 0100000001ff5d53a7070fa9f0a9d12af729d2cbaf355ef1173c106a84cf9ef3a46bff03b202000000ffffffff01005504790a0000001976a914627777996288556166614462639988446255776688ac000000000000000001000000000000000000000000ffffffff6b483045022100dced4d67dd74647d0036077ee5b435838934377b1d296dd9da852772911e3be2022063dd346bd812a894968b8acacead7e7beff48947657a82f1e8f9c38876d4c905012103aba0a09f5b44138a46a2e5d26b8659923d84c4ba9437e22c3828cac43d0edb49 true
	`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		var allowHighFee bool = false
		if len(args) > 1 {
			allowHighFee, err = strconv.ParseBool(args[1])
			if err != nil {
				log.Error("allow_high_fee bool true or false", err)
				return
			}
		}

		var rs string
		rs, err = getResString("sendRawTransaction", []interface{}{args[0], allowHighFee})
		if err != nil {
			log.Error(cmd.Use+" err: ", err)
		} else {
			output(rs)
		}
	},
}

//TxSignCmd TxSignCmd
var TxSignCmd = &cobra.Command{
	Use:     "txSign {private_key} {raw_tx}",
	Short:   "txSign {private_key} {raw_tx}; sign rawTx",
	Aliases: []string{"txsign", "TxSign", "signRawTx", "signrawtx", "SignRawTx"},
	Example: `
//txSign {private_key} {raw_tx}

txSign 2ad045c0df865c8f84479ea06adf00cbbfec705fb9402ea117ce2ef242a9d260 0100000001ff5d53a7070fa9f0a9d12af729d2cbaf355ef1173c106a84cf9ef3a46bff03b202000000ffffffff01005504790a0000001976a914627777996288556166614462639988446255776688ac000000000000000001000000000000000000000000ffffffff00
	`,
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		var rs string
		rs, err = getResString("txSign", []interface{}{args[0], args[1]})
		if err != nil {
			log.Error(cmd.Use+" err: ", err)
		} else {
			output(rs)
		}
	},
}

//GetUtxoCmd GetUtxoCmd
var GetUtxoCmd = &cobra.Command{
	Use:     "getUtxo {tx_hash} {vout} [include_mempool],bool,defalut true}",
	Short:   "getUtxo {tx_hash} {vout} [include_mempool]; vout:index of the output; include_mempool: default=true,include the mempool , get information about an unspent transaction output",
	Aliases: []string{"getutxo", "GetUtxo"},
	Example: `
getutxo a97cf4d67bbe5ce57d1d2f4fc18ae2ee19e1048cbb1a14d8d94273bfef83f371 0 true
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
