package tests

import (
	"testing"

	"gangleri.io/pkg/humanbytes"
)

// Test importing and using the core package from another package

func TestParse(t *testing.T) {
	b, e := humanbytes.ParseBytes("100 MB")
	if e != nil || b != 104857600 {
		t.Error("100 MB should be 1024 Bytes")
	}
}
