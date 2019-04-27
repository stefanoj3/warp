package warp_test

import (
	"testing"
	"time"

	"github.com/stefanoj3/warp"
)

func TestARPScannerCanPerformAScan(t *testing.T) {
	scanner, err := warp.NewARPScanner()
	if err != nil {
		t.Fatalf("no error expected when creating a new scanner, got %s", err.Error())
	}

	_, err = scanner.Scan()
	if err != nil {
		t.Errorf("no error expected when scanning, got %s", err.Error())
	}
}

func TestARPScannerCanPerformAScansInParallel(t *testing.T) {
	scanner, err := warp.NewARPScanner()
	if err != nil {
		t.Fatalf("no error expected when creating a new scanner, got %s", err.Error())
	}

	for i := 0; i < 30; i++ {
		go func() {
			_, err := scanner.Scan()
			if err != nil {
				t.Errorf("no error expected when scanning, got %s", err.Error())
			}
		}()
	}

	time.Sleep(500 * time.Millisecond)
}

func TestARPScannerWithNoopExecutor(t *testing.T) {
	scanner, err := warp.NewARPScanner(
		warp.NoopExecutorOption,
	)
	if err != nil {
		t.Fatalf("no error expected when creating a new scanner, got %s", err.Error())
	}

	entries, err := scanner.Scan()
	if err != nil {
		t.Errorf("no error expected when scanning, got %s", err.Error())
	}

	if len(entries) != 0 {
		t.Errorf("no entries expected with noop scanner, got %d", len(entries))
	}
}

func TestARPScannerWithCustomExecutor(t *testing.T) {
	scanner, err := warp.NewARPScanner(
		warp.CustomCommandExecutorOption("echo"),
	)
	if err != nil {
		t.Fatalf("no error expected when creating a new scanner, got %s", err.Error())
	}

	entries, err := scanner.Scan()
	if err != nil {
		t.Errorf("no error expected when scanning, got %s", err.Error())
	}

	if len(entries) != 0 {
		t.Errorf("no entries expected with custom executor and invlaid output, got %d", len(entries))
	}
}
