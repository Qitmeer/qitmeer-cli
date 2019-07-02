package main

import (
	"fmt"

	"github.com/HalalChain/qitmeer-cli/server"
)

func main() {

	server.Start()

	ch := make(chan int)

	<-ch

	fmt.Println("wallet")
}
