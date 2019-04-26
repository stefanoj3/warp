package warp_test

import (
	"fmt"

	"github.com/stefanoj3/warp"
)

func ExampleARPScanner_Scan() {
	scanner := warp.NewARPScanner()
	entries, err := scanner.Scan()
	if err != nil {
		panic(err)
	}

	for _, entry := range entries {
		fmt.Println(entry)
	}
}
