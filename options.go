package warp

import "os/exec"

// Option is the type used to represent additional configuration for the ARPScanner
type Option func(*ARPScanner) error

// NoopExecutorOption applies a noop implementation of the Executor
func NoopExecutorOption(s *ARPScanner) error {
	s.arpCommandExecutor = func() ([]byte, error) {
		return []byte{}, nil
	}
	return nil
}

// CustomCommandExecutorOption applies an implementation of the executor that executes the command provided
// For example it can be used to fetch the ARP table of a remote machine
func CustomCommandExecutorOption(cmd string, args ...string) Option {
	return func(s *ARPScanner) error {
		s.arpCommandExecutor = func() ([]byte, error) {
			return exec.Command(cmd, args...).CombinedOutput()
		}
		return nil
	}
}
