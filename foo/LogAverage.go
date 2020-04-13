package foo

import (
	"github.com/obviyus/goids/app"
)

// AverageLog takes input 2 ProcessLog structs and returns avergage
// of their CPU and Memory percentages
func AverageLog(prevLog app.ProcessLog,
	newLog app.ProcessLog) app.ProcessLog {

	averageLog := newLog

	averageLog.CPUAverages.Latest = (newLog.CPUPercent + prevLog.CPUPercent) / 2.0
	averageLog.MemAverages.Latest = (newLog.MemoryPercent + prevLog.MemoryPercent) / 2.0

	return (averageLog)
}
