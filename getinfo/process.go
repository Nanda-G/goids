package getinfo

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/shirou/gopsutil/process"
)

// ProcessInfo lists information about processes running on host machine
func ProcessInfo() error {
	processes, err := process.Processes()
	if err != nil {
		return err
	}

	writer := new(tabwriter.Writer)

	// Format in tab-separated columns with a tab stop of 8.
	writer.Init(os.Stdout, 0, 8, 8, '\t', tabwriter.AlignRight)
	fmt.Fprintf(writer, "PID\tName\tCPU Usage\n")

	for _, proc := range processes {
		pid := proc.Pid

		name, _ := proc.Name()
		cpuPercent, _ := proc.CPUPercent()

		if cpuPercent > 1 {
			fmt.Fprintf(writer, "%v\t%v\t%v%%\n", pid, name, int(cpuPercent))
			writer.Flush()
		}
	}
	return nil
}
