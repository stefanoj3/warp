package warp

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"regexp"
	"strings"
)

var arpRegexp = regexp.MustCompile(`^[^\d\.]+([\d\.]+).+\s+([a-f0-9:]{11,17})\s+on\s+([^\s]+)\s+.+$`)

func entriesFromARP(executor ARPCommandExecutor) ([]Entry, error) {
	output, err := executor()
	if err != nil {
		return nil, fmt.Errorf("darwin.entriesFromARP: failed to scan arp table: %s", err.Error())
	}

	var entries []Entry

	reader := bufio.NewReader(bytes.NewBuffer(output))

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, fmt.Errorf("darwin.entriesFromARP: failed to read buffer: %s", err.Error())
		}

		m := arpRegexp.FindStringSubmatch(string(line))
		if len(m) != 4 {
			continue
		}

		entry, err := entryFromRawData(m[1], normalizeMac(m[2]), m[3])
		if err != nil {
			return nil, fmt.Errorf("darwin.entriesFromARP: failed to create entry(%s): %s", line, err.Error())
		}

		entries = append(entries, entry)
	}

	return entries, nil

}

// normalizeMac adds missing leading zeroes - the ARP table in OSX can contain MAC addresses that golang
// is unable to parse like: f1:ec:4e:20:f2:6 which is missing the leading zero for the last portion of the
// address (:6 instead of :06), this method would translate the MAC to f1:ec:4e:20:f2:06
func normalizeMac(rawMac string) string {
	parts := strings.Split(rawMac, ":")

	for i, part := range parts {
		if len(part) < 2 {
			parts[i] = "0" + part
		}
	}

	return strings.Join(parts, ":")
}
