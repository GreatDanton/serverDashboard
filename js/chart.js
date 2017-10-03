var ctx = document.getElementById('memory-chart');

// MemoryChartData is used to display display points in
// memory chart
var MemoryChartData = [];


var memoryChart = new Chart(ctx, {
    type: 'line',
    data: {
        datasets: [{
            label: 'MemoryConsumption',
            backgroundColor: 'rgba(54, 162, 235, 0.2)',
            borderColor: 'rgb(54, 162, 235)',
            pointBorderWidth: 0,
            pointRadius: 0,
            hitRadius: 10,
            hoverRadius: 10,
            pointHoverBackgroundColor: 'rgb(54, 162, 235)',
            data: MemoryChartData
        }]
    },
    options: {
        animation: false,
        scales: {
            xAxes: [{
                type: 'time',
                unit: 'minute',
                time: {
                    displayFormats: {
                        minute: "h:mm",
                        second: "h:mm",
                        milisecond: "h:mm",
                        hour: "h:mm",
                        day: "h:mm",
                        week: "h:mm",
                        month: "h:mm",
                        quarter: "h:mm",
                        year: "h:mm"
                    },
                    tooltipFormat: 'h:mm:ss',
                },
                position: 'bottom',
            }],
            yAxes: [{
                display: true,
                ticks: {
                    beginAtZero: true,
                    max: 100
                }
            }]
        }
    }
});


var cpu = document.getElementById('cpu-chart');
var CpuChartData = [];

var cpuChart = new Chart(cpu, {
    type: 'line',
    data: {
        datasets: [{
            label: 'Cpu consumption',
            backgroundColor: 'rgba(255, 99, 132, 0.2)',
            borderColor: 'rgb(255, 99, 132)',
            pointBorderWidth: 0,
            pointRadius: 0,
            hitRadius: 10,
            hoverRadius: 10,
            pointHoverBackgroundColor: 'rgb(255, 99, 132)',
            data: CpuChartData
        }]
    },
    options: {
        animation: false,
        scales: {
            xAxes: [{
                type: 'time',
                unit: 'minute',
                time: {
                    displayFormats: {
                        minute: "h:mm",
                        second: "h:mm",
                        milisecond: "h:mm",
                        hour: "h:mm",
                        day: "h:mm",
                        week: "h:mm",
                        month: "h:mm",
                        quarter: "h:mm",
                        year: "h:mm"
                    },
                    tooltipFormat: 'h:mm:ss',
                },
                position: 'bottom',
            }],
            yAxes: [{
                display: true,
                ticks: {
                    beginAtZero: true,
                    max: 100
                }
            }]
        }
    }
});


