package kademlia

import (
	"net"
	"strconv"
)

const NODE_ID_BYTE = 20
const NODE_ID_LENGTH = 160

type Node struct {
	ID   [NODE_ID_BYTE]byte
	IP   net.IP
	Port uint16
}

func NewNode(id [NODE_ID_BYTE]byte, ip net.IP, port uint16) *Node {
	return &Node{
		ID:   id,
		IP:   ip,
		Port: port,
	}
}

func (n *Node) DistanceTo(node *Node) []byte {
	distance := make([]byte, len(n.ID))

	for i := 0; i < NODE_ID_BYTE; i++ {
		distance[i] = n.ID[i] ^ node.ID[i]
	}

	return distance
}

func (n *Node) Address() string {
	return n.IP.String() + ":" + strconv.FormatUint(uint64(n.Port), 10)
}
