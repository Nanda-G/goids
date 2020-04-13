package foo

import (
	"testing"
)

func TestLogProcessInfo(t *testing.T) {

	t.Run("Testing JSON encoding", func(t *testing.T) {
		// Argument 1 for JSON
		got := LogProcessInfo(1)

		if got != nil {
			t.Error("LogProcess using JSON failed: ", got)
		}
	})

	t.Run("Testing gob encoding", func(t *testing.T) {
		// Argument 0 for Gob
		got := LogProcessInfo(0)

		if got != nil {
			t.Error("LogProcess using gob failed: ", got)
		}
	})

}

func BenchmarkLogProcessInfo(b *testing.B) {

	b.Run("Benchmark log to JSON", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			LogProcessInfo(1)
		}
	})

	b.Run("Benchmark log to gob", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			LogProcessInfo(0)
		}
	})
}
