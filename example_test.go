package warp_test

import (
	"fmt"

	"github.com/stefanoj3/warp"
)

func ExampleARPScanner_Scan() {
	scanner, err := warp.NewARPScanner()
	if err != nil {
		panic(err)
	}

	entries, err := scanner.Scan()
	if err != nil {
		panic(err)
	}

	for _, entry := range entries {
		fmt.Println(entry)
	}
}

func ExampleNoopExecutorOption() {
	scanner, err := warp.NewARPScanner(
		warp.NoopExecutorOption,
	)
	if err != nil {
		panic(err)
	}

	entries, err := scanner.Scan()
	if err != nil {
		panic(err)
	}

	if len(entries) != 0 {
		panic("noop executor should produce no entries")
	}
}

func ExampleCustomCommandExecutorOption() {
	scanner, err := warp.NewARPScanner(
		warp.CustomCommandExecutorOption("echo"),
	)
	if err != nil {
		panic(err)
	}

	entries, err := scanner.Scan()
	if err != nil {
		panic(err)
	}

	if len(entries) != 0 {
		panic(
			"customer executor should produce no entries if the output is not compatible to the one of the ARP table",
		)
	}
}
