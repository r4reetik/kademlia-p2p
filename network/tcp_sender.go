package network

import (
	"kademlia-p2p/domain"
	"net"
)

func SendMessage(msg domain.Message, ip string, port string) error {
	encodedMsg, err := msg.Encode()
	if err != nil {
		return err
	}

	conn, err := net.Dial("tcp", ip+":"+port)
	if err != nil {
		return err
	}

	_, err = conn.Write(encodedMsg)
	if err != nil {
		return err
	}

	conn.Close()
	return nil
}
