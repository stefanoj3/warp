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
		return nil, fmt.Errorf("windows.entriesFromARP: failed to scan arp table: %s", err.Error())
	}

	var entries []Entry

	reader := bufio.NewReader(bytes.NewBuffer(output))

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, fmt.Errorf("windows.entriesFromARP: failed to read buffer: %s", err.Error())
		}

		m := arpRegexp.FindStringSubmatch(string(line))
		if len(m) != 3 {
			continue
		}

		entry, err := entryFromRawData(m[1], m[2], m[3])
		if err != nil {
			return nil, fmt.Errorf("windows.entriesFromARP: failed to create entry(%s): %s", line, err.Error())
		}

		entries = append(entries, entry)
	}

	return entries, nil

}
