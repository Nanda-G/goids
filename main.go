package main

import (
	"fmt"

	"github.com/dustin/go-humanize"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
)

func main() {
	cpu, _ := cpu.Info()
	disk, _ := disk.Usage("/")
	// fmt.Printf("Raw CPU data: %v", c)
	fmt.Printf("Cores: %v \tCPU Model: %v \n", cpu[0].Cores, cpu[0].ModelName)
	fmt.Printf("Total: %v \tFree: %v \t Used Percent: %v%%", humanize.Bytes(disk.Total), humanize.Bytes(disk.Free),
		int(disk.UsedPercent))

	//TODO: List all processes
	// p, _ := process.Processes()
	fmt.Printf("\nList of Processes Running: \n")

}
