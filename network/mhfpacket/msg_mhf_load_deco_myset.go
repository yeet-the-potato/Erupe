package mhfpacket

import (
	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfLoadDecoMyset represents the MSG_MHF_LOAD_DECO_MYSET
type MsgMhfLoadDecoMyset struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfLoadDecoMyset) Opcode() network.PacketID {
	return network.MSG_MHF_LOAD_DECO_MYSET
}

// Parse parses the packet from binary
func (m *MsgMhfLoadDecoMyset) Parse(bf *byteframe.ByteFrame) error {
	panic("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfLoadDecoMyset) Build(bf *byteframe.ByteFrame) error {
	panic("Not implemented")
}