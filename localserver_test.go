package main

import (
	"testing"
)

func TestFail(t *testing.T) {
	if 2-2 < 0 {
		t.Error("this is a failed test")
	}
}
