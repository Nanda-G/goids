package main

import (
	"fmt"
	"time"

	"github.com/schollz/progressbar"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/winservices"
)

func main() {
	v, _ := mem.VirtualMemory()
	services, _ := winservices.ListServices()
	bar := progressbar.New(100)

	for {
		fmt.Printf("Services: %v", services)
		fmt.Printf(" Total: %v, Free:%v, UsedPercent:%v%%\n", v.Total, v.Free, int64(v.UsedPercent))
		bar.Set(int(v.UsedPercent))

		time.Sleep(2 * time.Second)
	}
}
