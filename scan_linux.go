package warp

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"regexp"
)

var arpRegexp = regexp.MustCompile(`^([\d\.]+)\s+dev\s+(\w+)\s+\w+\s+([a-f0-9:]{17})\s+\w+$`)

func entriesFromARP() ([]Entry, error) {
	output, err := exec.Command("ip", "neigh").CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("linux.entriesFromARP: failed to scan arp table: %s", err.Error())
	}

	var entries []Entry

	reader := bufio.NewReader(bytes.NewBuffer(output))

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, fmt.Errorf("linux.entriesFromARP: failed to read buffer: %s", err.Error())
		}

		m := arpRegexp.FindStringSubmatch(string(line))
		if len(m) != 4 {
			continue
		}

		entries = append(
			entries,
			Entry{
				IP:        m[1],
				Interface: m[2],
				MAC:       m[3],
			},
		)
	}

	return entries, nil

}
