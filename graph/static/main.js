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
  this.getDataTypeList();
  this.getDataByArea();
  this.getDates();
  this.getDataByRoom();
}
Graph_Trigger.prototype.getDates = function(areas) {
  $.ajax({
    url: '/getDates',
    type: "get",
    dataType: 'json',
  }).then(this.putDates.bind(this));
}
Graph_Trigger.prototype.putDates = function(dates) {
  this.view_s.dates = dates
}
Graph_Trigger.prototype.getDataTypeList = function(areas) {
  $.ajax({
    url: '/getDataTypeList',
    type: "get",
    dataType: 'json',
  }).then(this.createDataTypeSelectBox.bind(this));
}
Graph_Trigger.prototype.createDataTypeSelectBox = function(data_type_list) {
  $('#data_type_select_area').append(
    '<select id="data_type_select">'
    + '</select>'
  );
  
  for(let key in data_type_list){
    $('#data_type_select').append('<option value=' + data_type_list[key]['data_type_no'] + '>' + data_type_list[key]['data_type_name'] + '</option>')
  }
  // $('#data_type_select').change(this.getDataByArea.bind(this))

}
Graph_Trigger.prototype.createAreaSelectBox = function(areas) {
  $('#area_select_area').append(
    '<select id="area_select">'
    + '</select>'
  );
  
  for(let key in areas){
    $('#area_select').append('<option value=' + areas[key]['area_no'] + '>' + areas[key]['area_name'] + '</option>')
  }
  $('#area_select').change(this.getDataByArea.bind(this))

}
Graph_Trigger.prototype.getDataByArea = function() {
  $.ajax({
    url: '/getData1',
    type: "get",
    dataType: 'json',
  }).then(this.drowChart.bind(this));
}
Graph_Trigger.prototype.getDataByRoom = function() {
  $.ajax({
    url: '/getData1',
    type: "get",
    dataType: 'json',
  }).then(this.drowLineChart.bind(this));
}

Graph_Trigger.prototype.drowChart = function(data) {
  var ctx = document.getElementById('myChart1');
  var myChart = new Chart(ctx, {
    type: 'bar',
    data: {
    labels: this.view_s.room_name_list,
    datasets: [
      {
        label: '平均気温(度）',
        data: data,
        backgroundColor: "rgb(0, 255, 0)"
      },
    ],
  },
  options: {
    title: {
      display: true,
      text: '各部屋の現在の気温'
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
Graph_Trigger.prototype.drowLineChart = function(room_data) {
  var ctx = document.getElementById('myChart2');
  var myChart = new Chart(ctx, {
    type: 'line',
    data: {
    labels: this.view_s.dates,
    datasets: [
      {
        label: '最高気温(度）',
        data: room_data,
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