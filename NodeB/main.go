package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
)

var (
	id   = flag.String("id", "", "ID")
	addr = flag.String("Address", "", "Address")
)

func handlerMeow(s network.Stream) {
	log.Println("Meow! It worked!")
	s.Close()
}

func handleStream(stream network.Stream) {
	log.Println("Got a new stream!")

	// Create a buffer stream for non blocking read and write.
	rw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))

	go readData(rw)
	go writeData(rw)

	// 'stream' will stay open until you close it (or the other side closes it).
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

		rw.WriteString(fmt.Sprintf("%s : %s\n", h3.ID().ShortString(), sendData))
		rw.Flush()
	}
}

var h3 host.Host

func main() {
	flag.Parse()
	h3, err := libp2p.New(libp2p.ListenAddrs(), libp2p.EnableRelay())
	if err != nil {
		log.Printf("Failed to create h3: %v", err)
		return
	}

	log.Printf("ID: %s", h3.ID())
	for _, addr := range h3.Addrs() {
		log.Printf("Adder: %s", addr)
	}

	realyaddr, err := peer.AddrInfoFromString(fmt.Sprintf("%s/p2p/%s", *id, *addr))
	if err != nil {
		panic(err)
	}

	if err := h3.Connect(context.Background(), *realyaddr); err != nil {
		log.Printf("Failed to connect h3 and h2: %v", err)
		return
	}

	log.Printf("Connected to: %s", realyaddr.ID)

	h3.SetStreamHandler("/cats", handlerMeow)
	h3.SetStreamHandler("/chat", handleStream)

	select {}
}
