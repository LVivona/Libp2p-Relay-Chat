package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/peerstore"
	"github.com/libp2p/go-libp2p-core/protocol"
	ma "github.com/multiformats/go-multiaddr"
)

var (
	id    = flag.String("id", "", "ID")
	addr  = flag.String("Address", "", "Address")
	dest  = flag.String("d", "", "destination")
	proto = flag.String("p", "/cat", "protocol")
)

func main() {
	flag.Parse()

	h1, err := libp2p.New(libp2p.EnableRelay())
	if err != nil {
		log.Printf("Failed to create h1: %v", err)
		return
	}

	realyaddr, err := peer.AddrInfoFromString(fmt.Sprintf("%s/p2p/%s", *id, *addr))
	if err != nil {
		panic(err)
	}
	h3, err := peer.AddrInfoFromString(*dest)
	if err != nil {
		panic(err)
	}

	if err := h1.Connect(context.Background(), *realyaddr); err != nil {
		log.Printf("Failed to connect h1 and h2: %v", err)
		return
	}
	log.Println("made it")
	r, err := ma.NewMultiaddr("/p2p/" + realyaddr.ID.Pretty() + "/p2p-circuit/ipfs/" + h3.ID.Pretty())
	if err != nil {
		log.Println(err)
		return
	}

	relayInfo := peer.AddrInfo{
		ID:    h3.ID,
		Addrs: []ma.Multiaddr{r},
	}

	h1.Peerstore().AddAddr(relayInfo.ID, relayInfo.Addrs[0], peerstore.PermanentAddrTTL)

	if err := h1.Connect(context.Background(), relayInfo); err != nil {
		log.Printf("Failed to connect h1 and h3: %v", err)
		return
	}

	s, err := h1.NewStream(context.Background(), h3.ID, protocol.ID(*proto))
	if err != nil {
		log.Println("huh, this should have worked: ", err)
		return
	}

	if *proto == "/chat" {
		rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))
		go readData(rw)
		go writeData(rw)
		select {}
	} else {
		s.Read(make([]byte, 1))
	}

}

func readData(rw *bufio.ReadWriter) {
	for {
		str, _ := rw.ReadString('\n')

		if str == "" {
			return
		}
		if str != "\n" {
			// Green console colour: 	\x1b[32m
			// Reset console colour: 	\x1b[0m
			fmt.Printf("\x1b[32m%s\x1b[0m> ", str)
		}

	}
}

func writeData(rw *bufio.ReadWriter) {
	stdReader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		sendData, err := stdReader.ReadString('\n')
		if err != nil {
			log.Println(err)
			return
		}

		rw.WriteString(fmt.Sprintf("%s\n", sendData))
		rw.Flush()
	}
}
