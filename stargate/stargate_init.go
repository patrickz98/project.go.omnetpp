package stargate

import (
	"fmt"
	lg "log"
	"net"
	"os"
)

const (
	DefaultAddr = "localhost"
	DefaultPort = 8889
)

var (
	rendezvousAddr *net.UDPAddr
	config         Config
	broadcast      = &net.UDPAddr{
		IP:   net.IPv4(239, 11, 22, 33),
		Port: 10077,
	}
)

func SetConfig(conf Config) {

	config = conf

	addr := fmt.Sprintf("%s:%d", config.Addr, config.Port)

	var err error
	rendezvousAddr, err = net.ResolveUDPAddr("udp", addr)
	if err != nil {
		log.Fatalln(err)
	}
}

var log *lg.Logger

func init() {
	log = lg.New(os.Stderr, "Stargate ", lg.LstdFlags|lg.Lshortfile)

	SetConfig(Config{
		Addr: DefaultAddr,
		Port: DefaultPort,
	})
}
