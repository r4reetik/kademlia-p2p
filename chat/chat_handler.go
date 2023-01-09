package chat

import (
	"fmt"
	"kademlia-p2p/domain"
	"kademlia-p2p/kademlia"
	"kademlia-p2p/utils"
	"strings"

	"kademlia-p2p/network"
)

func handleMessage(target string, message string, peer *kademlia.Peer) {
	targetNodeId := utils.StringToBytes(target)

	targetNode := peer.FindNode(targetNodeId)
	for targetNode.ID != targetNodeId {
		findMsg := domain.NewMessage(4, []byte("/"+utils.BytesToString(peer.Node.ID)+"/tcp/"+utils.NetIPToString(peer.Node.IP)+":"+utils.Uint16ToString(peer.Node.Port)+"+"+target))
		err := network.SendMessage(findMsg, utils.NetIPToString(targetNode.IP), utils.Uint16ToString(targetNode.Port))
		if err != nil {
			fmt.Println(err)
		}
		targetNode = peer.FindNode(targetNodeId)
	}

	textMsg := domain.NewMessage(3, []byte("/"+utils.BytesToString(peer.Node.ID)+"/tcp/"+utils.NetIPToString(peer.Node.IP)+":"+utils.Uint16ToString(peer.Node.Port)+"+"+message))
	err := network.SendMessage(textMsg, utils.NetIPToString(targetNode.IP), utils.Uint16ToString(targetNode.Port))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Message sent to ", target)
}

func handlePing(target string, peer *kademlia.Peer) {
	splitedByColon := strings.Split(target, ":")
	if len(splitedByColon) < 2 {
		fmt.Println("Invalid message format")
		return
	}

	ip := splitedByColon[0]
	port := splitedByColon[1]

	pingMsg := domain.NewMessage(1, []byte("/"+utils.BytesToString(peer.Node.ID)+"/tcp/"+utils.NetIPToString(peer.Node.IP)+":"+utils.Uint16ToString(peer.Node.Port)))
	err := network.SendMessage(pingMsg, ip, port)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Ping sent to ", target)
}

func parseInput(input string, peer *kademlia.Peer) {
	splitedBySlash := strings.Split(input, "/")
	if len(splitedBySlash) < 2 {
		fmt.Println("Invalid message format")
		return
	}

	switch splitedBySlash[0] {
	case "@message":
		if len(splitedBySlash) < 3 {
			fmt.Println("Invalid message format")
			return
		}
		handleMessage(splitedBySlash[1], splitedBySlash[2], peer)
	case "@ping":
		if len(splitedBySlash) < 2 {
			fmt.Println("Invalid message format")
			return
		}
		handlePing(splitedBySlash[1], peer)
	}
}

func ListenChat(peer *kademlia.Peer) {
	for {
		var input string
		fmt.Scanln(&input)
		go parseInput(input, peer)
	}
}
