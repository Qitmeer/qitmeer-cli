package main

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/BurntSushi/toml"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var cfg = &Config{}

/*  */
func init() {
	rootCmd.PersistentFlags().StringVarP(&cfg.ConfigFile, "config", "c", "config.toml", "config file path")

	rootCmd.PersistentFlags().StringVarP(&cfg.RPCUser, "user", "u", "", "RPC username")
	rootCmd.PersistentFlags().StringVarP(&cfg.RPCPassword, "password", "P", "", "RPC password")
	rootCmd.PersistentFlags().StringVarP(&cfg.RPCServer, "server", "s", "127.0.0.1:18131", "RPC server to connect to")
	rootCmd.PersistentFlags().StringVar(&cfg.RPCCert, "cert", "", "RPC server certificate file path")

	rootCmd.PersistentFlags().BoolVar(&cfg.NoTLS, "notls", true, "Do not verify tls certificates (not recommended!)")
	rootCmd.PersistentFlags().BoolVar(&cfg.TLSSkipVerify, "skipverify", true, "Do not verify tls certificates (not recommended!)")

	rootCmd.PersistentFlags().StringVar(&cfg.Proxy, "proxy", "", "Connect via SOCKS5 proxy (eg. 127.0.0.1:9050)")
	rootCmd.PersistentFlags().StringVar(&cfg.ProxyUser, "proxyuser", "", "Username for proxy server")
	rootCmd.PersistentFlags().StringVar(&cfg.ProxyPass, "proxypass", "", "Password for proxy server")

	rootCmd.PersistentFlags().BoolVar(&cfg.TestNet, "testnet", false, "Connect to testnet")
	rootCmd.PersistentFlags().BoolVar(&cfg.SimNet, "simnet", false, "Connect to the simulation test network")

	rootCmd.PersistentFlags().BoolVar(&cfg.Debug, "debug", false, "debug print log")
	rootCmd.PersistentFlags().StringVar(&cfg.Timeout, "timeout", "30s", "rpc timeout,s:second h:hour m:minute")
}

type Config struct {
	ConfigFile string

	RPCUser       string
	RPCPassword   string
	RPCServer     string
	RPCCert       string
	NoTLS         bool
	TLSSkipVerify bool

	Proxy     string
	ProxyUser string
	ProxyPass string

	TestNet bool
	SimNet  bool

	Debug   bool
	Timeout string
}

//
func rootCmdPreRun(cmd *cobra.Command, args []string) error {

	cfgFromFile := &Config{}

	if cmd.Flag("config").Changed {
		_, decodeErr := toml.DecodeFile(cfg.ConfigFile, cfgFromFile)
		if decodeErr != nil {
			return fmt.Errorf("config file err: %s", decodeErr)
		}
	}

	if cmd.Flag("user").Changed {
		cfgFromFile.RPCUser = cfg.RPCUser
	}
	if cmd.Flag("password").Changed {
		cfgFromFile.RPCPassword = cfg.RPCPassword
	}
	if cmd.Flag("server").Changed {
		cfgFromFile.RPCServer = cfg.RPCServer
	}
	if cmd.Flag("cert").Changed {
		cfgFromFile.RPCCert = cfg.RPCCert
	}
	if cmd.Flag("notls").Changed {
		cfgFromFile.NoTLS = cfg.NoTLS
	}
	if cmd.Flag("skipverify").Changed {
		cfgFromFile.TLSSkipVerify = cfg.TLSSkipVerify
	}

	if cmd.Flag("proxy").Changed {
		cfgFromFile.Proxy = cfg.Proxy
	}
	if cmd.Flag("proxyuser").Changed {
		cfgFromFile.ProxyUser = cfg.ProxyUser
	}
	if cmd.Flag("proxypass").Changed {
		cfgFromFile.ProxyPass = cfg.ProxyPass
	}

	if cmd.Flag("debug").Changed {
		cfgFromFile.Debug = cfg.Debug
	}
	if cmd.Flag("timeout").Changed {
		cfgFromFile.Timeout = cfg.Timeout
	}

	// Multiple networks can't be selected simultaneously.
	numNets := 0
	if cfg.TestNet {
		cfgFromFile.TestNet = true
		numNets++
	}
	if cfg.SimNet {
		cfgFromFile.SimNet = true
		numNets++
	}
	if numNets > 1 {
		return fmt.Errorf("network: %s", "one of the testnet and simnet")
	}

	cfg = cfgFromFile

	if cfg.Debug {
		log.SetLevel(log.TraceLevel)
	}
	log.SetFormatter(&log.TextFormatter{
		DisableTimestamp: true,
	})

	//save
	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(cfg); err != nil {
		log.Fatal(err)
	}

	return ioutil.WriteFile(cfg.ConfigFile, buf.Bytes(), 0666)
}
