package foo

import (
	"math/rand"
	"testing"
	"time"

	"github.com/obviyus/goids/app"
)

func TestAvergeFunction(t *testing.T) {

	rand.Seed(time.Now().UnixNano())
	min := 0.0
	max := 100.0

	cpuX := (min + rand.Float64()*(max-min))
	memX := (float32(min) + rand.Float32()*(float32(max-min)))
	cpuY := (min + rand.Float64()*(max-min))
	memY := (float32(min) + rand.Float32()*(float32(max-min)))

	x := app.ProcessLog{
		PID:           12345,
		Name:          "serviceX.exe",
		Background:    true,
		CPUPercent:    cpuX,
		RunningTime:   1586640172643,
		MemoryPercent: memX,
		Status:        "",
	}

	y := app.ProcessLog{
		PID:           54321,
		Name:          "serviceY.exe",
		Background:    true,
		CPUPercent:    cpuY,
		RunningTime:   1586648592643,
		MemoryPercent: memY,
		Status:        "",
	}

	cpuAverage := (cpuX + cpuY) / 2.0
	memAverage := (memX + memY) / 2.0

	expected := app.ProcessLog{
		PID:           54321,
		Name:          "serviceY.exe",
		Background:    true,
		CPUPercent:    cpuAverage,
		RunningTime:   1586648592643,
		MemoryPercent: memAverage,
		Status:        "",
	}

	averaged := AverageLog(x, y)

	if averaged != expected {
		t.Errorf("expected %v but got %v", expected, averaged)
	}
}

func BenchmarkAverageFunction(b *testing.B) {

	rand.Seed(time.Now().UnixNano())
	min := 0.0
	max := 100.0

	x := app.ProcessLog{
		PID:           12345,
		Name:          "serviceX.exe",
		Background:    true,
		CPUPercent:    (min + rand.Float64()*(max-min)),
		RunningTime:   1586640172643,
		MemoryPercent: (float32(min) + rand.Float32()*(float32(max-min))),
		Status:        "",
	}

	y := app.ProcessLog{
		PID:           54321,
		Name:          "serviceY.exe",
		Background:    true,
		CPUPercent:    (min + rand.Float64()*(max-min)),
		RunningTime:   1586648592643,
		MemoryPercent: (float32(min) + rand.Float32()*(float32(max-min))),
		Status:        "",
	}

	for i := 0; i < b.N; i++ {
		AverageLog(x, y)
	}
}
