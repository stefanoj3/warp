package warp_test

import (
	"errors"
	"strings"
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

func TestCustomExecutorOptionShouldReturnAnErrorWhenFunctionIsNil(t *testing.T) {
	_, err := warp.NewARPScanner(
		warp.CustomExecutorOption(nil),
	)
	if err == nil {
		t.Fatal("an error is expected when nil is passed instead of a function to the custom executor")
	}

	if !strings.Contains(err.Error(), "CustomExecutorOption") {
		t.Fatal("the error is expected to come from CustomExecutorOption")
	}
}

func TestScanShouldReturnAnErrorWhenExecutorFails(t *testing.T) {
	customError := errors.New("my_custom_error")

	scanner, err := warp.NewARPScanner(
		warp.CustomExecutorOption(
			func() (bytes []byte, e error) {
				return nil, customError
			},
		),
	)
	if err != nil {
		t.Fatalf("no error expected when creating a new scanner, got %s", err.Error())
	}

	entries, err := scanner.Scan()
	if err == nil {
		t.Fatal("the error is expected to come from CustomExecutorOption")
	}

	if !strings.Contains(err.Error(), customError.Error()) {
		t.Fatal("the error is expected to come from CustomExecutorOption")
	}

	if entries != nil {
		t.Fatal("no entries are expected when an error is produced")
	}
}
