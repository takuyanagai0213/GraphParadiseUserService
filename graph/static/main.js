function Graph_Trigger() {
  this.view_s = $.extend(true, {}, this.def_s);
  return;
}
Graph_Trigger.prototype.def_s = new Object({
  rooms: {},
  room_name_list: []
});
Graph_Trigger.prototype.createGraph = function() {
  $.ajax({
    url: '/getRooms',
    type: "get",
    dataType: 'json',
  }).then(this.formatRoomData.bind(this));
  $.ajax({
    url: '/getAreas',
    type: "get",
    dataType: 'json',
  }).then(this.createAreaSelectBox.bind(this));
  this.getDataByArea();
}
Graph_Trigger.prototype.createAreaSelectBox = function(areas) {
  console.log(areas)
  for(let key in areas){
    $('#area_select').append('<option value=' + areas[key]['area_no'] + '>' + areas[key]['area_name'] + '</option>')
  }
  $('#area_select').change(this.getDataByArea.bind(this))

}
Graph_Trigger.prototype.getDataByArea = function() {
  console.log('get')
  $.ajax({
    url: '/getData1',
    type: "get",
    dataType: 'json',
  }).then(this.drowChart.bind(this));
}

Graph_Trigger.prototype.drowChart = function(data) {
  var ctx = document.getElementById('myChart1');
  var myChart = new Chart(ctx, {
    type: 'bar',
    data: {
    labels: this.view_s.room_name_list,
    datasets: [
      {
        label: '最高気温(度）',
        data: data,
        backgroundColor: "rgb(0, 255, 0)"
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
    }
  }
});
}

Graph_Trigger.prototype.formatRoomData = function(rooms) {
  for(let key in rooms){
    this.view_s.room_name_list.push(rooms[key]['room_name'])
  }

}
Graph_Trigger.prototype.drowLineChart = function() {
  var ctx = document.getElementById(target);
  var myChart = new Chart(ctx, {
    type: 'line',
    data: {
    labels: dates,
    datasets: [
      {
        label: '最高気温(度）',
        data: json,
        backgroundColor: "rgb(255, 100, 0)"
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
    }
  }
});
}
const chartClass = new Graph_Trigger();

chartClass.createGraph()