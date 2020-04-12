package tests

import (
	"testing"

	"github.com/obviyus/goids/getinfo"
)

func TestDiskInfo(t *testing.T) {
	got := getinfo.DiskInfo()

	if got != nil {
		t.Error("DiskInfo failed to run: ", got)
	}
}
