package warp

// ARPCommandExecutor runs the command needed to fetch the ARP table from the OS
// and returns its combined standard output and standard error.
type ARPCommandExecutor func() ([]byte, error)
