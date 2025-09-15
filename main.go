package main

import (
	"fmt"
	"net-cat/funcs"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(os.Args) > 2 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	} else if len(os.Args) == 2 {
		if _, err := strconv.Atoi(args[1]); err != nil {
			fmt.Println("[USAGE]: ./TCPChat $port")
			return
		}

		funcs.DEFULT_PORT = os.Args[1]
	}

	funcs.ServerHandler(funcs.DEFULT_PORT)
}
