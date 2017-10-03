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
	http.Handle("/echo", websocket.Handler(echo))
	http.Handle("/status", websocket.Handler(status))
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe", err)
	}
}

func echo(ws *websocket.Conn) {
	var err error
	for {
		var reply string
		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive message")
			break
		}

		fmt.Println("Received from client:", reply)
		msg := reply
		fmt.Println("Sending to client:" + reply)
		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Can't reply to client")
			break
		}

	}
}

type displayData struct {
	Memory stats.ChartMemory
	CPU    stats.ChartCPU
}

func status(ws *websocket.Conn) {
	for {
		memory, err := stats.UsedMemory()
		if err != nil {
			fmt.Println(err)
			break
		}
		cpu1, err := stats.UsedCPU()
		if err != nil {
			fmt.Println(err)
			break
		}
		time.Sleep(time.Second * 1)

		cpu2, err := stats.UsedCPU()
		if err != nil {
			fmt.Println(err)
			break
		}
		idleTicks := float32(cpu2.Idle - cpu1.Idle)
		totalTicks := float32(cpu2.Total - cpu1.Total)
		averageLoad := fmt.Sprintf("%.1f", (totalTicks-idleTicks)*100/totalTicks)
		cpu := stats.ChartCPU{AverageLoad: averageLoad, Time: cpu2.Time}

		// load data and send it via websockets
		data := displayData{Memory: memory, CPU: cpu}

		err = websocket.JSON.Send(ws, data)
		if err != nil {
			fmt.Println("Stats: Can't push to client:", err)
			break
		}
	}
}

func displayDashboard(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./index.html"))
	err := t.Execute(w, t)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
