# qitmeer-cli
The command line utility of Qitmeer

Configuration file config.toml will be made automatically

# Usage 

```
qitmeer cli is a RPC tool for the qitmeer network

Usage:
  qitmeer-cli [command]

Available Commands:
  createrawtransaction createRawTransaction
  decoderawtransaction decodeRawTransaction
  generate             generate {n}, cpu mine n blocks
  getUtxo              getUtxo tx_hash vout include_mempool,
  getblock             get block by number or hash
  getblockcount        get block count
  getblockhash         get block hash by number
  getblocktemplate     getblocktemplate
  getmempool           get mempool
  getrawtransaction    getrawtransaction
  help                 Help about any command
  sendrawtransaction   sendRawTransaction
  txSign               txSign private_key raw_tx

Flags:
      --cert string        RPC server certificate file path
  -c, --config string      config file path (default "config.toml")
      --debug              debug print log
  -h, --help               help for qitmeer-cli
      --notls              Do not verify tls certificates (not recommended!) (default true)
  -P, --password string    RPC password
      --proxy string       Connect via SOCKS5 proxy (eg. 127.0.0.1:9050)
      --proxypass string   Password for proxy server
      --proxyuser string   Username for proxy server
  -s, --server string      RPC server to connect to (default "127.0.0.1:18131")
      --simnet             Connect to the simulation test network
      --skipverify         Do not verify tls certificates (not recommended!) (default true)
      --testnet            Connect to testnet
      --timeout string     rpc timeout,s:second h:hour m:minute (default "30s")
  -u, --user string        RPC username

Use "qitmeer-cli [command] --help" for more information about a command.
```

