package foo

import (
	"encoding/gob"
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
		filename.WriteString("./logs/")
		filename.WriteString(pL.Name)

		if choice == 1 {
			filename.WriteString(".json")
			writeToJSON(pL, filename.String())
		} else {
			filename.WriteString(".gob")
			err := writeToGob(pL, filename.String())

			if err != nil {
				log.Fatal("Error encoding to gob")
			}
		}

	}
	return nil
}

func writeToJSON(pL app.ProcessLog, ext string) {

	if pL.Name != "" {
		file, _ := os.OpenFile(ext, os.O_CREATE, os.ModePerm)
		defer file.Close()
		enc := json.NewEncoder(file)
		err := enc.Encode(pL)
		if err != nil {
			log.Fatal("encode error:", err)
		}
	}
}

func writeToGob(pL app.ProcessLog, ext string) error {

	if pL.Name != "" {
		fi, _ := os.OpenFile(ext, os.O_CREATE, os.ModePerm)
		defer fi.Close()

		enc := gob.NewEncoder(fi)
		err := enc.Encode(pL)
		if err != nil {
			return err
		}
	}
	return nil
}
