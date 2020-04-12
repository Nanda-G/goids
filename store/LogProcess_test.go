package store

import (
	"testing"
)

func TestLogProcessInfo(t *testing.T) {
	got := LogProcessInfo()

	if got != nil {
		t.Error("LogProcessInfo failed to run: ", got)
	}
}

func BenchmarkLogProcessInfo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LogProcessInfo()
	}
}
