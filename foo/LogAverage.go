package foo

import (
	"encoding/gob"
	"os"

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

// ReadLog checks if pased ProcessLog exists in stored logs and returns
// ProcessLog with last saved value
func ReadLog(check app.ProcessLog, loc string) (app.ProcessLog, error) {

	openedLog := new(app.ProcessLog)

	fo, err := os.Open(loc)
	if err != nil {
		return *openedLog, err
	}

	dec := gob.NewDecoder(fo)
	err = dec.Decode(&openedLog)

	return *openedLog, nil
}
