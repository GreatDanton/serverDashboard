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

// stats establish websocket connection on /stats with our server
function stats() {
    var wsurl = getSocketURL("/status")
    var socket = new WebSocket(wsurl)
    socket.onopen = function () {
        console.log("Connected to /status")
    }

    socket.onmessage = function (e) {
        data = JSON.parse(e.data)
        console.log(data.TakenPerc);
        console.log(data.Time);

        if (MemoryChartData.length >= 1000) {
            MemoryChartData.splice(0, 1)
        }
        var newData = { x: data.Time, y: data.TakenPerc };
        MemoryChartData.push(newData);
        memoryChart.update();
    }

    socket.onclose = function () {
        console.log("Socket closed");
    }
}

function echo() {
    var socket = null;
    var wsurl = "ws://127.0.0.1:8080/echo"
    console.log("On load")

    socket = new WebSocket(wsurl)

    socket.onopen = function () {
        console.log("connected to " + wsurl);
    }
    socket.onclose = function (e) {
        console.log("connection closed" + e.code);
    }
    socket.onmessage = function (e) {
        console.log("Received: " + e.data)
    }

    var msgInput = document.getElementById('message');
    function send() {
        var msg = msgInput.value;

        socket.send(msg)
    }
    var btn = document.getElementById('btn-send');
    btn.addEventListener('click', function () {
        send();
    });
}

stats();

