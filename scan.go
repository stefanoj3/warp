package warp

import (
	"fmt"
	"net"
)

// An ARPScanner is responsible for scanning the ARP table
type ARPScanner struct {
}

// NewARPScanner returns a new instance of ARPScanner
func NewARPScanner() ARPScanner {
	return ARPScanner{}
}

// Scan tries to scan the ARP table and return the entries found if possible
func (s ARPScanner) Scan() ([]Entry, error) {
	entries, err := entriesFromARP()
	if err != nil {
		return nil, err
	}

	// filter by interface ??

	return entries, nil
}

func entryFromRawData(rawIP, rawMAC, rawInterface string) (Entry, error) {
	ip := net.ParseIP(rawIP)
	if ip == nil {
		return Entry{}, fmt.Errorf("entryFromRawData: failed to parse ip: %s", rawIP)
	}

	mac, err := net.ParseMAC(rawMAC)
	if err != nil {
		return Entry{}, fmt.Errorf("entryFromRawData: failed to parse MAC address (%s): %s", rawMAC, err.Error())
	}

	i, err := net.InterfaceByName(rawInterface)
	if err != nil {
		return Entry{}, fmt.Errorf("entryFromRawData: failed to parse interface name(%s): %s", rawInterface, err.Error())
	}

	return Entry{
		IP:        ip,
		MAC:       mac,
		Interface: *i,
	}, nil
}
