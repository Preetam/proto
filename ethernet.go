package protodecode

import (
	"net"
)

type EthernetFrame struct {
	Source      net.HardwareAddr `json:"source"`
	Destination net.HardwareAddr `json:"destination"`
	VlanTag     uint32           `json:"vlanTag"`
	EtherType   uint16           `json:"etherType"`
	Payload     []byte           `json:"payload"`
}

func DecodeEthernet(b []byte) EthernetFrame {
	frame := EthernetFrame{}

	i := 0 // just a helper for indexing

	frame.Source = net.HardwareAddr(b[i : i+6])
	i += 6
	frame.Destination = net.HardwareAddr(b[i : i+6])
	i += 6

	// Check for a VLAN tag
	if b[i] == 0x81 && b[i+1] == 0x00 {
		i += 2
		frame.VlanTag = uint32(0x81<<24) | uint32(0x00<<16) | uint32(b[i]<<8) | uint32(b[i+1])
		i += 2
	}

	// Check the Ethernet type
	if t := uint16(b[i])<<8 | uint16(b[i+1]); t >= 1536 { // if 1500 or less, then it's the payload length
		frame.EtherType = t
	}

	i += 2

	frame.Payload = b[i:]

	return frame
}