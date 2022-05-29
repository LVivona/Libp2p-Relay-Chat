package main

import (
	"context"
	"log"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/peer"

	relayv1 "github.com/libp2p/go-libp2p/p2p/protocol/circuitv1/relay"
)

func main() {
	run()
}

func run() {

	//config, err := ParseFlags()
	_, cancel := context.WithCancel(context.Background())
	defer cancel()
	relay, err := libp2p.New(libp2p.DisableRelay())
	if err != nil {
		panic(err)

	}

	_, err = relayv1.NewRelay(relay)
	if err != nil {
		log.Printf("Failed to instantiate h2 relay: %v", err)
		return
	}

	Relayinfo := peer.AddrInfo{
		ID:    relay.ID(),
		Addrs: relay.Addrs(),
	}
	/*
		kademliaDHT, err := dht.New(ctx, relay)
		if err != nil {
			panic(err)
		}

		logger.Debug("Bootstrapping the DHT")
		if err = kademliaDHT.Bootstrap(ctx); err != nil {
			panic(err)
		}

		var wg sync.WaitGroup
		for _, peerAddr := range config.BootstrapPeers {
			peerinfo, _ := peer.AddrInfoFromP2pAddr(peerAddr)
			wg.Add(1)
			go func() {
				defer wg.Done()
				if err := relay.Connect(ctx, *peerinfo); err != nil {
					logger.Warning(err)
				} else {
					logger.Info("Connection established with bootstrap node:", *peerinfo)
				}
			}()
		}
		wg.Wait()
	*/
	log.Printf("ID: %s", Relayinfo.ID)
	for _, addr := range Relayinfo.Addrs {
		log.Printf("Adder: %s", addr)
	}

	select {}
}
