console.log('www')
class chart {
  constructor(json){
    console.log(this.json)
  }
  drowChart(json, target){
    console.log(json)
    var ctx = document.getElementById(target);
    var myChart = new Chart(ctx, {
      type: 'line',
      data: {
      labels: ['8月1日', '8月2日', '8月3日', '8月4日', '8月5日', '8月6日', '8月7日'],
      datasets: [
        {
          label: '最高気温(度）',
          data: json,
          borderColor: "rgba(255,0,0,1)",
          backgroundColor: "rgba(0,0,0,0)"
        },
        {
          label: '最高気温(度）',
          data: json,
          borderColor: "rgba(255,0,0,1)",
          backgroundColor: "rgba(0,0,0,0)"
        },
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
      },
    }
    });
  }
}
const chartClass = new chart();
$.ajax({
  url: '/getData',
  type: "get",
  dataType: 'json',
}).then(function (json) {
  console.log(json)
  chartClass.drowChart(json, "myChart1");
});
$.ajax({
  url: '/getData',
  type: "get",
  dataType: 'json',
}).then(function (json) {
  console.log(json)
  chartClass.drowChart(json, "myChart2");
});
$.ajax({
  url: '/getData',
  type: "get",
  dataType: 'json',
}).then(function (json) {
  console.log(json)
  chartClass.drowChart(json, "myChart3");
});