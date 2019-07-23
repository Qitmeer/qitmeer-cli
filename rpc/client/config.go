package client

// Config rpc client config
type Config struct{
	RPCServer     string
	RPCUser       string
	RPCPassword   string	
	RPCCert       string
	NoTLS         bool
	TLSSkipVerify bool

	Proxy     string
	ProxyUser string
	ProxyPass string

	Timeout string
}