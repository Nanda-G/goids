package foo

import (
	"math/rand"
	"testing"
	"time"

	"github.com/obviyus/goids/app"
)

func TestAvergeFunction(t *testing.T) {
	test := app.ProcessLog{
		PID:           12345,
		Name:          "service.exe",
		Background:    true,
		CPUPercent:    12.0,
		RunningTime:   1586640172643,
		MemoryPercent: 0.041369982,
		Status:        "",
	}

	averaged := AverageLog(test, test)
	expected := test

	if averaged != expected {
		t.Errorf("expected %v but got %v", expected, averaged)
	}
}

func BenchmarkAverageFunction(b *testing.B) {

	rand.Seed(time.Now().UnixNano())
	min := 0.0
	max := 100.0

	test := app.ProcessLog{
		PID:           12345,
		Name:          "service.exe",
		Background:    true,
		CPUPercent:    (min + rand.Float64()*(max-min)),
		RunningTime:   1586640172643,
		MemoryPercent: (float32(min) + rand.Float32()*(float32(max-min))),
		Status:        "",
	}

	for i := 0; i < b.N; i++ {
		AverageLog(test, test)
	}
}
