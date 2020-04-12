package fetch

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/dustin/go-humanize"
	"github.com/shirou/gopsutil/disk"
)

// DiskInfo reads partitions on disk and returns Total, Free and Used space
func DiskInfo() error {
	parts, err := disk.Partitions(false)
	if err != nil {
		return err
	}

	var usage []*disk.UsageStat

	writer := new(tabwriter.Writer)

	// Format in tab-separated columns with a tab stop of 8.
	writer.Init(os.Stdout, 0, 8, 8, '\t', tabwriter.AlignRight)

	fmt.Fprintf(writer, "Drive\t%% Used\tTotal\tUsed\tFree\n")

	for _, part := range parts {
		u, err := disk.Usage(part.Mountpoint)
		if err != nil {
			return err
		}
		usage = append(usage, u)
		printUsage(u, writer)
	}
	return nil
}

func printUsage(u *disk.UsageStat, writer *tabwriter.Writer) {
	fmt.Fprintf(writer, "%v\t%v %%\t%v\t%v\t%v\n", u.Path, int(u.UsedPercent),
		humanize.Bytes(u.Total), humanize.Bytes(u.Free),
		humanize.Bytes(u.Used))
	writer.Flush()
}
