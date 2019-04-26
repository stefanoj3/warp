package warp

type ArpScanner struct {
}

func NewArpScanner() ArpScanner {
	return ArpScanner{}
}

func (s ArpScanner) Scan() ([]Entry, error) {
	entries, err := entriesFromArp()
	if err != nil {
		return nil, err
	}

	// filter by interface ??

	return entries, nil
}
