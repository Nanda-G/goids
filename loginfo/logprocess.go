package loginfo

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/process"
)

type processLog struct {
	Name          string
	PID           int32
	Background    bool
	CPUPercent    float64
	RunningTime   int64
	MemoryPercent float32
	Status        string
	CPUTimes      *cpu.TimesStat
}

// LogProcessInfo logs all currently running processes with:
// Name, PID, background, CPU Percent, Running Time,
// Memory Percent, Status and CPU times
func LogProcessInfo() {
	processes, err := process.Processes()
	if err != nil {
		log.Fatal(err)
	}

	for _, proc := range processes {
		pid := proc.Pid
		name, _ := proc.Name()
		background, _ := proc.Background()
		cpuPercent, _ := proc.CPUPercent()
		runningTime, _ := proc.CreateTime()
		memoryPercent, _ := proc.MemoryPercent()
		status, _ := proc.Status()
		cpuTimes, _ := proc.Times()

		pL := processLog{
			PID:           pid,
			Name:          name,
			Background:    background,
			CPUPercent:    cpuPercent,
			RunningTime:   runningTime,
			MemoryPercent: memoryPercent,
			Status:        status,
			CPUTimes:      cpuTimes,
		}

		jsonPL, _ := json.MarshalIndent(pL, "", "    ")
		var filename strings.Builder
		filename.WriteString(pL.Name)
		filename.WriteString(".json")
		err = ioutil.WriteFile(filename.String(), jsonPL, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
}
