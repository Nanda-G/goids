package fetch

import (
	"testing"
)

func TestDiskInfo(t *testing.T) {
	got := DiskInfo()

	if got != nil {
		t.Error("DiskInfo failed to run: ", got)
	}
}

func BenchmarkDiskInfo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DiskInfo()
	}
}
