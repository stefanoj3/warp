package warp

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
