package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"net"
	"strconv"
)

func GenerateNodeId(uname string) [20]byte {
	sha1 := sha1.New()
	sha1.Write([]byte(uname))
	var hash [20]byte
	copy(hash[:], sha1.Sum(nil))

	return hash
}

func GetOutboundIP() (net.IP, error) {
	conn, err := net.Dial("udp", "1.1.1.1:80")
	if err != nil {
		return nil, errors.New("tcp-listener.go: No outbound IP address for this host")
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP, nil
}

func StringToBytes(str string) [20]byte {
	byteString, err := hex.DecodeString(str)
	if err != nil {
		fmt.Println(err)
	}
	var bytes [20]byte
	copy(bytes[:], byteString[:])
	return bytes
}

func StringToNetIP(str string) net.IP {
	return net.ParseIP(str)
}

func StringToUint16(str string) uint16 {
	number, err := strconv.ParseUint(str, 10, 16)
	if err != nil {
		fmt.Println(err)
	}
	return uint16(number)
}

func BytesToString(bytes [20]byte) string {
	return hex.EncodeToString(bytes[:])
}

func NetIPToString(ip net.IP) string {
	return ip.String()
}

func Uint16ToString(num uint16) string {
	return fmt.Sprintf("%v", num)
}
