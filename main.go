package main

import (
	"fmt"
	"time"

	"github.com/schollz/progressbar"
	"github.com/shirou/gopsutil/mem"
)

func main() {
	v, _ := mem.VirtualMemory()

	for {
		bar := progressbar.New(100)
		bar.Add(int(v.UsedPercent))
		fmt.Printf(" Total: %v, Free:%v, UsedPercent:%v%%\n", v.Total, v.Free, int64(v.UsedPercent))
		time.Sleep(2 * time.Second)
	}
}
