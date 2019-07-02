package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/spf13/cobra"

	"github.com/HalalChain/qitmeer-cli/rpc/client"
)

// Config cli config file
type Config struct {
	ConfigFile string

	SimNet  bool
	TestNet bool

	RPC *client.Config
}

var cfg = &Config{
	RPC: &client.Config{}}

var rpcClient *client.RPCClient

var rootCmd = &cobra.Command{
	Use:               "cli",
	Long:              `cli is a RPC tool for noxd`,
	PersistentPreRunE: makeConfigFile,
}

func init() {

	//flags
	rootCmd.PersistentFlags().StringVar(&cfg.ConfigFile, "conf", "config.toml", "RPC username")
	rootCmd.PersistentFlags().StringVarP(&cfg.RPC.RPCUser, "user", "u", "", "RPC username")
	rootCmd.PersistentFlags().StringVarP(&cfg.RPC.RPCPassword, "password", "P", "", "RPC password")
	rootCmd.PersistentFlags().StringVarP(&cfg.RPC.RPCServer, "server", "s", "127.0.0.1:18131", "RPC server to connect to")
	rootCmd.PersistentFlags().StringVar(&cfg.RPC.RPCCert, "c", "", "RPC server certificate file path")
	rootCmd.PersistentFlags().BoolVar(&cfg.RPC.NoTLS, "notls", true, "Do not verify tls certificates (not recommended!)")
	rootCmd.PersistentFlags().BoolVar(&cfg.RPC.TLSSkipVerify, "skipverify", true, "Do not verify tls certificates (not recommended!)")
	rootCmd.PersistentFlags().StringVar(&cfg.RPC.Proxy, "proxy", "", "Connect via SOCKS5 proxy (eg. 127.0.0.1:9050)")
	rootCmd.PersistentFlags().StringVar(&cfg.RPC.ProxyUser, "proxyuser", "", "Username for proxy server")
	rootCmd.PersistentFlags().StringVar(&cfg.RPC.ProxyPass, "proxypass", "", "Password for proxy server")
	rootCmd.PersistentFlags().BoolVar(&cfg.TestNet, "testnet", false, "Connect to testnet")
	rootCmd.PersistentFlags().BoolVar(&cfg.SimNet, "simnet", false, "Connect to the simulation test network")

	//cmds
	rootCmd.AddCommand(GenerateCmd)
	rootCmd.AddCommand(GetBlockCountCmd)
	rootCmd.AddCommand(GetBlockTemplateCmd)
	rootCmd.AddCommand(GetBlockHashCmd)
	rootCmd.AddCommand(GetBlockCmd)
	rootCmd.AddCommand(GetMempoolCmd)
	rootCmd.AddCommand(GetRawTransactionCmd)
	rootCmd.AddCommand(CreateRawTransactionCmd)
	rootCmd.AddCommand(DecodeRawTransactionCmd)
	rootCmd.AddCommand(SendRawTransactionCmd)
	rootCmd.AddCommand(TxSignCmd)
	rootCmd.AddCommand(GetUtxoCmd)
}

//
func makeConfigFile(cmd *cobra.Command, args []string) error {
	cfg2 := &Config{}
	_, decodeErr := toml.DecodeFile(cfg.ConfigFile, cfg2)

	if decodeErr != nil {
		fmt.Println("config file err:", decodeErr)
	} else {
		if !cmd.Flag("user").Changed {
			cfg.RPC.RPCUser = cfg2.RPC.RPCUser
		}
		if !cmd.Flag("password").Changed {
			cfg.RPC.RPCPassword = cfg2.RPC.RPCPassword
		}
		if !cmd.Flag("server").Changed {
			cfg.RPC.RPCServer = cfg2.RPC.RPCServer
		}
	}

	//params.MainNetParams.DefaultPort
	//	preCfg := cfg

	// Multiple networks can't be selected simultaneously.
	numNets := 0
	if cfg.TestNet {
		numNets++
	}
	if cfg.SimNet {
		numNets++
	}
	if numNets > 1 {
		return fmt.Errorf("loadConfig: %s", "one of the testnet and simnet")
	}

	//save
	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(cfg); err != nil {
		log.Fatal(err)
	}

	err := ioutil.WriteFile(cfg.ConfigFile, buf.Bytes(), 0666)
	if err != nil {
		return err
	}

	rpcClient, err = client.NewRPCClient(cfg.RPC)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return
}
