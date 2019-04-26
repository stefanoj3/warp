package warp

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"regexp"
)

var arpRegexp = regexp.MustCompile(`^[^\d\.]+([\d\.]+).+\s+([a-f0-9\-]{11,17})\s+.+$`)

func entriesFromARP() ([]Entry, error) {
	output, err := exec.Command("arp", "-a").CombinedOutput()
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
				IP:        m[0],
				Interface: m[1],
				MAC:       m[2],
			},
		)
	}

	return entries, nil

}
