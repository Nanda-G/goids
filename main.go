package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/obviyus/goids/getinfo"
	"github.com/obviyus/goids/loginfo"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
)

func main() {
	cpu, _ := cpu.Info()
	fmt.Printf("Cores: %v \tCPU Model: %v \n", cpu[0].Cores, cpu[0].ModelName)

	// List partitions and total, used and free space
	getinfo.DiskInfo()

	infoStat, _ := host.Info()
	fmt.Printf("Total processes: %d\n", infoStat.Procs)

	// List processes with PID, name and CPU usage
	getinfo.ProcessInfo()

	// TODO: Logging all ProcessInfo outputs
	loginfo.LogProcessInfo()

	fmt.Print("Press 'Enter' to exit...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
