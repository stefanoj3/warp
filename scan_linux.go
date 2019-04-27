package warp

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"regexp"
)

var arpRegexp = regexp.MustCompile(`^([\d\.]+)\s+dev\s+(\w+)\s+\w+\s+([a-f0-9:]{17})\s+\w+$`)

func entriesFromARP(executor ARPCommandExecutor) ([]Entry, error) {
	output, err := executor()
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

		entry, err := entryFromRawData(m[1], m[3], m[2])
		if err != nil {
			return nil, fmt.Errorf("linux.entriesFromARP: failed to create entry(%s): %s", line, err.Error())
		}

		entries = append(entries, entry)
	}

	return entries, nil

}
