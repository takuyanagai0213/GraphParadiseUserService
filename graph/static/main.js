var rooms = [];
class chart {
  constructor(){
    this.rooms = []
  }
  createGraph() {
    const self = this;
    $.ajax({
      url: '/getData1',
      type: "get",
      dataType: 'json',
    }).then(function (json) {
      $.ajax({
        url: '/getRooms',
        type: "get",
        dataType: 'json',
      }).then(function (rooms) {
        self.drowChart(json, rooms, "myChart1");
        self.drowLineChart(json, rooms, "myChart2");
        self.drowPieChart(json, rooms, "myChart3");
        self.drowRadarChart(json, rooms, "myChart4");
      });
    });
    $.ajax({
      url: '/getDataForDaily',
      type: "get",
      dataType: 'json',
    }).then(function (json) {
      $.ajax({
        url: '/getDates',
        type: "get",
        dataType: 'json',
      }).then(function (dates) {
        self.drowLineChart(json, dates, "myChart2");
      });
    });
  }
  drowChart(json, rooms, target) {
    console.log(json)
    var ctx = document.getElementById(target);
    var myChart = new Chart(ctx, {
      type: 'bar',
      data: {
      labels: rooms,
      datasets: [
        {
          label: '最高気温(度）',
          data: json,
          // borderColor: "rgba(255,0,0,1)",
          backgroundColor: "rgb(0, 255, 0)"
        },
        // {
        //   label: '最高気温(度）',
        //   data: json,
        //   borderColor: "rgba(255,0,0,1)",
        //   backgroundColor: "rgba(0,0,0,0)"
        // },
      ],
    },
    options: {
      title: {
        display: true,
        text: '気温（8月1日~8月7日）'
      },
      scales: {
        yAxes: [{
          ticks: {
            suggestedMax: 40,
            suggestedMin: 0,
            stepSize: 10,
            callback: function(value, index, values){
              return  value +  '度'
            }
          }
        }]
      }
    }
  });
  }
  drowLineChart(json, dates, target) {
    console.log(json)
    var ctx = document.getElementById(target);
    var myChart = new Chart(ctx, {
      type: 'line',
      data: {
      labels: dates,
      datasets: [
        {
          label: '最高気温(度）',
          data: json,
          // borderColor: "rgba(255,0,0,1)",
          backgroundColor: "rgb(0, 255, 0)"
        },
        // {
        //   label: '最高気温(度）',
        //   data: json,
        //   borderColor: "rgba(255,0,0,1)",
        //   backgroundColor: "rgba(0,0,0,0)"
        // },
      ],
    },
    options: {
      title: {
        display: true,
        text: '気温（8月1日~8月7日）'
      },
      scales: {
        yAxes: [{
          ticks: {
            suggestedMax: 40,
            suggestedMin: 0,
            stepSize: 10,
            callback: function(value, index, values){
              return  value +  '度'
            }
          }
        }]
      }
    }
  });
  }
  drowPieChart(json, rooms, target) {
    console.log(json)
    var ctx = document.getElementById(target);
    var myChart = new Chart(ctx, {
      type: 'pie',
      data: {
      labels: rooms,
      datasets: [
        {
          label: '最高気温(度）',
          data: json,
          // borderColor: "rgba(255,0,0,1)",
          backgroundColor: "rgb(0, 255, 0)"
        },
        // {
        //   label: '最高気温(度）',
        //   data: json,
        //   borderColor: "rgba(255,0,0,1)",
        //   backgroundColor: "rgba(0,0,0,0)"
        // },
      ],
    },
    options: {
      title: {
        display: true,
        text: '気温（8月1日~8月7日）'
      },
      scales: {
        yAxes: [{
          ticks: {
            suggestedMax: 40,
            suggestedMin: 0,
            stepSize: 10,
            callback: function(value, index, values){
              return  value +  '度'
            }
          }
        }]
      }
    }
  });
  }
  drowRadarChart(json, rooms, target) {
    console.log(json)
    var ctx = document.getElementById(target);
    var myChart = new Chart(ctx, {
      type: 'radar',
      data: {
      labels: rooms,
      datasets: [
        {
          label: '最高気温(度）',
          data: json,
          // borderColor: "rgba(255,0,0,1)",
          backgroundColor: "rgb(0, 255, 0)"
        },
        // {
        //   label: '最高気温(度）',
        //   data: json,
        //   borderColor: "rgba(255,0,0,1)",
        //   backgroundColor: "rgba(0,0,0,0)"
        // },
      ],
    },
    options: {
      title: {
        display: true,
        text: '気温（8月1日~8月7日）'
      },
      scales: {
        yAxes: [{
          ticks: {
            suggestedMax: 40,
            suggestedMin: 0,
            stepSize: 10,
            callback: function(value, index, values){
              return  value +  '度'
            }
          }
        }]
      }
    }
  });
  }
}
const chartClass = new chart();

chartClass.createGraph()