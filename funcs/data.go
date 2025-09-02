package funcs

import (
	"net"
	"sync"
)

var DEFULT_PORT = "8989"

var serverLogHistory []string

var clients = make(map[net.Conn]*Client) // Stores all currently connected clients
var chanMessage = make(chan string)

var clientsMutex sync.Mutex

type Client struct {
	Name    string
	Conn    net.Conn // The actual TCP connection
	message string
}

var LinuxAscii = "Welcome to TCP-Chat!\n" +
	"         _nnnn_\n" +
	"        dGGGGMMb\n" +
	"       @p~qp~~qMb\n" +
	"       M|@||@) M|\n" +
	"       @,----.JM|\n" +
	"      JS^\\__/  qKL\n" +
	"     dZP        qKRb\n" +
	"    dZP          qKKb\n" +
	"   fZP            SMMb\n" +
	"   HZM            MMMM\n" +
	"   FqM            MMMM\n" +
	" __| \".        |\\dS\"qML\n" +
	" |    `.       | `' \\Zq\n" +
	"_)      \\.___.,|     .'\n" +
	"\\____   )MMMMMP|   .'\n" +
	"     `-'       `--'\n"
