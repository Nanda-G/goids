package main

import (
	"fmt"

	"goids/getinfo"

	"github.com/shirou/gopsutil/cpu"
)

func main() {
	cpu, _ := cpu.Info()
	// fmt.Printf("Raw CPU data: %v", c)
	fmt.Printf("Cores: %v \tCPU Model: %v \n", cpu[0].Cores, cpu[0].ModelName)

	getinfo.DiskInfo()

	//TODO: List all processes
	// p, _ := process.Processes()
	fmt.Printf("\nList of Processes Running: \n")

}
