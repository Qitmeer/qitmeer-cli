module github.com/HalalChain/qitmeer-cli

go 1.12

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/samuel/go-socks v0.0.0-20130725190102-f6c5f6a06ef6
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/cobra v0.0.4
)

replace (
	golang.org/x/crypto v0.0.0-20181203042331-505ab145d0a9 => github.com/golang/crypto v0.0.0-20181203042331-505ab145d0a9
	golang.org/x/sys v0.0.0-20181205085412-a5c9d58dba9a => github.com/golang/sys v0.0.0-20181205085412-a5c9d58dba9a
	golang.org/x/text v0.3.0 => github.com/golang/text v0.3.0
	gopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405 => github.com/go-check/check v0.0.0-20161208181325-20d25e280405
	gopkg.in/yaml.v2 v2.2.2 => github.com/go-yaml/yaml v0.0.0-20181115110504-51d6538a90f8
)
