package warp_test

import (
	"testing"
	"time"

	"github.com/stefanoj3/warp"
)

func TestArpScannerCanPerformAScan(t *testing.T) {
	scanner := warp.NewArpScanner()

	_, err := scanner.Scan()
	if err != nil {
		t.Errorf("no error expected when scanning")
	}
}

func TestArpScannerCanPerformAScansInParallel(t *testing.T) {
	scanner := warp.NewArpScanner()

	for i := 0; i < 30; i++ {
		go func() {
			_, err := scanner.Scan()
			if err != nil {
				t.Errorf("no error expected when scanning")
			}
		}()
	}

	time.Sleep(500 * time.Millisecond)
}
