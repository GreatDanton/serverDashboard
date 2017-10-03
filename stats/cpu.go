package stats

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"time"
)

// CPU is used for parsing cpu ticks data
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

// UsedCPU returns total number of cpu ticks and number of idle ticks
// To get an average cpu load we have to call UsedCPU with time difference
// of at least one second. By calculating the difference between cpu ticks
// we get the average cpu load.
func UsedCPU() (CPU, error) {
	chartData := CPU{}
	file, err := os.Open("/proc/stat")
	if err != nil {
		return chartData, err
	}
	defer file.Close()
	var total uint64
	var idle uint64
	// for each line of file we check if cpu exists in first field of line
	// if it does we gather the total number of ticks and number of idle ticks
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
