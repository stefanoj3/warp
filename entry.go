package warp

import "net"

// Entry represents an entry in the ARP table
type Entry struct {
	MAC       net.HardwareAddr
	IP        net.IP
	Interface net.Interface
}
