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

// Config file
type Config struct {
	configFile string

	Debug  bool //print log
	Format bool //output by json format

	client.Config
}

var preCfg *Config

func init() {
	log.SetFormatter(&log.TextFormatter{
		DisableTimestamp: true,
	})
	bindFlags()
}

func main() {
	commands.RootCmd.PersistentPreRunE = LoadConfig

	if err := commands.RootCmd.Execute(); err != nil {
		log.Error("cmd execute err: ", err)
		os.Exit(1)
	}
	return
}

// InitFlags init flags
func bindFlags() {
	preCfg = &Config{}
	gFlags := commands.RootCmd.PersistentFlags()

	gFlags.StringVarP(&preCfg.configFile, "config", "c", "cli.toml", "config file path")

	gFlags.StringVarP(&preCfg.RPCServer, "server", "s", "127.0.0.1:18131", "RPC server to connect to")
	gFlags.StringVarP(&preCfg.RPCUser, "user", "u", "", "RPC username")
	gFlags.StringVarP(&preCfg.RPCPassword, "password", "P", "", "RPC password")
	gFlags.StringVar(&preCfg.RPCCert, "cert", "", "RPC server certificate file path")

	gFlags.BoolVar(&preCfg.NoTLS, "notls", true, "Do not verify tls certificates (not recommended!)")
	gFlags.BoolVar(&preCfg.TLSSkipVerify, "skipverify", true, "Do not verify tls certificates (not recommended!)")

	gFlags.StringVar(&preCfg.Proxy, "proxy", "", "Connect via SOCKS5 proxy (eg. 127.0.0.1:9050)")
	gFlags.StringVar(&preCfg.ProxyUser, "proxyuser", "", "Username for proxy server")
	gFlags.StringVar(&preCfg.ProxyPass, "proxypass", "", "Password for proxy server")

	gFlags.StringVar(&preCfg.Timeout, "timeout", "30s", "rpc timeout,s:second h:hour m:minute")

	gFlags.BoolVar(&preCfg.Debug, "debug", false, "debug print log")
	gFlags.BoolVar(&preCfg.Format, "format", false, "print json format")
}

// LoadConfig merge config file and flags
func LoadConfig(cmd *cobra.Command, args []string) (err error) {
	//cmd := commands.RootCmd

	fileCfg := &Config{}

	fileCfg.configFile = preCfg.configFile

	_, decodeErr := toml.DecodeFile(fileCfg.configFile, fileCfg)
	if decodeErr != nil {
		return fmt.Errorf("config file err: %s", decodeErr)
	}

	if cmd.Flag("debug").Changed {
		fileCfg.Debug = preCfg.Debug
	}
	if fileCfg.Debug {
		log.SetLevel(log.TraceLevel)
	}

	if cmd.Flag("server").Changed {
		fileCfg.RPCServer = preCfg.RPCServer
	}
	if cmd.Flag("user").Changed {
		fileCfg.RPCUser = preCfg.RPCUser
	}
	if cmd.Flag("password").Changed {
		fileCfg.RPCPassword = preCfg.RPCPassword
	}
	if cmd.Flag("cert").Changed {
		fileCfg.RPCCert = preCfg.RPCCert
	}
	if cmd.Flag("notls").Changed {
		fileCfg.NoTLS = preCfg.NoTLS
	}
	if cmd.Flag("skipverify").Changed {
		fileCfg.TLSSkipVerify = preCfg.TLSSkipVerify
	}

	if cmd.Flag("proxy").Changed {
		fileCfg.Proxy = preCfg.Proxy
	}
	if cmd.Flag("proxyuser").Changed {
		fileCfg.ProxyUser = preCfg.ProxyUser
	}
	if cmd.Flag("proxypass").Changed {
		fileCfg.ProxyPass = preCfg.ProxyPass
	}

	if cmd.Flag("timeout").Changed {
		fileCfg.Timeout = preCfg.Timeout
	}

	if cmd.Flag("debug").Changed {
		fileCfg.Debug = preCfg.Debug
	}
	if cmd.Flag("format").Changed {
		fileCfg.Format = preCfg.Format
	}

	log.Debug("fileCfg: ", *fileCfg)

	commands.RPCCfg = &fileCfg.Config
	commands.Format = fileCfg.Format

	//save
	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(*fileCfg); err != nil {
		log.Fatal(err)
	}

	return ioutil.WriteFile(fileCfg.configFile, buf.Bytes(), 0666)
}

// FileExists reports whether the named file or directory exists.
func FileExists(filePath string) (bool, error) {
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
