package foo

import (
	"testing"
)

func TestLogProcessInfo(t *testing.T) {
	got := LogProcessInfo()

	if got != nil {
		t.Error("DiskInfo failed to run: ", got)
	}
}

func BenchmarkLogProcessInfo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LogProcessInfo()
	}
}
