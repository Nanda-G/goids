package foo

import (
	"encoding/gob"
	"log"
	"os"
	"strings"

	"github.com/obviyus/goids/app"
	"github.com/shirou/gopsutil/process"
)

// LogProcessInfo logs all currently running processes with:
// Name, PID, background, CPU Percent, Running Time,
// Memory Percent, Status and CPU times
func LogProcessInfo(choice int) error {

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

		pL := app.ProcessLog{
			PID:           proc.Pid,
			Name:          name,
			Background:    background,
			CPUPercent:    cpuPercent,
			RunningTime:   runningTime,
			MemoryPercent: memoryPercent,
			Status:        status,
		}

		var filename strings.Builder
		filename.WriteString("/logs/")
		filename.WriteString(pL.Name)

		filename.WriteString(".gob")
		err := writeToGob(pL, filename.String())

		if err != nil {
			log.Fatal("Error encoding to gob: ", err)
		}

	}
	return nil
}

func writeToGob(pL app.ProcessLog, loc string) error {

	if pL.Name != "" {
		fi, err := os.OpenFile(loc, os.O_CREATE|os.O_WRONLY, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
		defer fi.Close()

		enc := gob.NewEncoder(fi)
		err = enc.Encode(pL)
		if err != nil {
			return err
		}
	}
	return nil
}
