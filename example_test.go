package warp_test

import (
	"fmt"

	"github.com/stefanoj3/warp"
)

func ExampleArpScanner_Scan() {
	scanner := warp.NewArpScanner()
	entries, err := scanner.Scan()
	if err != nil {
		panic(err)
	}

	for _, entry := range entries {
		fmt.Println(entry)
	}
}
