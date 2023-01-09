package network

import (
	"fmt"
	"kademlia-p2p/domain"
	"kademlia-p2p/kademlia"
	"kademlia-p2p/utils"
	"strings"
)

func OpTypeHandler(msg domain.Message, peer *kademlia.Peer) {
	switch msg.OPtype {
	case 1:
		handlePing(msg, peer)
	case 2:
		handlePong(msg, peer)
	case 3:
		handleText(msg, peer)
	case 4:
		handleFindNode(msg, peer)
	case 5:
		handleNearNode(msg, peer)
	case 6:
		handleFoundNode(msg, peer)
	}
}

func handlePing(msg domain.Message, peer *kademlia.Peer) {
	splitedData := strings.Split(string(msg.Data), "/")
	id := splitedData[1]

	address := splitedData[3]
	splitedAddress := strings.Split(address, ":")
	ip := splitedAddress[0]
	port := splitedAddress[1]

	remoteNode := kademlia.NewNode(utils.StringToBytes(id), utils.StringToNetIP(ip), utils.StringToUint16(port))
	peer.AddNode(remoteNode)
	fmt.Printf("Ping received from ----- ID: %s, IP: %s, Port: %s\n", id, ip, port)

	pongMsg := domain.NewMessage(2, []byte("/"+utils.BytesToString(peer.Node.ID)+"/tcp/"+utils.NetIPToString(peer.Node.IP)+":"+utils.Uint16ToString(peer.Node.Port)))
	err := SendMessage(pongMsg, ip, port)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Pong sent to ----- ID: %s, IP: %s, Port: %s\n", id, ip, port)
}

func handlePong(msg domain.Message, peer *kademlia.Peer) {
	splitedData := strings.Split(string(msg.Data), "/")
	id := splitedData[1]

	address := splitedData[3]
	splitedAddress := strings.Split(address, ":")
	ip := splitedAddress[0]
	port := splitedAddress[1]

	remoteNode := kademlia.NewNode(utils.StringToBytes(id), utils.StringToNetIP(ip), utils.StringToUint16(port))
	peer.AddNode(remoteNode)

	fmt.Printf("Pong received from ----- ID: %s, IP: %s, Port: %s\n", id, ip, port)
}

func handleText(msg domain.Message, peer *kademlia.Peer) {
	splitedByPlus := strings.Split(string(msg.Data), "+")

	splitedData := strings.Split(splitedByPlus[0], "/")
	id := splitedData[1]

	address := splitedData[3]
	splitedAddress := strings.Split(address, ":")
	ip := splitedAddress[0]
	port := splitedAddress[1]

	remoteNode := kademlia.NewNode(utils.StringToBytes(id), utils.StringToNetIP(ip), utils.StringToUint16(port))
	peer.AddNode(remoteNode)

	msgReceived := splitedByPlus[1]

	fmt.Printf("Text received from ----- ID: %s, IP: %s, Port: %s\n", id, ip, port)
	fmt.Println("Message received: ", string(msgReceived))
}

func handleFindNode(msg domain.Message, peer *kademlia.Peer) {
	splitedByPlus := strings.Split(string(msg.Data), "+")

	splitedData := strings.Split(splitedByPlus[0], "/")
	id := splitedData[1]

	address := splitedData[3]
	splitedAddress := strings.Split(address, ":")
	ip := splitedAddress[0]
	port := splitedAddress[1]

	remoteNode := kademlia.NewNode(utils.StringToBytes(id), utils.StringToNetIP(ip), utils.StringToUint16(port))
	peer.AddNode(remoteNode)

	targetNodeId := splitedByPlus[1]

	resultNode := peer.FindNode(utils.StringToBytes(targetNodeId))

	if resultNode.ID == utils.StringToBytes(targetNodeId) {
		foundNodeMsg := domain.NewMessage(6, []byte("/"+utils.BytesToString(peer.Node.ID)+"/tcp/"+utils.NetIPToString(peer.Node.IP)+":"+utils.Uint16ToString(peer.Node.Port)+"+"+"/"+utils.BytesToString(resultNode.ID)+"/tcp/"+utils.NetIPToString(resultNode.IP)+":"+utils.Uint16ToString(resultNode.Port)))
		err := SendMessage(foundNodeMsg, ip, port)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("Found node sent to ----- ID: %s, IP: %s, Port: %s\n", id, ip, port)
		fmt.Printf("Found node ----- ID: %s, IP: %s, Port: %s\n", utils.BytesToString(resultNode.ID), utils.NetIPToString(resultNode.IP), utils.Uint16ToString(resultNode.Port))
	} else {
		nearNodeMsg := domain.NewMessage(5, []byte("/"+utils.BytesToString(peer.Node.ID)+"/tcp/"+utils.NetIPToString(peer.Node.IP)+":"+utils.Uint16ToString(peer.Node.Port)+"+"+"/"+utils.BytesToString(resultNode.ID)+"/tcp/"+utils.NetIPToString(resultNode.IP)+":"+utils.Uint16ToString(resultNode.Port)+"+"+targetNodeId))
		err := SendMessage(nearNodeMsg, ip, port)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("Near node sent to ----- ID: %s, IP: %s, Port: %s\n", id, ip, port)
		fmt.Printf("Near node ----- ID: %s, IP: %s, Port: %s\n", utils.BytesToString(resultNode.ID), utils.NetIPToString(resultNode.IP), utils.Uint16ToString(resultNode.Port))
	}
}

func handleNearNode(msg domain.Message, peer *kademlia.Peer) {
	splitedByPlus := strings.Split(string(msg.Data), "+")

	splitedData := strings.Split(splitedByPlus[0], "/")
	id := splitedData[1]

	address := splitedData[3]
	splitedAddress := strings.Split(address, ":")
	ip := splitedAddress[0]
	port := splitedAddress[1]

	remoteNode := kademlia.NewNode(utils.StringToBytes(id), utils.StringToNetIP(ip), utils.StringToUint16(port))
	peer.AddNode(remoteNode)

	nearNodeData := splitedByPlus[1]
	targetNodeId := splitedByPlus[2]

	splitedData = strings.Split(nearNodeData, "/")
	id = splitedData[1]

	address = splitedData[3]
	splitedAddress = strings.Split(address, ":")
	ip = splitedAddress[0]
	port = splitedAddress[1]

	nearNodeId := utils.StringToBytes(id)
	nearNodeIp := utils.StringToNetIP(ip)
	nearNodePort := utils.StringToUint16(port)

	nearNode := kademlia.NewNode(nearNodeId, nearNodeIp, nearNodePort)
	peer.AddNode(nearNode)

	fmt.Printf("Near node received from ----- ID: %s, IP: %s, Port: %s\n", id, ip, port)
	fmt.Printf("Near node ----- ID: %s, IP: %s, Port: %s\n", utils.BytesToString(nearNode.ID), utils.NetIPToString(nearNode.IP), utils.Uint16ToString(nearNode.Port))

	findNodeMsg := domain.NewMessage(4, []byte("/"+utils.BytesToString(peer.Node.ID)+"/tcp/"+utils.NetIPToString(peer.Node.IP)+":"+utils.Uint16ToString(peer.Node.Port)+"+"+targetNodeId))
	err := SendMessage(findNodeMsg, utils.NetIPToString(nearNodeIp), utils.Uint16ToString(nearNodePort))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Find node sent to ----- ID: %s, IP: %s, Port: %s\n", utils.BytesToString(nearNodeId), utils.NetIPToString(nearNodeIp), utils.Uint16ToString(nearNodePort))
}

func handleFoundNode(msg domain.Message, peer *kademlia.Peer) {
	splitedByPlus := strings.Split(string(msg.Data), "+")

	splitedData := strings.Split(splitedByPlus[0], "/")
	id := splitedData[1]

	address := splitedData[3]
	splitedAddress := strings.Split(address, ":")
	ip := splitedAddress[0]
	port := splitedAddress[1]

	remoteNode := kademlia.NewNode(utils.StringToBytes(id), utils.StringToNetIP(ip), utils.StringToUint16(port))
	peer.AddNode(remoteNode)

	fmt.Printf("Found node received from ----- ID: %s, IP: %s, Port: %s\n", id, ip, port)
	fmt.Printf("Found node ----- ID: %s, IP: %s, Port: %s\n", utils.BytesToString(remoteNode.ID), utils.NetIPToString(remoteNode.IP), utils.Uint16ToString(remoteNode.Port))

	foundNodeData := splitedByPlus[1]

	splitedData = strings.Split(foundNodeData, "/")
	id = splitedData[1]

	address = splitedData[3]
	splitedAddress = strings.Split(address, ":")
	ip = splitedAddress[0]
	port = splitedAddress[1]

	foundNodeId := utils.StringToBytes(id)
	foundNodeIp := utils.StringToNetIP(ip)
	foundNodePort := utils.StringToUint16(port)

	foundNode := kademlia.NewNode(foundNodeId, foundNodeIp, foundNodePort)
	peer.AddNode(foundNode)
}
