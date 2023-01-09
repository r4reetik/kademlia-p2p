package main

import (
	"fmt"
	"sync"

	"kademlia-p2p/chat"
	"kademlia-p2p/kademlia"
	"kademlia-p2p/network"
	"kademlia-p2p/utils"
)

func main() {
	var wg sync.WaitGroup

	nodeId := utils.GenerateNodeId("r4reetik")
	ip, err := utils.GetOutboundIP()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("IP: ", ip)

	node := kademlia.NewNode(nodeId, ip, 17080)

	peer := kademlia.NewPeer(node)

	wg.Add(1)
	go network.ListenTCP(peer)

	wg.Add(1)
	go chat.ListenChat(peer)

	wg.Wait()
}
