package kademlia

import (
	"net"
)

const K = 20

type Bucket []*Node

func NewBucket() Bucket {
	return make(Bucket, K)
}

func ping(node *Node) bool {
	conn, err := net.Dial("tcp", node.Address())
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

func (b *Bucket) Add(node *Node) {
	for _, n := range *b {
		if n.ID == node.ID {
			n.IP = node.IP
			n.Port = node.Port
			return
		}
	}

	if len(*b) == K && !ping((*b)[0]) {
		*b = (*b)[1:]
		*b = append(*b, node)
	} else {
		*b = append(*b, node)
	}
}

type RoutingTable struct {
	Buckets [NODE_ID_LENGTH]Bucket
}

func NewRoutingTable() *RoutingTable {
	return &RoutingTable{}
}

func getBucketIndex(distance []byte) int {
	index := 0
	for _, val := range distance {
		if val == 0 {
			index += 8
			continue
		}
		for i := 0; i < 8; i++ {
			if val&0x80 == 0 {
				index++
				val = val << 1
			} else {
				return index
			}
		}
	}
	return index
}

func (rt *RoutingTable) Add(node *Node, self *Node) int {
	if node.ID == self.ID {
		return 0
	}

	distance := node.DistanceTo(self)
	bucketIndex := getBucketIndex(distance)
	rt.Buckets[bucketIndex].Add(node)
	return bucketIndex
}

func (rt *RoutingTable) Find(targetNodeID [20]byte, self *Node) *Node {
	targetNode := NewNode(targetNodeID, nil, 0)
	distance := targetNode.DistanceTo(self)
	bucketIndex := getBucketIndex(distance)
	bucket := rt.Buckets[bucketIndex]

	for _, node := range bucket {
		if node.ID == targetNodeID {
			return node
		}
	}

	for i := bucketIndex; i >= 0; i-- {
		bucket := rt.Buckets[i]
		for _, node := range bucket {
			return node
		}
	}

	for i := bucketIndex + 1; i < NODE_ID_LENGTH; i++ {
		bucket := rt.Buckets[i]
		for _, node := range bucket {
			return node
		}	
	}

	return nil
}
