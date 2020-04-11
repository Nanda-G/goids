package getinfo

import (
	"fmt"
	"strconv"

	"github.com/shirou/gopsutil/disk"
)

// DiskInfo reads partitions on disk and returns Total, Free and Used space
func DiskInfo() {
	parts, err := disk.Partitions(false)
	check(err)

	var usage []*disk.UsageStat

	for _, part := range parts {
		u, err := disk.Usage(part.Mountpoint)
		check(err)
		usage = append(usage, u)
		printUsage(u)
	}
}

func printUsage(u *disk.UsageStat) {
	fmt.Printf(u.Path + "\t" + strconv.FormatFloat(u.UsedPercent, 'f', 2, 64) + "%% full \t")
	fmt.Printf("Total \t: " + strconv.FormatUint(u.Total/1024/1024/1024, 10) + " GiB \t")
	fmt.Printf("Free \t:  " + strconv.FormatUint(u.Free/1024/1024/1024, 10) + " GiB \t")
	fmt.Printf("Used \t:  " + strconv.FormatUint(u.Used/1024/1024/1024, 10) + " GiB \n")
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
