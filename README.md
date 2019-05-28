# qitmeer-cli
The command line utility of Qitmeer

Configuration file config.toml  will maked automatic 

# useage

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
  getmempool           get mempool
  getrawtransaction    getrawtransaction
  help                 Help about any command
  sendrawtransaction   sendRawTransaction
  txSign               txSign private_key raw_tx

Flags:
      --c string           RPC server certificate file path
      --conf string        RPC username (default "config.toml")
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
  -u, --user string        RPC username

Use "qitmeer-cli [command] --help" for more information about a command.
```

