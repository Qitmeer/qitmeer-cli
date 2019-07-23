package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/HalalChain/qitmeer-cli/commands"
	"github.com/HalalChain/qitmeer-cli/rpc/client"
)

var (
	configFile string
	rpcCfg     *client.Config
	debug      bool
	format     bool
)

func init() {
	InitFlag()
	commands.RootCmd.PersistentPreRunE = MergeConfig
}

func main() {
	if err := commands.RootCmd.Execute(); err != nil {
		log.Error("cmd execute err: ", err)
		os.Exit(1)
	}
	return
}

// InitFlag init cmd flags
func InitFlag() {

	rootCmd := commands.RootCmd

	rpcCfg := &client.Config{}

	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "config.toml", "config file path")

	rootCmd.PersistentFlags().StringVarP(&rpcCfg.RPCServer, "server", "s", "127.0.0.1:18131", "RPC server to connect to")
	rootCmd.PersistentFlags().StringVarP(&rpcCfg.RPCUser, "user", "u", "", "RPC username")
	rootCmd.PersistentFlags().StringVarP(&rpcCfg.RPCPassword, "password", "P", "", "RPC password")
	rootCmd.PersistentFlags().StringVar(&rpcCfg.RPCCert, "cert", "", "RPC server certificate file path")

	rootCmd.PersistentFlags().BoolVar(&rpcCfg.NoTLS, "notls", true, "Do not verify tls certificates (not recommended!)")
	rootCmd.PersistentFlags().BoolVar(&rpcCfg.TLSSkipVerify, "skipverify", true, "Do not verify tls certificates (not recommended!)")

	rootCmd.PersistentFlags().StringVar(&rpcCfg.Proxy, "proxy", "", "Connect via SOCKS5 proxy (eg. 127.0.0.1:9050)")
	rootCmd.PersistentFlags().StringVar(&rpcCfg.ProxyUser, "proxyuser", "", "Username for proxy server")
	rootCmd.PersistentFlags().StringVar(&rpcCfg.ProxyPass, "proxypass", "", "Password for proxy server")

	rootCmd.PersistentFlags().StringVar(&rpcCfg.Timeout, "timeout", "30s", "rpc timeout,s:second h:hour m:minute")

	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "debug print log")
	rootCmd.PersistentFlags().BoolVar(&format, "format", false, "print json format")
}

// MergeConfig merge config file and flags
func MergeConfig(cmd *cobra.Command, args []string) error {

	type Config struct {
		client.Config
		Debug  bool
		Format bool
	}

	cfgFromFile := &Config{}

	if cmd.Flag("config").Changed {
		_, decodeErr := toml.DecodeFile(configFile, cfgFromFile)
		if decodeErr != nil {
			return fmt.Errorf("config file err: %s", decodeErr)
		}
	}

	if cmd.Flag("server").Changed {
		cfgFromFile.RPCServer = rpcCfg.RPCServer
	}
	if cmd.Flag("user").Changed {
		cfgFromFile.RPCUser = rpcCfg.RPCUser
	}
	if cmd.Flag("password").Changed {
		cfgFromFile.RPCPassword = rpcCfg.RPCPassword
	}
	if cmd.Flag("cert").Changed {
		cfgFromFile.RPCCert = rpcCfg.RPCCert
	}
	if cmd.Flag("notls").Changed {
		cfgFromFile.NoTLS = rpcCfg.NoTLS
	}
	if cmd.Flag("skipverify").Changed {
		cfgFromFile.TLSSkipVerify = rpcCfg.TLSSkipVerify
	}

	if cmd.Flag("proxy").Changed {
		cfgFromFile.Proxy = rpcCfg.Proxy
	}
	if cmd.Flag("proxyuser").Changed {
		cfgFromFile.ProxyUser = rpcCfg.ProxyUser
	}
	if cmd.Flag("proxypass").Changed {
		cfgFromFile.ProxyPass = rpcCfg.ProxyPass
	}

	if cmd.Flag("timeout").Changed {
		cfgFromFile.Timeout = rpcCfg.Timeout
	}

	if cmd.Flag("debug").Changed {
		cfgFromFile.Debug = debug
	}
	if cmd.Flag("format").Changed {
		cfgFromFile.Format = format
	}

	if cfgFromFile.Debug {
		log.SetLevel(log.TraceLevel)
	}
	log.SetFormatter(&log.TextFormatter{
		DisableTimestamp: true,
	})

	commands.RPCCfg = &cfgFromFile.Config
	commands.Format = cfgFromFile.Format

	//save
	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(cfgFromFile); err != nil {
		log.Fatal(err)
	}

	return ioutil.WriteFile(configFile, buf.Bytes(), 0666)
}
