package foo

import (
	"github.com/obviyus/goids/app"
)

// AverageLog takes input 2 ProcessLog structs and returns avergage
// of their CPU and Memory percentages
func AverageLog(prevLog app.ProcessLog,
	newLog app.ProcessLog) app.ProcessLog {

	averageLog := app.ProcessLog{
		PID:           newLog.PID,
		Name:          newLog.Name,
		Background:    newLog.Background,
		CPUPercent:    ((newLog.CPUPercent + prevLog.CPUPercent) / 2.0),
		RunningTime:   newLog.RunningTime,
		MemoryPercent: ((newLog.MemoryPercent + prevLog.MemoryPercent) / 2.0),
		Status:        newLog.Status,
	}

	return (averageLog)
}
