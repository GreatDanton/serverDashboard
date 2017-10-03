package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/greatdanton/serverDashboard/stats"

	"golang.org/x/net/websocket"
)

func main() {
	fmt.Println("Starting server: http://127.0.0.1:8080")

	http.HandleFunc("/", displayDashboard)
	http.Handle("/status", websocket.Handler(status))
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe", err)
	}
}

type displayData struct {
	Memory stats.ChartMemory
	CPU    stats.ChartCPU
}

// status is streaming data to client/frontend via websockets
func status(ws *websocket.Conn) {
	for {
		// get data about memory consumption
		memory, err := stats.UsedMemory()
		if err != nil {
			fmt.Println(err)
			break
		}
		// get cpu load data
		cpu1, err := stats.UsedCPU()
		if err != nil {
			fmt.Println(err)
			break
		}
		// waiting for one second so we can calculate the
		// the cpu load
		time.Sleep(time.Second * 1)
		// get cpu load data after one second
		cpu2, err := stats.UsedCPU()
		if err != nil {
			fmt.Println(err)
			break
		}
		// calculate average cpu load
		idleTicks := float32(cpu2.Idle - cpu1.Idle)
		totalTicks := float32(cpu2.Total - cpu1.Total)
		averageLoad := fmt.Sprintf("%.1f", (totalTicks-idleTicks)*100/totalTicks)
		cpu := stats.ChartCPU{AverageLoad: averageLoad, Time: cpu2.Time}

		// gather memory and cpu data together, and send it
		// as JSON to client via websocket
		data := displayData{Memory: memory, CPU: cpu}
		err = websocket.JSON.Send(ws, data)
		if err != nil {
			fmt.Println("Stats: Can't push to client:", err)
			break
		}
	}
}

var t = template.Must(template.ParseFiles("./index.html"))

// displayDashboard is displaying index page of server dashboard
func displayDashboard(w http.ResponseWriter, r *http.Request) {
	err := t.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
