package app

import "github.com/shirou/gopsutil/cpu"

// ProcessLog is used for encoding process info to JSON
type ProcessLog struct {
	Name          string         `json:"name"`
	PID           int32          `json:"pid"`
	Background    bool           `json:"background"`
	CPUPercent    float64        `json:"cpupercent"`
	RunningTime   int64          `json:"runningtime"`
	MemoryPercent float32        `json:"memorypercent"`
	Status        string         `json:"status"`
	CPUTimes      *cpu.TimesStat `json:"cputimes"`
}

// AverageLog is used for encoding average process info to JSON
type AverageLog struct {
	Name          string         `json:"name"`
	CPUPercent    float64        `json:"cpupercent"`
	MemoryPercent float32        `json:"memorypercent"`
	CPUTimes      *cpu.TimesStat `json:"cputimes"`
}
