package warp

import "os/exec"

// LocalExecutor runs the command needed to fetch the ARP table from the local machine
func LocalExecutor() ([]byte, error) {
	return exec.Command("arp", "-n", "-a").CombinedOutput()
}
