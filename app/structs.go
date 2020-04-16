package app

// ProcessLog is used for encoding process info to JSON
type ProcessLog struct {
	Name          string  `json:"name"`
	PID           int32   `json:"pid"`
	Background    bool    `json:"background"`
	CPUPercent    float64 `json:"cpupercent"`
	RunningTime   int64   `json:"runningtime"`
	MemoryPercent float32 `json:"memorypercent"`
	Status        string  `json:"status"`
	CPUAverages   CPUAverages
	MemAverages   MemAverages
}

type CPUAverages struct {
	Latest   float64
	LastDay  float64
	LastWeek float64
	AllTime  float64
}

type MemAverages struct {
	Latest   float32
	LastDay  float32
	LastWeek float32
	AllTime  float32
}
