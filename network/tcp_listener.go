package network

import (
	"fmt"
	"kademlia-p2p/domain"
	"kademlia-p2p/kademlia"
	"net"
)

func HandleConnection(conn net.Conn, peer *kademlia.Peer) {
	defer conn.Close()

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
	}

	msg := domain.Message{}
	err = msg.Decode(buf[:n])
	if err != nil {
		fmt.Println(err)
	}

	OpTypeHandler(msg, peer)
}

func ListenTCP(peer *kademlia.Peer) error {
	listener, err := net.ListenTCP("tcp", &net.TCPAddr{IP: peer.Node.IP, Port: int(peer.Node.Port)})
	if err != nil {
		return err
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return err
		}
		go HandleConnection(conn, peer)
	}
}
