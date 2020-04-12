package foo

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/obviyus/goids/app"
	"github.com/shirou/gopsutil/process"
)

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

		pL := app.ProcessLog{
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

func writeToJSON(pL app.ProcessLog) {

	var filename strings.Builder
	filename.WriteString("./logs/")
	filename.WriteString(pL.Name)
	filename.WriteString(".json")

	file, _ := os.OpenFile(filename.String(), os.O_CREATE, os.ModePerm)
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.Encode(pL)
}
