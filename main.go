package main

// "flag"
// "fmt"
// "log"
// "log/slog"
// "net"

// "github.com/tidwall/resp"

const defaultListenAddr = ":5001"

type Config struct {
	ListenAddr string
}

type Message struct {
	cmd  Command
	peer *Peer
}
