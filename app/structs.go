package app

import "github.com/shirou/gopsutil/cpu"

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
