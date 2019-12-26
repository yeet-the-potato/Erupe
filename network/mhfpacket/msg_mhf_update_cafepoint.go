package mhfpacket

import (
	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfUpdateCafepoint represents the MSG_MHF_UPDATE_CAFEPOINT
type MsgMhfUpdateCafepoint struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfUpdateCafepoint) Opcode() network.PacketID {
	return network.MSG_MHF_UPDATE_CAFEPOINT
}

// Parse parses the packet from binary
func (m *MsgMhfUpdateCafepoint) Parse(bf *byteframe.ByteFrame) error {
	panic("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfUpdateCafepoint) Build(bf *byteframe.ByteFrame) error {
	panic("Not implemented")
}