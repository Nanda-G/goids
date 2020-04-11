package getinfo

import (
	"fmt"
	"log"

	"github.com/shirou/gopsutil/process"
)

// ProcessInfo lists information about processes running on host machine
func ProcessInfo() {
	processes, err := process.Processes()
	if err != nil {
		log.Fatal(err)
	}

	for _, proc := range processes {
		pid := proc.Pid
		// (time.Sleep(time.Second * 1))

		name, _ := proc.Name()
		cpuPercent, _ := proc.CPUPercent()

		fmt.Printf("PID: %v \tName: %v \tCPU Usage: %v%%\n", pid, name, int(cpuPercent))
	}
}
