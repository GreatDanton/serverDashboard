package stats

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Memory type holding data of computer memory
type Memory struct {
	Total     int // Total amount of memory in kb
	Available int // Available memory in kb
	Taken     int // Taken memory in kb
	Time      int64
}

// ChartMemory is used to display x: time, y: percentage of memory used
type ChartMemory struct {
	TakenPerc string
	Time      int64
}

// UsedMemory returns amount of used memory in percentage
// or error if an error happens
func UsedMemory() (ChartMemory, error) {
	chartMemory := ChartMemory{}
	mem, err := getMemoryInfo()
	if err != nil {
		return chartMemory, err
	}
	percentage := fmt.Sprintf("%.1f", float32(mem.Taken)*100/float32(mem.Total))
	chartMemory.TakenPerc = percentage
	chartMemory.Time = mem.Time
	return chartMemory, nil
}

// getMemoryInfo parses memory info from /proc/meminfo
// and returns Memory type
func getMemoryInfo() (Memory, error) {
	mem := Memory{}
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		return mem, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "MemTotal") {
			// parse fields in line and pick number only
			total := strings.Fields(line)[1]
			mem.Total, err = strconv.Atoi(total)
			if err != nil {
				return mem, err
			}
		} else if strings.Contains(line, "MemAvailable") {
			avail := strings.Fields(line)[1]
			mem.Available, err = strconv.Atoi(avail)
			if err != nil {
				return mem, err
			}
		}
	}
	mem.Taken = mem.Total - mem.Available
	mem.Time = time.Now().Unix() * 1000
	return mem, nil
}
