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
	iterface string
	locAddr  string
)

const (
	maxTransferUnitSize = 1500
)

func init() {
	flag.BoolVar(&h, "h", false, "This usage.")
	flag.BoolVar(&isServer, "s", false, "Start as server. Without this option, it's started as receiver.")
	flag.StringVar(&group, "g", "239.0.0.1:12345", "Address of multicast group.")
	flag.StringVar(&iterface, "i", "", "Interface name to listen multicast on.")
	flag.StringVar(&locAddr, "l", "", "Local address to send datagram.")
}

//receive joins a multicast group to receive datagrams.
func receive() {
	log.Printf("Start to join multicast group: %s\n", group)
	addr, err := net.ResolveUDPAddr("udp", group)
	if err != nil {
		log.Fatal(err)
	}
	var i *net.Interface
	if len(iterface) > 0 {
		i, err = net.InterfaceByName(iterface)
		i.MulticastAddrs()
		if err != nil {
			log.Fatal(err)
		}
	}
	con, err := net.ListenMulticastUDP("udp", i, addr)
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

//send creates a server to broadcast in multicast group.
func send() {
	log.Printf("Start server in multicast group:%s\n", group)
	addr, err := net.ResolveUDPAddr("udp", group)
	if err != nil {
		log.Fatal("Failed to resolve udp address !", err)
	}
	var laddr *net.UDPAddr
	if len(locAddr) > 0 {
		laddr, err = net.ResolveUDPAddr("udp", locAddr)
		if err != nil {
			log.Fatal("Failed to resolve udp address !", err)
		}
	}

	conn, err := net.DialUDP("udp", laddr, addr)
	if err != nil {
		log.Fatal("Failed to start !", err)
	}

	for {
		time.Sleep(time.Second)
		log.Println("Send hello message from ", conn.LocalAddr())
		conn.Write([]byte("hello, world\n"))

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
		send()
	} else {
		receive()
	}

}

