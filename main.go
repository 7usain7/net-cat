package main

import (
	"fmt"
	"net-cat/funcs"
	"os"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	} else if len(os.Args) == 2 {
		funcs.DEFULT_PORT = os.Args[1]
	}

	funcs.ServerHandler(funcs.DEFULT_PORT)
}
