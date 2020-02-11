package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

var (
	isServer bool
	h        bool
	group    string
)

const (
	maxTransferUnitSize = 1500
)

func init() {
	flag.BoolVar(&h, "h", false, "This usage.")
	flag.BoolVar(&isServer, "s", false, "Start as server. Without this option, it's started as receiver.")
	flag.StringVar(&group, "g", "239.0.0.1:12345", "Address of multicast group.")
}

//Receive joins a multicast group to receive datagrams.
func Receive(mgroup string) {
	log.Printf("Start to join multicast group: %s\n", mgroup)
	addr, err := net.ResolveUDPAddr("udp4", mgroup)
	if err != nil {
		log.Fatal(err)
	}

	con, err := net.ListenMulticastUDP("udp4", nil, addr)
	if err != nil {
		log.Fatal("Failed to jion multicase group!", err)
	}
	if con.SetReadBuffer(maxTransferUnitSize) != nil {
		log.Fatal("Failed to set datagram size.")
	}
	for {
		b := make([]byte, maxTransferUnitSize)
		n, src, err := con.ReadFromUDP(b)
		if err != nil {
			log.Fatal("ReadFromUDP failed:", err)
		}
		log.Println(n, "bytes read from", src)
		log.Println(string(b[:n]))
	}
}

//Send creates a server to broadcast in multicast group.
func Send(mgroup string) {
	log.Printf("Start server in multicast group:%s\n", mgroup)
	addr, err := net.ResolveUDPAddr("udp4", mgroup)
	if err != nil {
		log.Fatal("Failed to start !", err)
	}

	conn, err := net.DialUDP("udp4", nil, addr)
	if err != nil {
		log.Fatal("Failed to start !", err)
	}

	for {
		conn.Write([]byte("hello, world\n"))
		log.Println("Send hello message.")
		time.Sleep(1 * time.Second)
	}
}

func usage() {
	fmt.Println(`Usage: ` + os.Args[0] + ` [-s] [-g xxx.xxx.xxx.xxx:xxxxx]
	
	Options:`)
	flag.PrintDefaults()
}

func main() {
	flag.Parse()

	if h {
		usage()
		return
	}

	if isServer {
		Send(group)
	} else {
		Receive(group)
	}
}
