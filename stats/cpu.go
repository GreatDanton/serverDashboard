package stats

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"time"
)

// CPU is used for holding cpu ticks data
type CPU struct {
	Idle  uint64
	Total uint64
	Time  int64
}

// ChartCPU is used for holding final chart data that
// is send to frontend of the application
type ChartCPU struct {
	AverageLoad string
	Time        int64
}

// UsedCPU returns average cpu load data, for displaying
// in chart
func UsedCPU() (CPU, error) {
	chartData := CPU{}
	file, err := os.Open("/proc/stat")
	if err != nil {
		return chartData, err
	}
	defer file.Close()
	var total uint64
	var idle uint64

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line) // split line into fields
		if fields[0] == "cpu" {
			for i := 1; i < len(fields); i++ {
				val, err := strconv.ParseUint(fields[i], 10, 64)
				if err != nil {
					return chartData, err
				}
				total += val // total number of ticks
				if i == 4 {  // idle is 5th fields in cpu line
					idle = val
				}
			}
			chartData.Time = time.Now().Unix() * 1000
			chartData.Total = total
			chartData.Idle = idle
			return chartData, nil
		}
	}
	return chartData, nil
}
