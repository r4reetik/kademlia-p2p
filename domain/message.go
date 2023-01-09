package domain

import (
	"encoding/binary"
)

type Message struct {
	Length uint32
	OPtype uint8
	Data   []byte
}

func NewMessage(opType uint8, data []byte) Message {
	return Message{
		Length: uint32(len(data) + 1),
		OPtype: opType,
		Data:   data,
	}
}

func (m *Message) Encode() ([]byte, error) {
	encodedMsg := make([]byte, m.Length+4)
	binary.BigEndian.PutUint32(encodedMsg, m.Length)
	encodedMsg[4] = m.OPtype
	copy(encodedMsg[5:], m.Data)
	return encodedMsg, nil
}

func (m *Message) Decode(encodedMsg []byte) error {
	m.Length = binary.BigEndian.Uint32(encodedMsg)
	m.OPtype = encodedMsg[4]
	m.Data = encodedMsg[5:]
	return nil
}
