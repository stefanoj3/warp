package warp

// An Entry represents an entry in the ARP table
type Entry struct {
	MAC       string
	IP        string
	Interface string
}
