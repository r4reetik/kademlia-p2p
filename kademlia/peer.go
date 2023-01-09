package kademlia

type Peer struct {
	Node         *Node
	RoutingTable *RoutingTable
}

func NewPeer(node *Node) *Peer {
	return &Peer{
		Node:         node,
		RoutingTable: NewRoutingTable(),
	}
}

func (p *Peer) AddNode(node *Node) int {
	return p.RoutingTable.Add(node, p.Node)
}

func (p *Peer) FindNode(targetNodeID [20]byte) *Node {
	return p.RoutingTable.Find(targetNodeID, p.Node)
}