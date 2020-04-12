package store

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/process"
)

type processLog struct {
	Name          string         `json:"name"`
	PID           int32          `json:"pid"`
	Background    bool           `json:"background"`
	CPUPercent    float64        `json:"cpupercent"`
	RunningTime   int64          `json:"runningtime"`
	MemoryPercent float32        `json:"memorypercent"`
	Status        string         `json:"status"`
	CPUTimes      *cpu.TimesStat `json:"cputimes"`
}

// LogProcessInfo logs all currently running processes with:
// Name, PID, background, CPU Percent, Running Time,
// Memory Percent, Status and CPU times
func LogProcessInfo() error {

	_, err := os.Stat("logs")

	// If logs folder does not exist, create it
	if os.IsNotExist(err) {
		errDir := os.MkdirAll("logs", 0755)
		if errDir != nil {
			log.Fatal(err)
		}
	}

	processes, err := process.Processes()
	if err != nil {
		log.Fatal(err)
	}

	for _, proc := range processes {
		name, _ := proc.Name()
		background, _ := proc.Background()
		cpuPercent, _ := proc.CPUPercent()
		runningTime, _ := proc.CreateTime()
		memoryPercent, _ := proc.MemoryPercent()
		status, _ := proc.Status()
		cpuTimes, _ := proc.Times()

		pL := processLog{
			PID:           proc.Pid,
			Name:          name,
			Background:    background,
			CPUPercent:    cpuPercent,
			RunningTime:   runningTime,
			MemoryPercent: memoryPercent,
			Status:        status,
			CPUTimes:      cpuTimes,
		}

		go writeToJSON(pL)
		if err != nil {
			return err
		}
	}
	return nil
}

func writeToJSON(pL processLog) {

	var filename strings.Builder
	filename.WriteString("./logs/")
	filename.WriteString(pL.Name)
	filename.WriteString(".json")

	file, _ := os.OpenFile(filename.String(), os.O_CREATE, os.ModePerm)
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.Encode(pL)
}
