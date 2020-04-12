package store

import (
	"os"
	"testing"
)

func TestDiskInfo(t *testing.T) {
	got := LogProcessInfo()

	if got != nil {
		t.Error("DiskInfo failed to run: ", got)
	}
}

func BenchmarkDiskInfo(b *testing.B) {
	// Discard output of DiskInfo during benchmark
	os.Stdout, _ = os.Open(os.DevNull)

	for i := 0; i < b.N; i++ {
		LogProcessInfo()
	}
}
