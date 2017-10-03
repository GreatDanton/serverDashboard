// Creates path for websocket
function getSocketURL(path) {
    if (path[0] == "/") { // avoid url//path problem
        path = path.slice(1);
    }
    var url = "" + window.location;
    if (url.indexOf("https") > -1) {
        start = url.indexOf(":")
        socketURL = "wss" + url.slice(start) + path;
        return socketURL
    } else { // http protocol
        start = url.indexOf(":")
        socketURL = "ws" + url.slice(start) + path;
        return socketURL
    }
}

// stats establish websocket connection on /stats with backend application
// and displays data on frontend
function stats() {
    var wsurl = getSocketURL("/status")
    var socket = new WebSocket(wsurl)
    socket.onopen = function () {
        console.log("Connected to /status")
    }

    socket.onmessage = function (e) {
        var data = JSON.parse(e.data)
        var memory = data.Memory
        var cpu = data.CPU

        if (MemoryChartData.length >= 500) {
            MemoryChartData.splice(0, 1)
        }
        var newData = { x: memory.Time, y: memory.TakenPerc };
        MemoryChartData.push(newData);
        memoryChart.update();

        // remove first element when cpu data arr is bigger than 1000
        if (CpuChartData.length >= 500) {
            CpuChartData.splice(0, 1)
        }
        // append new data to cpu data arr and update chart
        var newCpuData = { x: cpu.Time, y: cpu.AverageLoad }
        CpuChartData.push(newCpuData)
        cpuChart.update();
    }

    socket.onclose = function () {
        console.log("Socket closed");
    }
}

// receive data from backend and display it on chart
stats();

